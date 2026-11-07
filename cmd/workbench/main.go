package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/fang/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/workbench"
	"github.com/spf13/cobra"
)

func main() {
	if err := fang.Execute(context.Background(), command(), fang.WithVersion("0.1.0")); err != nil {
		os.Exit(1)
	}
}

func command() *cobra.Command {
	root := &cobra.Command{Use: "workbench", Short: "A component workbench for Charm", SilenceUsage: true}
	root.AddCommand(&cobra.Command{Use: "list", Short: "List registered Go stories", RunE: func(cmd *cobra.Command, args []string) error {
		for _, d := range story.Definitions() {
			fmt.Fprintf(cmd.OutOrStdout(), "%-26s %s\n", d.ID, d.Title)
		}
		return nil
	}})
	for _, name := range []string{"run", "session", "capture"} {
		name := name
		c := story.DefaultConfig()
		encoded := ""
		output := ""
		plain := false
		cmd := &cobra.Command{Use: name + " STORY", Short: map[string]string{"run": "Run a story in this terminal", "session": "Run an isolated preview session", "capture": "Render a deterministic ANSI frame"}[name], Hidden: name == "session", Args: cobra.ExactArgs(1)}
		cmd.Flags().IntVar(&c.Columns, "columns", c.Columns, "Terminal columns")
		cmd.Flags().IntVar(&c.Rows, "rows", c.Rows, "Terminal rows")
		cmd.Flags().StringVar((*string)(&c.Theme), "theme", string(c.Theme), "dark, light, contrast, or mono")
		cmd.Flags().StringVar(&c.Profile, "profile", c.Profile, "truecolor, ansi256, ansi16, or none")
		cmd.Flags().BoolVar(&c.ASCII, "ascii", false, "Use ASCII/text fallbacks")
		cmd.Flags().BoolVar(&c.ReducedMotion, "reduced-motion", false, "Disable automatic animation")
		cmd.Flags().BoolVar(&c.Live, "live", false, "Run wall-clock animations")
		cmd.Flags().StringVar(&c.State, "state", c.State, "ready, empty, loading, error, or disabled")
		cmd.Flags().StringVar(&c.Value, "value", c.Value, "Input fixture value")
		cmd.Flags().IntVar(&c.Series, "series", c.Series, "Chart series count")
		cmd.Flags().Float64Var(&c.Amount, "amount", c.Amount, "Progress or chart baseline, 0–100")
		cmd.Flags().BoolVar(&c.Horizontal, "horizontal", false, "Use horizontal bars")
		cmd.Flags().StringVar(&encoded, "config", "", "Internal encoded configuration")
		cmd.Flags().MarkHidden("config")
		if name == "capture" {
			cmd.Flags().StringVarP(&output, "output", "o", "", "Write a frame to a file")
			cmd.Flags().BoolVar(&plain, "plain", false, "Strip ANSI styling")
		}
		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			d, err := story.Find(args[0])
			if err != nil {
				return err
			}
			if encoded != "" {
				c, err = workbench.DecodeConfig(encoded)
				if err != nil {
					return err
				}
			} else {
				if !cmd.Flags().Changed("columns") {
					c.Columns = d.Defaults.Columns
				}
				if !cmd.Flags().Changed("rows") {
					c.Rows = d.Defaults.Rows
				}
				if os.Getenv("NO_COLOR") != "" {
					c.Profile = "none"
				}
			}
			if err = c.Validate(); err != nil {
				return err
			}
			if name == "capture" {
				frame, err := story.Frame(args[0], c)
				if err != nil {
					return err
				}
				if plain {
					frame = ansi.Strip(frame)
				}
				if output != "" {
					return os.WriteFile(output, []byte(frame+"\n"), 0644)
				}
				_, err = fmt.Fprintln(cmd.OutOrStdout(), frame)
				return err
			}
			return run(args[0], c)
		}
		root.AddCommand(cmd)
	}
	addr := "127.0.0.1:7331"
	serve := &cobra.Command{Use: "serve", Short: "Serve native terminal previews on loopback", RunE: func(cmd *cobra.Command, args []string) error { return serve(addr) }}
	serve.Flags().StringVar(&addr, "listen", addr, "Loopback address")
	root.AddCommand(serve)
	out := "docs"
	generate := &cobra.Command{Use: "generate", Short: "Export Go stories, MDX, source and static frames", RunE: func(cmd *cobra.Command, args []string) error { return workbench.Generate(out) }}
	generate.Flags().StringVar(&out, "out", out, "Documentation project directory")
	root.AddCommand(generate)
	return root
}

func run(id string, c story.Config) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	var events io.Writer
	if os.Getenv("CHARMED_EVENTS_FD") == "3" {
		file := os.NewFile(3, "story-events")
		defer file.Close()
		events = file
	}
	emit := func(value any) {
		if events != nil {
			_ = json.NewEncoder(events).Encode(value)
		}
	}
	m, err := story.New(id, c, emit)
	if err != nil {
		return err
	}
	defer m.Close()
	_, err = tea.NewProgram(m, tea.WithContext(ctx), tea.WithWindowSize(c.Columns, c.Rows), tea.WithColorProfile(c.ColorProfile())).Run()
	if errors.Is(err, tea.ErrProgramKilled) || errors.Is(err, context.Canceled) {
		return nil
	}
	return err
}

func serve(addr string) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	ip := net.ParseIP(host)
	if host != "localhost" && (ip == nil || !ip.IsLoopback()) {
		return errors.New("preview server must bind to loopback")
	}
	executable, err := os.Executable()
	if err != nil {
		return err
	}
	preview := workbench.NewServer(executable)
	defer preview.Close()
	server := &http.Server{Addr: addr, Handler: preview, ReadHeaderTimeout: 5 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		preview.Close()
		shutdown, cancel := context.WithTimeout(context.Background(), 4*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdown)
	}()
	fmt.Fprintln(os.Stderr, "Charmed preview service: http://"+addr)
	err = server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

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
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/cli"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/workbench"
	"github.com/spf13/cobra"
)

func main() {
	root := command()
	if err := root.ExecuteContext(context.Background()); err != nil {
		printCommandError(root, err)
		os.Exit(1)
	}
}

func command() *cobra.Command {
	root := &cobra.Command{
		Use: "workbench", Short: "A component workbench for Charm",
		Example: "workbench run screen/services\nworkbench capture chart/timeseries --plain",
		Version: "0.1.0", SilenceUsage: true, SilenceErrors: true,
	}
	cli.HelpFormatter{}.Install(root)
	root.AddCommand(manCommand())
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
		cmd.Flags().IntVar(&c.Sigma, "sigma", c.Sigma, "Maximum nested host-spread band: 1, 2 or 3")
		cmd.Flags().StringVar(&c.ComparisonVariant, "comparison-variant", c.ComparisonVariant, "Comparison presentation: stacked or bands")
		cmd.Flags().StringVar(&c.LogPane, "log-pane", c.LogPane, "Log inspector pane: requests, request, response or logs (http aliases response)")
		cmd.Flags().StringVar(&c.SymbolStyle, "symbols", c.SymbolStyle, "Symbol presentation: both, glyph or label")
		cmd.Flags().StringVar(&c.URLStyle, "url-format", c.URLStyle, "URL display: path-query, path, host-path or full")
		cmd.Flags().BoolVar(&c.VerboseDetails, "verbose-details", false, "Label HTTP summary fields")
		cmd.Flags().BoolVar(&c.ColumnLabels, "column-labels", false, "Show HTTP table column labels")
		cmd.Flags().BoolVar(&c.NoFolds, "no-folds", false, "Keep log fields and HTTP sections expanded without fold controls")
		cmd.Flags().BoolVar(&c.NoHelp, "no-help", false, "Hide all keyboard instructions below log fixtures")
		cmd.Flags().BoolVar(&c.JSONFolded, "fold-json", false, "Show a nested JSON container folded")
		cmd.Flags().BoolVar(&c.SingleRequest, "single-request", false, "Show one request with UUID headers and a detailed response")
		cmd.Flags().StringSliceVar(&c.CollapsedPanes, "collapsed-panes", nil, "Fold inspector panes: requests, request, response, logs")
		cmd.Flags().BoolVar(&c.TraceLogs, "trace-logs", false, "Show related logs across the selected trace")
		cmd.Flags().BoolVar(&c.DetachedLogDetails, "detached-log-details", false, "Show selected log fields outside the hierarchy")
		cmd.Flags().StringVar(&c.InspectorLayout, "log-layout", "", "Inspector layout: sidebar, timeline, columns or focused")
		cmd.Flags().StringSliceVar(&c.TimelineGroups, "timeline-group-by", nil, "Group timeline by service, host, port or endpoint")
		cmd.Flags().BoolVar(&c.Timeline, "timeline", false, "Show HTTP requests on a shared time axis")
		cmd.Flags().StringVar((*string)(&c.KeyLayout), "key-layout", "", "Navigation preset: standard or vim")
		cmd.Flags().StringVar((*string)(&c.PaletteBackdrop), "palette-backdrop", "", "Palette backdrop: dim, monochrome or none")
		cmd.Flags().BoolVar(&c.HighlightMatches, "highlight-matches", false, "Underline command search matches")
		cmd.Flags().StringVar(&c.CommandQuery, "command-query", "", "Initial command palette search")
		cmd.Flags().BoolVar(&c.CommandPalette, "command-palette", false, "Open the floating command palette")
		cmd.Flags().BoolVar(&c.RawBody, "raw-body", false, "Show captured HTTP body formatting")
		cmd.Flags().StringVar(&c.Emphasis, "emphasis", c.Emphasis, "Atom emphasis: bright, normal or dim")
		cmd.Flags().StringVar(&c.Scenario, "scenario", c.Scenario, "Deployment fixture scenario (see deployment guide)")
		cmd.Flags().StringVar(&c.ReferenceID, "reference-version", c.ReferenceID, "Reference version identifier")
		cmd.Flags().StringVar(&c.CandidateID, "candidate-version", c.CandidateID, "Candidate version identifier")
		cmd.Flags().BoolVar(&c.Boundaries, "boundaries", false, "Enable illustrative warn/error deviation boundaries")
		cmd.Flags().Float64Var(&c.Amount, "amount", c.Amount, "Progress or chart baseline, 0–100")
		cmd.Flags().BoolVar(&c.Legend, "legend", false, "Show a legend using the plot marker and series styles")
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
				if !cmd.Flags().Changed("boundaries") {
					c.Boundaries = d.Defaults.Boundaries
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
					return os.WriteFile(output, []byte(frame+"\n"), 0o600)
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
	generate := &cobra.Command{Use: "generate", Short: "Build the static guide from Go stories", RunE: func(cmd *cobra.Command, args []string) error { return workbench.Generate(out) }}
	generate.Flags().StringVar(&out, "out", out, "Documentation output parent directory")
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

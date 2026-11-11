package cli_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/spf13/cobra"

	"github.com/lanej/charmed/cli"
	"github.com/lanej/charmed/theme"
)

func TestHelpFromLiveCommandTree(t *testing.T) {
	root := &cobra.Command{Use: "release"}
	root.AddGroup(&cobra.Group{ID: "read", Title: "Read releases:"})
	root.PersistentFlags().String("profile", "local", "Choose a profile")

	show := &cobra.Command{
		Use: "show VERSION", GroupID: "read", Aliases: []string{"inspect"},
		Short: "Show release", Long: "Inspect a release by version.",
		Example: "release show v1 --format json", Run: func(*cobra.Command, []string) { t.Fatal("render executed show") },
	}
	root.AddCommand(show, &cobra.Command{Use: "internal", Hidden: true}, &cobra.Command{Use: "glossary", Short: "Release terms"})
	show.Flags().StringP("format", "f", "table", "Write a `FORMAT` document")
	show.Flags().Bool("trace-internal", false, "Internal traces")

	if err := show.Flags().MarkHidden("trace-internal"); err != nil {
		t.Fatal(err)
	}

	if err := show.MarkFlagRequired("format"); err != nil {
		t.Fatal(err)
	}

	format := cli.HelpFormatter{Width: 86, Theme: theme.New(theme.Mono)}

	rootHelp := format.Format(root)
	for _, want := range []string{"release [command]", "Read releases\n  show VERSION", "Show release", "Help topics\n  glossary"} {
		if !strings.Contains(rootHelp, want) {
			t.Fatalf("root help missing %q:\n%s", want, rootHelp)
		}
	}

	show.Flags().Int("limit", 3, "Maximum number of services")

	page := format.Format(show)
	for _, want := range []string{"Inspect a release by version.", "release show VERSION [flags]", "show, inspect", show.Example, "-f, --format FORMAT", "Write a FORMAT document (required) (default \"table\")", "--limit int", "Global flags\n  --profile string", "(default \"local\")"} {
		if !strings.Contains(page, want) {
			t.Fatalf("command help missing %q:\n%s", want, page)
		}
	}

	if strings.Contains(rootHelp, "internal") || strings.Contains(page, "trace-internal") || strings.Count(page, "--profile") != 1 || page != ansi.Strip(page) {
		t.Fatalf("unexpected hidden, duplicate or styled content:\n%s\n%s", rootHelp, page)
	}
}

func TestHelpReflowsInTerminalCells(t *testing.T) {
	cmd := &cobra.Command{Use: "inspect 対象", Short: "Inspect a release."}
	cmd.Flags().String("地域", "西", "地域ごとに表示する。 Output for every selected region.")
	cmd.Flags().Bool("all", false, "Include every region in the release")

	for _, width := range []int{12, 32, 45, 80} {
		page := (cli.HelpFormatter{Width: width, Theme: theme.New(theme.Dark)}).Format(cmd)

		plain := ansi.Strip(page)
		if plain == page || !strings.Contains(plain, "--地域") {
			t.Fatalf("missing styled flag at %d cells:\n%s", width, page)
		}

		for line := range strings.SplitSeq(plain, "\n") {
			if cells := ansi.StringWidth(line); cells > width {
				t.Errorf("%d-cell line in %d-cell help: %q", cells, width, line)
			}
		}

		if width == 32 && !strings.Contains(plain, "  --all\n    Include every region") {
			t.Errorf("expected stacked description:\n%s", plain)
		}

		if width == 80 && !strings.Contains(plain, "  --all          Include every region") {
			t.Errorf("expected aligned description:\n%s", plain)
		}
	}
}

func TestInstalledHelpUsesCobraRoutesAndOutput(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "0")
	t.Setenv("TTY_FORCE", "0")

	var previous string

	for _, args := range [][]string{{"show", "--help"}, {"show", "-h"}, {"help", "show"}} {
		var stdout, stderr bytes.Buffer

		root := &cobra.Command{Use: "release"}
		root.SetOut(&stdout)
		root.SetErr(&stderr)
		root.SetArgs(args)
		cli.HelpFormatter{}.Install(root)
		root.AddCommand(&cobra.Command{Use: "show", Run: func(*cobra.Command, []string) { t.Fatal("executed instead of displaying help") }})

		if err := root.ExecuteContext(t.Context()); err != nil {
			t.Fatal(err)
		}

		page := stdout.String()
		if !strings.Contains(page, "release show [flags]") || !strings.Contains(page, "-h, --help") || page != ansi.Strip(page) || stderr.Len() != 0 {
			t.Fatalf("%v: stdout %q, stderr %q", args, page, stderr.String())
		}

		if previous != "" && page != previous {
			t.Errorf("%v produced different help:\n%s", args, page)
		}

		previous = page
	}
}

func TestInstalledHelpPreservesChildOverrideAndReportsWriteErrors(t *testing.T) {
	root := &cobra.Command{Use: "release"}
	child := &cobra.Command{Use: "inspect"}
	called := false

	child.SetHelpFunc(func(*cobra.Command, []string) { called = true })
	root.AddCommand(child)
	cli.HelpFormatter{}.Install(root)

	if err := child.Help(); err != nil || !called {
		t.Fatalf("child help override: called=%v err=%v", called, err)
	}

	var stderr bytes.Buffer

	root.SetOut(brokenWriter{})
	root.SetErr(&stderr)

	if err := root.Help(); err != nil || !strings.Contains(stderr.String(), "closed output") {
		t.Fatalf("Cobra help write failure: err=%v diagnostic=%q", err, stderr.String())
	}
}

func TestHelpHonorsPaletteAndNoColor(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	t.Setenv("TERM", "xterm-256color")
	t.Setenv("NO_COLOR", "")

	cmd := &cobra.Command{Use: "release"}
	palette := theme.New(theme.Light)
	palette.Accent = lipgloss.Color("#123456")
	cli.HelpFormatter{Theme: palette}.Install(cmd)

	var output bytes.Buffer
	cmd.SetOut(&output)

	if err := cmd.Help(); err != nil || !strings.Contains(output.String(), "38;2;18;52;86") {
		t.Fatalf("missing custom heading color: %v %q", err, output.String())
	}

	t.Setenv("NO_COLOR", "please")
	output.Reset()

	if err := cmd.Help(); err != nil || output.String() != ansi.Strip(output.String()) {
		t.Fatalf("expected plain output: %v %q", err, output.String())
	}
}

type brokenWriter struct{}

func (brokenWriter) Write([]byte) (int, error) { return 0, errors.New("closed output") }

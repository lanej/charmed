package workbench

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

type Snapshot struct {
	Name   string       `json:"name"`
	Config story.Config `json:"config"`
	Frame  string       `json:"frame"`
}
type ManifestEntry struct {
	story.Definition
	Code      string     `json:"code"`
	Snapshots []Snapshot `json:"snapshots"`
}

// Generate writes reproducible documentation inputs from the compiled Go registry.
func Generate(out string) error {
	if err := story.Validate(); err != nil {
		return err
	}

	root, err := moduleRoot()
	if err != nil {
		return err
	}

	var entries []ManifestEntry

	for _, d := range story.Definitions() {
		code, err := exampleCode(root, d)
		if err != nil {
			return err
		}

		entry := ManifestEntry{Definition: d, Code: code}

		variants := []struct {
			name   string
			change func(*story.Config)
		}{
			{"Default", func(*story.Config) {}},
			{"Narrow", func(c *story.Config) {
				c.Columns, c.Rows = 40, 16
				if d.ID == "foundations/permissions" || d.ID == "diagram/state-transitions" {
					c.Rows = d.Defaults.Rows
				}

				if d.ID == "foundations/glyphs" {
					c.Rows = (len(theme.Glyphs())+1)/2 + 2
				}
			}},
			{"256 colors", func(c *story.Config) { c.Profile = "ansi256" }},
			{"16 colors", func(c *story.Config) { c.Profile = "ansi16" }},
			{"Monochrome", func(c *story.Config) {
				c.Theme = theme.Mono
				c.ASCII = true
				c.Profile = "none"
				c.ReducedMotion = true
			}},
		}
		if d.ID == "diagram/state-transitions" {
			for _, trace := range []struct{ name, state string }{{"Last 3 transitions", "last-3"}, {"Whole path", "all"}} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{
					name: trace.name, change: func(c *story.Config) { c.TraceWindow = trace.state },
				})
			}
		}
		if d.ID == "collection/command-palette" || d.ID == "pattern/log-inspector" {
			for _, sample := range []struct{ name, query string }{{"Commands", ""}, {"Commands filtered", "laytm"}, {"Commands highlighted", "laytm"}, {"Commands unavailable", "Scope: trace"}, {"Commands empty", "zzzzzz"}} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{name: sample.name, change: func(c *story.Config) {
					c.CommandPalette = true
					c.Value = sample.query
					c.HighlightMatches = sample.name == "Commands highlighted"

					c.CommandQuery = sample.query
					if d.ID == "pattern/log-inspector" && sample.name == "Commands unavailable" {
						c.State = "loading"
					}

					c.InspectorLayout = "timeline"
					c.DetachedLogDetails = true
					c.TimelineGroups = []string{"service", "endpoint"}
				}})
			}

			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{name: "Commands narrow", change: func(c *story.Config) { c.CommandPalette = true; c.Columns = 40; c.Rows = 20 }})

		}

		if d.ID == "input/text" {
			variants = append(variants,
				struct {
					name   string
					change func(*story.Config)
				}{"Empty", func(c *story.Config) { c.Value = "" }},
				struct {
					name   string
					change func(*story.Config)
				}{"Disabled", func(c *story.Config) { c.State = "disabled" }},
			)
		}

		if d.ID == "foundations/permissions" || d.Family == "commands" || d.Family == "logs" || d.Family == "http" || d.ID == "text/syntax" {
			for _, mode := range []theme.Mode{theme.Light, theme.Contrast} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{name: string(mode), change: func(c *story.Config) { c.Theme = mode }})
			}
		}

		if d.Family == "logs" && slices.Contains(d.Controls, "state") {
			for _, state := range []string{"empty", "error", "loading"} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{name: strings.ToUpper(state[:1]) + state[1:], change: func(c *story.Config) { c.State = state }})
			}
		}

		if d.ID == "pattern/log-inspector" || d.ID == "collection/log-tree" || d.ID == "collection/http-log" {
			for _, mode := range []string{"glyph", "label"} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{name: strings.ToUpper(mode[:1]) + mode[1:] + " only", change: func(c *story.Config) { c.SymbolStyle = mode }})
			}
		}

		if d.ID == "collection/log-tree" || d.ID == "pattern/log-inspector" {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{name: "Focused leaf", change: func(c *story.Config) { c.DetachedLogDetails = true }})
		}

		if d.ID == "collection/http-log" {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{name: "Timeline", change: func(c *story.Config) { c.Timeline = true; c.Legend = true }}, struct {
				name   string
				change func(*story.Config)
			}{name: "Timeline grouped", change: func(c *story.Config) {
				c.Timeline = true
				c.Legend = true
				c.TimelineGroups = []string{"service", "endpoint"}
			}})
		}

		if d.ID == "pattern/log-inspector" || d.ID == "collection/http-log" {
			for _, format := range []string{"full", "host-path", "path"} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{name: "URL " + format, change: func(c *story.Config) { c.URLStyle = format }})
			}

			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{name: "Labeled summary", change: func(c *story.Config) { c.VerboseDetails = true; c.ColumnLabels = true }})
		}

		if d.ID == "pattern/log-inspector" {
			for _, preset := range []string{"sidebar", "timeline", "columns", "focused"} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{name: "Layout " + preset, change: func(c *story.Config) {
					c.InspectorLayout = preset
					c.DetachedLogDetails = true
					c.Columns = 120

					c.Rows = 44
					if preset == "columns" {
						c.Columns = 160
					}

					if preset == "timeline" {
						c.TimelineGroups = []string{"service", "endpoint"}
					}

					if preset == "focused" {
						c.LogPane = "logs"
						c.Columns = 60
						c.Rows = 30
					}
				}})
			}

			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{name: "Focused narrow", change: func(c *story.Config) { c.DetachedLogDetails = true; c.Columns = 40; c.Rows = 24; c.LogPane = "logs" }})
			for _, variant := range []struct {
				name, pane string
				trace      bool
			}{{"Narrow details", "http", false}, {"Narrow logs", "logs", false}, {"Trace context", "", true}} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{name: variant.name, change: func(c *story.Config) {
					c.LogPane = variant.pane

					c.TraceLogs = variant.trace
					if variant.pane != "" {
						c.Columns, c.Rows = 40, 24
					}
				}})
			}
		}

		if d.ID == "collection/log-tree" || d.ID == "collection/http-log" || d.ID == "pattern/log-inspector" {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{name: "No folds", change: func(c *story.Config) { c.NoFolds = true }})
		}

		if d.ID == "collection/http-log" {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{name: "Raw body", change: func(c *story.Config) { c.RawBody = true }})
		}

		if slices.Contains(d.Controls, "legend") {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{"With legend", func(c *story.Config) { c.Legend = true }})
		}

		if d.ID == "chart/panel" {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{"Narrow with legend", func(c *story.Config) {
				c.Columns, c.Rows = 40, 16
				c.Series, c.Legend = 4, true
			}})
		}

		if slices.Contains(d.Controls, "comparisonVariant") {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{
				name: "Disconnected bands", change: func(c *story.Config) { c.ComparisonVariant = "bands" },
			})
		}

		if slices.Contains(d.Controls, "emphasis") {
			for _, emphasis := range []string{"bright", "normal", "dim"} {
				variants = append(variants, struct {
					name   string
					change func(*story.Config)
				}{
					name: strings.ToUpper(emphasis[:1]) + emphasis[1:], change: func(c *story.Config) { c.Emphasis = emphasis },
				})
			}
		}

		if slices.Contains(d.Controls, "boundaries") {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{name: "No boundaries", change: func(c *story.Config) { c.Boundaries = false }})
		}

		for _, scenario := range d.Scenarios {
			if scenario == d.Defaults.Scenario {
				continue
			}

			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{
				name:   strings.ToUpper(scenario[:1]) + strings.ReplaceAll(scenario[1:], "-", " "),
				change: func(c *story.Config) { c.Scenario = scenario },
			})
		}

		for _, v := range variants {
			c := d.Defaults
			v.change(&c)

			frame, err := story.Frame(d.ID, c)
			if err != nil {
				return fmt.Errorf("capture %s: %w", d.ID, err)
			}

			entry.Snapshots = append(entry.Snapshots, Snapshot{Name: v.name, Config: c, Frame: frame})
		}

		entries = append(entries, entry)
	}

	return writeSite(out, root, entries)
}

func moduleRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
		if err == nil && strings.Contains(string(data), "module github.com/lanej/charmed") {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("run generate from the Charmed repository")
		}

		dir = parent
	}
}

func sourceSymbol(path, symbol string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	if symbol == "" {
		return string(data), nil
	}

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, path, data, parser.ParseComments)
	if err != nil {
		return "", err
	}

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Name.Name == symbol {
			start := fset.Position(fn.Pos()).Offset
			if fn.Doc != nil {
				start = fset.Position(fn.Doc.Pos()).Offset
			}

			end := fset.Position(fn.End()).Offset

			return string(data[start:end]), nil
		}
	}

	return "", fmt.Errorf("source symbol %s missing from %s", symbol, path)
}

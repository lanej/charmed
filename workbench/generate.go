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
			{"Narrow", func(c *story.Config) { c.Columns = 40; c.Rows = 16 }},
			{"256 colors", func(c *story.Config) { c.Profile = "ansi256" }},
			{"16 colors", func(c *story.Config) { c.Profile = "ansi16" }},
			{"Monochrome", func(c *story.Config) {
				c.Theme = theme.Mono
				c.ASCII = true
				c.Profile = "none"
				c.ReducedMotion = true
			}},
		}
		if slices.Contains(d.Controls, "legend") {
			variants = append(variants, struct {
				name   string
				change func(*story.Config)
			}{"With legend", func(c *story.Config) { c.Legend = true }})
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

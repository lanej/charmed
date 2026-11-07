package workbench

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
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
	root, err := moduleRoot()
	if err != nil {
		return err
	}
	directory := filepath.Join(out, "generated")
	if err = os.MkdirAll(directory, 0755); err != nil {
		return err
	}
	var entries []ManifestEntry
	for _, d := range story.Definitions() {
		code, err := sourceSymbol(filepath.Join(root, d.Source), d.Symbol)
		if err != nil {
			return err
		}
		entry := ManifestEntry{Definition: d, Code: code}
		variants := []struct {
			name   string
			change func(*story.Config)
		}{
			{"Default", func(*story.Config) {}},
			{"Light", func(c *story.Config) { c.Theme = theme.Light }},
			{"Narrow", func(c *story.Config) { c.Columns = 40; c.Rows = 16 }},
			{"Monochrome", func(c *story.Config) {
				c.Theme = theme.Mono
				c.ASCII = true
				c.Profile = "none"
				c.ReducedMotion = true
			}},
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
		filename := strings.ReplaceAll(d.ID, "/", "_")
		if err = writeStories(directory, filename, entry); err != nil {
			return err
		}
	}
	if err = writeJSONFile(filepath.Join(directory, "manifest.json"), entries); err != nil {
		return err
	}
	return writeJSONFile(filepath.Join(directory, "coverage.json"), CoverageInventory())
}

func writeJSONFile(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0644)
}
func js(value any) string { data, _ := json.Marshal(value); return string(data) }

func writeStories(directory, filename string, entry ManifestEntry) error {
	parts := strings.Split(entry.Title, "/")
	name := parts[len(parts)-1]
	mdx := fmt.Sprintf("import { StoryCanvas } from '../components/StoryCanvas';\nimport { GoSource } from '../components/GoSource';\n\n# %s\n\n%s\n\n<StoryCanvas storyId=%s />\n\n## Go example\n\n<GoSource storyId=%s />\n", name, entry.Description, js(entry.ID), js(entry.ID))
	return os.WriteFile(filepath.Join(directory, filename+".mdx"), []byte(mdx), 0644)
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

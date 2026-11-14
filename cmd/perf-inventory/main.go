// Command perf-inventory lists native fixture coverage without executing commands.
package main

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
)

type (
	surface struct {
		Symbol     string   `json:"symbol"`
		Source     string   `json:"source"`
		Examples   []string `json:"consumer_references"`
		Kind       string   `json:"kind"`
		Benchmarks []string `json:"benchmarks"`
		Status     string   `json:"status"`
	}
	inventory struct {
		Fixtures []story.Definition `json:"fixtures"`
		Surfaces []surface          `json:"surfaces"`
	}
)

func typeName(expression ast.Expr) string {
	switch value := expression.(type) {
	case *ast.Ident:
		return value.Name
	case *ast.StarExpr:
		return typeName(value.X)
	case *ast.IndexExpr:
		return typeName(value.X)
	case *ast.IndexListExpr:
		return typeName(value.X)
	default:
		return ""
	}
}

// These indirect/enum views have no constructor reference in consumer snippets.
// Links identify exercised workloads, not proof of every internal branch.
var explicitViews = map[string][]string{
	"feedback.Evidence.View": {"feedback/evidence"},
	"feedback.Outcome.View":  {"feedback/outcome"},
	"feedback.Status.View":   {"feedback/status"},
	"gitui.ChangeGroup.View": {"git/staged", "git/unstaged", "git/untracked", "git/conflicts"},
	"gitui.List.View":        {"gitui/BenchmarkListLifecycle"},
	"text.ChangeKind.View":   {"text/change"},
	"text.ValueState.View":   {"text/value-state"},
}

func collect() (inventory, error) {
	result := inventory{Fixtures: story.Definitions()}

	fixtureIDs := map[string]bool{}
	for _, definition := range result.Fixtures {
		fixtureIDs[definition.ID] = true
	}

	references := map[string][]string{}

	for _, definition := range result.Fixtures {
		file, err := parser.ParseFile(token.NewFileSet(), definition.ExampleSource, nil, 0)
		if err != nil {
			return inventory{}, err
		}

		aliases := map[string]string{}

		for _, imp := range file.Imports {
			path := strings.Trim(imp.Path.Value, "\"")
			if !strings.HasPrefix(path, "github.com/lanej/charmed/") {
				continue
			}

			alias := filepath.Base(path)
			if imp.Name != nil {
				alias = imp.Name.Name
			}

			aliases[alias] = filepath.Base(path)
		}

		ast.Inspect(file, func(node ast.Node) bool {
			selector, ok := node.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			ident, ok := selector.X.(*ast.Ident)
			if !ok {
				return true
			}

			if pkg := aliases[ident.Name]; pkg != "" {
				key := pkg + "." + selector.Sel.Name
				references[key] = append(references[key], definition.ID)
			}

			return true
		})
	}

	for _, directory := range []string{"chart", "collection", "feedback", "gitui", "input", "layout", "navigation", "pattern", "text", "theme", "tui", "runtime"} {
		files, err := filepath.Glob(filepath.Join(directory, "*.go"))
		if err != nil {
			return inventory{}, err
		}

		for _, name := range files {
			if strings.HasSuffix(name, "_test.go") {
				continue
			}

			file, parseErr := parser.ParseFile(token.NewFileSet(), name, nil, 0)
			if parseErr != nil {
				return inventory{}, parseErr
			}

			for _, declaration := range file.Decls {
				if group, ok := declaration.(*ast.GenDecl); ok {
					for _, spec := range group.Specs {
						if declared, ok := spec.(*ast.TypeSpec); ok && declared.Name.IsExported() {
							symbol := directory + "." + declared.Name.Name
							result.Surfaces = append(result.Surfaces, surface{Symbol: symbol, Source: name, Examples: references[symbol], Kind: "type", Status: "data-provider-or-view-type"})
						}
					}
				}

				function, ok := declaration.(*ast.FuncDecl)
				if !ok || !function.Name.IsExported() {
					continue
				}

				symbol := directory + "." + function.Name.Name
				kind := "function"
				examples := references[symbol]

				if function.Recv != nil {
					if function.Name.Name != "View" {
						continue
					}

					owner := typeName(function.Recv.List[0].Type)
					if !ast.IsExported(owner) {
						continue
					}

					symbol = directory + "." + owner + ".View"
					examples = append(examples, references[directory+"."+owner]...)
					examples = append(examples, references[directory+".New"+owner]...)
					kind = "view"
				}

				benchmarks := []string{}
				status := "support-or-unmapped-function"

				for _, example := range examples {
					benchmarks = append(benchmarks, "story/BenchmarkCatalog/"+example)
				}

				if kind == "view" {
					status = "workload-linked-view"

					if symbol == "runtime.Root.View" {
						benchmarks = append(benchmarks, "story/BenchmarkCatalog")
					}

					if len(benchmarks) == 0 {
						for _, id := range explicitViews[symbol] {
							if id == "gitui/BenchmarkListLifecycle" {
								benchmarks = append(benchmarks, id)
							} else {
								if _, exists := fixtureIDs[id]; !exists {
									return inventory{}, fmt.Errorf("unknown performance fixture %s", id)
								}

								benchmarks = append(benchmarks, "story/BenchmarkCatalog/"+id)
							}
						}
					}

					if len(benchmarks) == 0 {
						return inventory{}, fmt.Errorf("unmapped public view %s; add fixture or direct benchmark", symbol)
					}
				}

				result.Surfaces = append(result.Surfaces, surface{Symbol: symbol, Source: name, Examples: examples, Kind: kind, Benchmarks: benchmarks, Status: status})
			}
		}
	}

	return result, nil
}

func run() error {
	result, err := collect()
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(result)
}

func main() {
	if err := run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(2)
		}

		os.Exit(1)
	}
}

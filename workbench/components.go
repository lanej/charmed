package workbench

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"net/url"
	"path/filepath"
	"slices"

	"github.com/lanej/charmed/catalog"
	"github.com/lanej/charmed/story"
)

type (
	relationLink struct{ ID, Name, Kind, Role, Href, SourceHref string }
	useLink      struct{ Name, Role, Href, SourceHref string }
	exampleLink  struct {
		ID, Name, Kind, Href, SourceHref string
		Uses                             []useLink
	}
	componentCard struct {
		catalog.Component
		API, Href, SourceHref, CoverageLabel, EmptyLabel string
		Dependencies, Dependents                         []relationLink
		Examples                                         []exampleLink
	}
)

type componentGuide struct {
	Cards         []componentCard
	Mapped, Total int
}

func (g *componentGuide) find(id string) *componentCard {
	for i := range g.Cards {
		if g.Cards[i].ID == id {
			return &g.Cards[i]
		}
	}

	return nil
}

func buildComponentGuide(root string, entries []ManifestEntry) (*componentGuide, error) {
	graph, err := story.ComponentGraph()
	if err != nil {
		return nil, err
	}

	return projectComponentGuide(root, entries, graph)
}

func projectComponentGuide(root string, entries []ManifestEntry, graph *catalog.Graph) (*componentGuide, error) {
	data := graph.Data()
	guide := &componentGuide{Total: len(entries)}

	byEntry := make(map[string]ManifestEntry, len(entries))
	for _, e := range entries {
		byEntry[e.ID] = e
	}

	sources := sourceLinks{root: root, files: map[string]map[string]int{}}
	for _, c := range data.Components {
		source, sourceErr := sources.link(c.Source)
		if sourceErr != nil {
			return nil, sourceErr
		}

		card := componentCard{Component: c, API: c.Package + "." + c.Symbol, Href: "dependencies.html#component-" + url.PathEscape(c.ID), SourceHref: source}
		switch c.Coverage {
		case catalog.Unaudited:
			card.CoverageLabel = "Not audited"
			card.EmptyLabel = "No relationships recorded; this component has not been audited."
		case catalog.Partial:
			card.CoverageLabel = "Partial map"
			card.EmptyLabel = "No relationships recorded in this partial map."
		case catalog.Complete:
			card.CoverageLabel = "Audited map"
			card.EmptyLabel = "No reusable parts declared in the audited scope."
		default:
			return nil, fmt.Errorf("unknown component coverage %q", c.Coverage)
		}

		guide.Cards = append(guide.Cards, card)
	}

	for _, r := range data.Relations {
		source, sourceErr := sources.link(r.Source)
		if sourceErr != nil {
			return nil, sourceErr
		}

		from, to := guide.find(r.From), guide.find(r.To)
		from.Dependencies = append(from.Dependencies, relationLink{ID: to.ID, Name: to.Name, Kind: string(r.Kind), Role: r.Role, Href: to.Href, SourceHref: source})
		to.Dependents = append(to.Dependents, relationLink{ID: from.ID, Name: from.Name, Kind: string(r.Kind), Role: r.Role, Href: from.Href, SourceHref: source})
	}

	mapped := map[string]bool{}

	for _, e := range data.Examples {
		entry, exists := byEntry[e.EntryID]
		if !exists {
			return nil, fmt.Errorf("unknown example entry %s", e.EntryID)
		}

		source, sourceErr := sources.link(e.Source)
		if sourceErr != nil {
			return nil, sourceErr
		}

		href := storyPage(entry) + "#" + url.PathEscape(e.EntryID)
		kind := "Consumer usage"

		switch e.Kind {
		case catalog.ConsumerExample:
		case catalog.InteractiveExample:
			kind = "Interactive fixture"
			href = storyPage(entry) + "#panel-" + url.PathEscape(e.EntryID) + "-0"
		default:
			return nil, fmt.Errorf("unknown example kind %q", e.Kind)
		}

		view := exampleLink{ID: e.ID, Name: entry.Title, Kind: kind, Href: href, SourceHref: source}
		for _, u := range e.Uses {
			evidence, evidenceErr := sources.link(u.Source)
			if evidenceErr != nil {
				return nil, evidenceErr
			}

			c := guide.find(u.Component)
			view.Uses = append(view.Uses, useLink{Name: c.Name, Role: u.Role, Href: c.Href, SourceHref: evidence})
		}

		seenComponents := map[string]bool{}
		for _, u := range e.Uses {
			if seenComponents[u.Component] {
				continue
			}

			seenComponents[u.Component] = true
			card := guide.find(u.Component)
			copyView := view
			copyView.Uses = slices.Clone(view.Uses)
			card.Examples = append(card.Examples, copyView)
		}

		mapped[e.EntryID] = true
	}

	guide.Mapped = len(mapped)

	for i := range entries {
		examples := graph.ExamplesForEntry(entries[i].ID)
		if len(examples) > 0 {
			entries[i].Component = guide.find(examples[0].Primary)
			entries[i].ComponentExamples = examplesForEntry(guide, examples[0].Primary, examples)
		}
	}

	return guide, nil
}

func examplesForEntry(guide *componentGuide, id string, examples []catalog.Example) []exampleLink {
	var result []exampleLink

	for _, e := range examples {
		for _, view := range guide.find(id).Examples {
			if view.ID == e.ID {
				result = append(result, view)
				break
			}
		}
	}

	return result
}

func writeComponentGraph(directory string) error {
	graph, err := story.ComponentGraph()
	if err != nil {
		return err
	}

	document := struct {
		Version int `json:"version"`
		catalog.Data
	}{Version: 1, Data: graph.Data()}

	data, err := json.MarshalIndent(document, "", "  ")
	if err != nil {
		return err
	}

	return writeGuideFile(directory, "component-graph.json", append(data, '\n'))
}

type sourceLinks struct {
	root  string
	files map[string]map[string]int
}

func (s *sourceLinks) link(source catalog.Source) (string, error) {
	symbols, exists := s.files[source.Path]
	if !exists {
		fset := token.NewFileSet()

		file, err := parser.ParseFile(fset, filepath.Join(s.root, source.Path), nil, 0)
		if err != nil {
			return "", err
		}

		symbols = map[string]int{}

		ast.Inspect(file, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.TypeSpec:
				symbols[n.Name.Name] = fset.Position(n.Pos()).Line
				if structure, ok := n.Type.(*ast.StructType); ok {
					for _, field := range structure.Fields.List {
						for _, name := range field.Names {
							symbols[n.Name.Name+"."+name.Name] = fset.Position(field.Pos()).Line
						}
					}
				}
			case *ast.FuncDecl:
				name := n.Name.Name
				if n.Recv != nil && len(n.Recv.List) > 0 {
					name = receiverName(n.Recv.List[0].Type) + "." + name
				}

				symbols[name] = fset.Position(n.Pos()).Line
			}

			return true
		})

		s.files[source.Path] = symbols
	}

	line, exists := symbols[source.Symbol]
	if !exists {
		return "", fmt.Errorf("component evidence %s:%s references no declaration", source.Path, source.Symbol)
	}

	return fmt.Sprintf("https://github.com/lanej/charmed/blob/main/%s#L%d", source.Path, line), nil
}

func receiverName(expr ast.Expr) string {
	switch n := expr.(type) {
	case *ast.Ident:
		return n.Name
	case *ast.StarExpr:
		return receiverName(n.X)
	case *ast.IndexExpr:
		return receiverName(n.X)
	case *ast.IndexListExpr:
		return receiverName(n.X)
	default:
		return ""
	}
}

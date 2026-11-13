// Package catalog describes reusable component relationships independently of examples and rendering.
package catalog

import (
	"fmt"
	"path"
	"slices"
	"strings"
	"unicode"
)

type Level string

const (
	Foundations Level = "Foundations"
	Atoms       Level = "Atoms"
	Molecules   Level = "Molecules"
	Organisms   Level = "Organisms"
	Templates   Level = "Templates"
	Screens     Level = "Screens"
)

func Levels() []Level { return []Level{Foundations, Atoms, Molecules, Organisms, Templates, Screens} }

type Coverage string

const (
	// Unaudited is intentional zero: no recorded edges never establishes a leaf.
	Unaudited Coverage = ""
	Partial   Coverage = "partial"
	Complete  Coverage = "complete"
)

type RelationKind string

const (
	Composes    RelationKind = "composes"
	Specializes RelationKind = "specializes"
)

// Source identifies a repository Go declaration, including an optional field or method.
// The pure model validates its shape; publication verifies the declaration exists.
type Source struct {
	Path   string `json:"path"`
	Symbol string `json:"symbol"`
}

// Component is a reusable capability, not a fixture or a layout occurrence.
// Level is declared classification; Coverage scopes the completeness of recorded relationships.
type Component struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Package  string   `json:"package"`
	Symbol   string   `json:"symbol"`
	Level    Level    `json:"level"`
	Coverage Coverage `json:"coverage"`
	Source   Source   `json:"source"`
}

// Relation points from a component to one reusable part or specialized base.
// Different roles may reference the same target without duplicating its identity.
type Relation struct {
	From   string       `json:"from"`
	To     string       `json:"to"`
	Kind   RelationKind `json:"kind"`
	Role   string       `json:"role"`
	Source Source       `json:"source"`
}

type Use struct {
	Component string `json:"component"`
	Role      string `json:"role"`
	Source    Source `json:"source"`
}

type ExampleKind string

const (
	ConsumerExample    ExampleKind = "consumer"
	InteractiveExample ExampleKind = "interactive"
)

// Example is one source-backed consumer or fixture. EntryID retains its catalog URL.
// Multiple examples can demonstrate the same component without changing its dependencies.
type Example struct {
	ID      string      `json:"id"`
	EntryID string      `json:"entryId"`
	Primary string      `json:"primary"`
	Kind    ExampleKind `json:"kind"`
	Source  Source      `json:"source"`
	Uses    []Use       `json:"uses"`
}

type Data struct {
	Components []Component `json:"components"`
	Relations  []Relation  `json:"relations"`
	Examples   []Example   `json:"examples"`
}

// Graph owns a validated copy. Queries and Data return detached slices; declaration order is retained.
type Graph struct {
	data Data
	byID map[string]Component
}

func New(data Data) (*Graph, error) {
	g := &Graph{data: clone(data), byID: make(map[string]Component, len(data.Components))}
	if err := g.validate(); err != nil {
		return nil, err
	}

	return g, nil
}

func clone(data Data) Data {
	result := Data{Components: slices.Clone(data.Components), Relations: slices.Clone(data.Relations), Examples: slices.Clone(data.Examples)}
	for i := range result.Examples {
		result.Examples[i].Uses = slices.Clone(data.Examples[i].Uses)
	}

	return result
}

func (g *Graph) Data() Data                            { return clone(g.data) }
func (g *Graph) Component(id string) (Component, bool) { c, ok := g.byID[id]; return c, ok }
func (g *Graph) Dependencies(id string) []Relation {
	var result []Relation

	for _, r := range g.data.Relations {
		if r.From == id {
			result = append(result, r)
		}
	}

	return result
}

func (g *Graph) Dependents(id string) []Relation {
	var result []Relation

	for _, r := range g.data.Relations {
		if r.To == id {
			result = append(result, r)
		}
	}

	return result
}

func (g *Graph) ExamplesUsing(id string) []Example {
	var result []Example

	for _, e := range g.data.Examples {
		for _, u := range e.Uses {
			if u.Component == id {
				e.Uses = slices.Clone(e.Uses)
				result = append(result, e)

				break
			}
		}
	}

	return result
}

func (g *Graph) ExamplesForEntry(id string) []Example {
	var result []Example

	for _, e := range g.data.Examples {
		if e.EntryID == id {
			e.Uses = slices.Clone(e.Uses)
			result = append(result, e)
		}
	}

	return result
}

func (g *Graph) Example(id string) (Example, bool) {
	for _, e := range g.data.Examples {
		if e.ID == id {
			e.Uses = slices.Clone(e.Uses)
			return e, true
		}
	}

	return Example{}, false
}

func validID(id string) bool {
	for _, part := range strings.Split(id, "/") {
		if part == "" {
			return false
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return false
			}
		}
	}

	return true
}

func validText(s string) bool {
	return strings.TrimSpace(s) != "" && strings.IndexFunc(s, unicode.IsControl) < 0
}

func (s Source) validate() error {
	if !validText(s.Path) || path.IsAbs(s.Path) || path.Clean(s.Path) != s.Path || strings.HasPrefix(s.Path, "../") || strings.Contains(s.Path, "\\") || !strings.HasSuffix(s.Path, ".go") || !validText(s.Symbol) {
		return fmt.Errorf("invalid source %q:%q", s.Path, s.Symbol)
	}

	return nil
}

func rank(level Level) int {
	switch level {
	case Foundations:
		return 0
	case Atoms:
		return 1
	case Molecules:
		return 2
	case Organisms:
		return 3
	case Templates:
		return 4
	case Screens:
		return 5
	default:
		return -1
	}
}

func (g *Graph) validate() error {
	for _, c := range g.data.Components {
		if !validID(c.ID) || !validText(c.Name) || !validText(c.Package) || !validText(c.Symbol) || rank(c.Level) < 0 {
			return fmt.Errorf("invalid component %q", c.ID)
		}

		if _, exists := g.byID[c.ID]; exists {
			return fmt.Errorf("duplicate component %q", c.ID)
		}

		switch c.Coverage {
		case Unaudited, Partial, Complete:
		default:
			return fmt.Errorf("invalid coverage for %q", c.ID)
		}

		if err := c.Source.validate(); err != nil {
			return fmt.Errorf("component %s: %w", c.ID, err)
		}

		g.byID[c.ID] = c
	}

	seen := map[string]bool{}

	for _, r := range g.data.Relations {
		from, fromOK := g.byID[r.From]

		to, toOK := g.byID[r.To]
		if !fromOK || !toOK {
			return fmt.Errorf("relationship references unknown component: %s -> %s", r.From, r.To)
		}

		switch r.Kind {
		case Composes, Specializes:
		default:
			return fmt.Errorf("invalid relationship kind %q", r.Kind)
		}

		if rank(to.Level) > rank(from.Level) {
			return fmt.Errorf("component %s cannot depend on higher-level %s", r.From, r.To)
		}

		if !validText(r.Role) {
			return fmt.Errorf("relationship %s -> %s needs a role", r.From, r.To)
		}

		if err := r.Source.validate(); err != nil {
			return err
		}

		key := r.From + "\x00" + string(r.Kind) + "\x00" + r.Role + "\x00" + r.To
		if seen[key] {
			return fmt.Errorf("duplicate relationship %s -> %s", r.From, r.To)
		}

		seen[key] = true
	}

	if err := g.validateExamples(); err != nil {
		return err
	}

	return g.validateCycles()
}

func (g *Graph) validateExamples() error {
	seen := map[string]bool{}
	for _, e := range g.data.Examples {
		if !validID(e.ID) || !validID(e.EntryID) || seen[e.ID] {
			return fmt.Errorf("invalid or duplicate example %q", e.ID)
		}

		seen[e.ID] = true
		switch e.Kind {
		case ConsumerExample, InteractiveExample:
		default:
			return fmt.Errorf("invalid example kind %q", e.Kind)
		}

		if err := e.Source.validate(); err != nil {
			return err
		}

		primary := false
		roles := map[string]bool{}

		for _, u := range e.Uses {
			if _, ok := g.byID[u.Component]; !ok {
				return fmt.Errorf("example %s uses unknown component %s", e.ID, u.Component)
			}

			if !validText(u.Role) || roles[u.Role] {
				return fmt.Errorf("example %s needs unique nonempty usage roles", e.ID)
			}

			if err := u.Source.validate(); err != nil {
				return err
			}

			roles[u.Role] = true
			primary = primary || u.Component == e.Primary
		}

		if !primary {
			return fmt.Errorf("example %s must use its primary component %s", e.ID, e.Primary)
		}
	}

	return nil
}

func (g *Graph) validateCycles() error {
	visiting, visited := map[string]bool{}, map[string]bool{}

	var visit func(string) error

	visit = func(id string) error {
		if visiting[id] {
			return fmt.Errorf("component relationship cycle at %s", id)
		}

		if visited[id] {
			return nil
		}

		visiting[id] = true
		for _, r := range g.Dependencies(id) {
			if err := visit(r.To); err != nil {
				return err
			}
		}

		delete(visiting, id)
		visited[id] = true

		return nil
	}
	for _, c := range g.data.Components {
		if err := visit(c.ID); err != nil {
			return err
		}
	}

	return nil
}

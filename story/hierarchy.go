package story

import (
	"fmt"
	"slices"
	"strings"

	"github.com/lanej/charmed/catalog"
)

// Levels returns the catalog's learning order, independent of Go import paths.
func Levels() []string {
	levels := catalog.Levels()

	result := make([]string, len(levels))
	for i, l := range levels {
		result[i] = string(l)
	}

	return result
}

// Validate checks the catalog's declared composition graph before publication.
// Peers can compose peers, but a lower level cannot depend on a higher one.
func Validate() error {
	definitions := Definitions()
	if err := validateDefinitions(definitions); err != nil {
		return err
	}

	graph, err := ComponentGraph()
	if err != nil {
		return err
	}

	entries := make(map[string]bool, len(definitions))
	for _, d := range definitions {
		entries[d.ID] = true
	}

	for _, e := range graph.Data().Examples {
		if !entries[e.EntryID] {
			return fmt.Errorf("example %s references unknown catalog entry %s", e.ID, e.EntryID)
		}
	}

	return nil
}

func validateDefinitions(definitions []Definition) error {
	byID := make(map[string]Definition, len(definitions))
	for _, d := range definitions {
		if _, exists := byID[d.ID]; exists {
			return fmt.Errorf("duplicate story %q", d.ID)
		}
		if slices.Index(Levels(), d.Level) < 0 {
			return fmt.Errorf("story %q has unknown atomic level %q", d.ID, d.Level)
		}
		for _, segment := range strings.Split(d.ID, "/") {
			if segment == "" {
				return fmt.Errorf("invalid story ID %q", d.ID)
			}
			for _, r := range segment {
				if !(r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '-') {
					return fmt.Errorf("invalid story ID %q", d.ID)
				}
			}
		}
		byID[d.ID] = d
	}
	for _, d := range definitions {
		seen := make(map[string]bool)
		for _, id := range d.Composes {
			child, exists := byID[id]
			if !exists {
				return fmt.Errorf("story %q composes unknown story %q", d.ID, id)
			}
			if seen[id] {
				return fmt.Errorf("story %q repeats component %q", d.ID, id)
			}
			seen[id] = true
			if slices.Index(Levels(), child.Level) > slices.Index(Levels(), d.Level) {
				return fmt.Errorf("%s story %q cannot compose %s story %q", d.Level, d.ID, child.Level, id)
			}
		}
	}
	visiting, visited := make(map[string]bool), make(map[string]bool)
	var visit func(string) error
	visit = func(id string) error {
		if visiting[id] {
			return fmt.Errorf("composition cycle at %q", id)
		}
		if visited[id] {
			return nil
		}
		visiting[id] = true
		for _, child := range byID[id].Composes {
			if err := visit(child); err != nil {
				return err
			}
		}
		delete(visiting, id)
		visited[id] = true
		return nil
	}
	for _, d := range definitions {
		if err := visit(d.ID); err != nil {
			return err
		}
	}
	return nil
}

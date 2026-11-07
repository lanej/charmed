package story

import (
	"strings"
	"testing"
)

func TestCatalogCompositionGraph(t *testing.T) {
	if err := Validate(); err != nil {
		t.Fatal(err)
	}
	for _, d := range Definitions() {
		if !strings.HasPrefix(d.MDXPath(), strings.ToLower(d.Level)+"/") {
			t.Fatalf("%s lost its atomic directory", d.ID)
		}
	}
}

func TestCompositionRejectsBrokenLinksAndLayerInversions(t *testing.T) {
	for _, tc := range []struct {
		name        string
		definitions []Definition
		message     string
	}{
		{"missing", []Definition{{ID: "parent", Level: "Molecules", Composes: []string{"missing"}}}, "unknown story"},
		{"inversion", []Definition{{ID: "atom", Level: "Atoms", Composes: []string{"screen"}}, {ID: "screen", Level: "Screens"}}, "cannot compose"},
		{"peer-cycle", []Definition{{ID: "a", Level: "Molecules", Composes: []string{"b"}}, {ID: "b", Level: "Molecules", Composes: []string{"a"}}}, "cycle"},
		{"self-cycle", []Definition{{ID: "a", Level: "Atoms", Composes: []string{"a"}}}, "cycle"},
		{"repeated-story", []Definition{{ID: "a", Level: "Atoms"}, {ID: "a", Level: "Atoms"}}, "duplicate story"},
		{"repeated-child", []Definition{{ID: "a", Level: "Atoms"}, {ID: "b", Level: "Molecules", Composes: []string{"a", "a"}}}, "repeats component"},
		{"unknown-level", []Definition{{ID: "a", Level: "Other"}}, "unknown atomic level"},
		{"path-traversal", []Definition{{ID: "../escape", Level: "Atoms"}}, "invalid story ID"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := validateDefinitions(tc.definitions); err == nil || !strings.Contains(err.Error(), tc.message) {
				t.Fatalf("expected %s, got %v", tc.message, err)
			}
		})
	}
	if err := validateDefinitions([]Definition{{ID: "atom", Level: "Atoms"}, {ID: "inner", Level: "Molecules", Composes: []string{"atom"}}, {ID: "outer", Level: "Molecules", Composes: []string{"inner"}}}); err != nil {
		t.Fatalf("valid peer composition rejected: %v", err)
	}
}

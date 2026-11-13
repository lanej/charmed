package catalog_test

import (
	"strings"
	"testing"

	"github.com/lanej/charmed/catalog"
)

func fixture() catalog.Data {
	source := catalog.Source{Path: "example/parts.go", Symbol: "Part"}

	return catalog.Data{
		Components: []catalog.Component{
			{ID: "base", Name: "Base", Package: "example", Symbol: "Base", Level: catalog.Molecules, Coverage: catalog.Partial, Source: source},
			{ID: "owner", Name: "Owner", Package: "example", Symbol: "Owner", Level: catalog.Organisms, Coverage: catalog.Partial, Source: source},
			{ID: "other", Name: "Other", Package: "example", Symbol: "Other", Level: catalog.Organisms, Source: source},
		},
		Relations: []catalog.Relation{{From: "owner", To: "base", Kind: catalog.Composes, Role: "request", Source: source}, {From: "other", To: "base", Kind: catalog.Specializes, Role: "base", Source: source}},
		Examples:  []catalog.Example{{ID: "sample/usage", EntryID: "sample", Primary: "owner", Kind: catalog.ConsumerExample, Source: source, Uses: []catalog.Use{{Component: "owner", Role: "primary", Source: source}, {Component: "base", Role: "control", Source: source}}}},
	}
}

func TestGraphRetainsSharedIdentityAndOwnsItsData(t *testing.T) {
	data := fixture()

	graph, err := catalog.New(data)
	if err != nil {
		t.Fatal(err)
	}

	data.Components[0].Name = "changed"
	data.Examples[0].Uses[0].Component = "changed"

	if c, ok := graph.Component("base"); !ok || c.Name != "Base" {
		t.Fatal("component ownership lost")
	}

	if c, _ := graph.Component("other"); c.Coverage != catalog.Unaudited {
		t.Fatal("empty relationships inferred audited coverage")
	}

	if len(graph.Dependents("base")) != 2 || graph.Dependencies("other")[0].Kind != catalog.Specializes {
		t.Fatal("shared identity or relation kind lost")
	}

	snapshot := graph.Data()
	snapshot.Relations[0].Role = "changed"
	snapshot.Examples[0].Uses[0].Component = "changed"
	examples := graph.ExamplesUsing("owner")
	examples[0].Uses[0].Component = "changed"
	entry := graph.ExamplesForEntry("sample")
	entry[0].Uses[0].Component = "changed"

	e, ok := graph.Example("sample/usage")
	if !ok || e.Uses[0].Component != "owner" || graph.Dependencies("owner")[0].Role != "request" {
		t.Fatal("query exposed mutable storage")
	}

	if len(graph.Dependencies("missing")) != 0 {
		t.Fatal("unknown query fabricated a relationship")
	}
}

func TestGraphAllowsRepeatedTypesInDistinctRoles(t *testing.T) {
	data := fixture()
	response := data.Relations[0]
	response.Role = "response"
	data.Relations = append(data.Relations, response)
	data.Examples[0].Uses = append(data.Examples[0].Uses, catalog.Use{Component: "base", Role: "another control", Source: response.Source})

	graph, err := catalog.New(data)
	if err != nil {
		t.Fatal(err)
	}

	if len(graph.Dependencies("owner")) != 2 {
		t.Fatal("distinct roles collapsed")
	}

	if len(graph.ExamplesUsing("base")) != 1 {
		t.Fatal("repeated use duplicated the example")
	}
}

func TestGraphRejectsBrokenRelationships(t *testing.T) {
	cases := []struct {
		name, message string
		change        func(*catalog.Data)
	}{
		{"unknown endpoint", "unknown component", func(d *catalog.Data) { d.Relations[0].To = "missing" }},
		{"unknown kind", "relationship kind", func(d *catalog.Data) { d.Relations[0].Kind = "inherits" }},
		{"duplicate identity", "duplicate component", func(d *catalog.Data) { d.Components = append(d.Components, d.Components[0]) }},
		{"duplicate edge", "duplicate relationship", func(d *catalog.Data) { d.Relations = append(d.Relations, d.Relations[0]) }},
		{"empty role", "needs a role", func(d *catalog.Data) { d.Relations[0].Role = "" }},
		{"layer inversion", "higher-level", func(d *catalog.Data) { d.Components[0].Level = catalog.Screens }},
		{"cycle", "cycle", func(d *catalog.Data) {
			d.Components[0].Level = catalog.Organisms
			reverse := d.Relations[0]
			reverse.From, reverse.To = reverse.To, reverse.From
			d.Relations = append(d.Relations, reverse)
		}},
		{"self cycle", "cycle", func(d *catalog.Data) { d.Relations[0].To = "owner" }},
		{"invalid coverage", "coverage", func(d *catalog.Data) { d.Components[0].Coverage = "guessed" }},
		{"invalid level", "invalid component", func(d *catalog.Data) { d.Components[0].Level = "Widget" }},
		{"missing primary", "primary component", func(d *catalog.Data) { d.Examples[0].Primary = "other" }},
		{"unknown example use", "unknown component", func(d *catalog.Data) { d.Examples[0].Uses[0].Component = "missing" }},
		{"duplicate example", "duplicate example", func(d *catalog.Data) { d.Examples = append(d.Examples, d.Examples[0]) }},
		{"duplicate usage role", "usage roles", func(d *catalog.Data) { d.Examples[0].Uses[1].Role = "primary" }},
		{"bad example kind", "example kind", func(d *catalog.Data) { d.Examples[0].Kind = "other" }},
		{"unsafe evidence", "invalid source", func(d *catalog.Data) { d.Relations[0].Source.Path = "../private.go" }},
		{"control role", "needs a role", func(d *catalog.Data) { d.Relations[0].Role = "bad\x00role" }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := fixture()
			tc.change(&data)

			if _, err := catalog.New(data); err == nil || !strings.Contains(err.Error(), tc.message) {
				t.Fatalf("got %v; want %s", err, tc.message)
			}
		})
	}
}

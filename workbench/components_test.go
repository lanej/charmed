package workbench

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/lanej/charmed/catalog"
	"github.com/lanej/charmed/story"
)

func TestComponentGuideVerifiesEvidenceAndRetainsCatalogLinks(t *testing.T) {
	root, err := moduleRoot()
	if err != nil {
		t.Fatal(err)
	}

	entries := []ManifestEntry{}
	for _, d := range story.Definitions() {
		entries = append(entries, ManifestEntry{Definition: d, ExampleFrame: "consumer output for " + d.ID})
	}

	guide, err := buildComponentGuide(root, entries)
	if err != nil {
		t.Fatal(err)
	}

	replay := guide.find("diagram/trace-replay")
	if replay == nil || len(replay.Dependencies) != 1 || replay.Dependencies[0].Name != "State diagram" {
		t.Fatal("replay graph includes fixture controls")
	}

	if replay.ExampleFrame != "consumer output for diagram/trace-replay" || !strings.HasSuffix(replay.ExampleHref, "#diagram%2Ftrace-replay") {
		t.Fatal("preview must use the component's own consumer example")
	}

	state := guide.find("diagram/state-diagram")
	if state == nil || len(state.Examples) == 0 || !strings.Contains(state.Examples[0].Href, "molecules-diagram.html") {
		t.Fatal("canonical identity broke the existing catalog URL")
	}

	sources := sourceLinks{root: root, files: map[string]map[string]int{}}
	if _, err := sources.link(catalog.Source{Path: "collection/grid_inspector.go", Symbol: "GridInspector.Missing"}); err == nil {
		t.Fatal("nonexistent source evidence accepted")
	}

	if _, err := sources.link(catalog.Source{Path: filepath.ToSlash("collection/grid_inspector.go"), Symbol: "GridInspector.Grid"}); err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		if entry.ID == "diagram/trace-replay" && (entry.Component == nil || len(entry.ComponentExamples) != 2) {
			t.Fatal("catalog binding conflated consumer and fixture examples")
		}
	}
}

func TestComponentProjectionKeepsRepeatedRolesInOneExample(t *testing.T) {
	root, err := moduleRoot()
	if err != nil {
		t.Fatal(err)
	}

	graph, err := story.ComponentGraph()
	if err != nil {
		t.Fatal(err)
	}

	data := graph.Data()
	repeated := data.Examples[0].Uses[0]
	repeated.Role = "second grid"
	data.Examples[0].Uses = append(data.Examples[0].Uses, repeated)

	graph, err = catalog.New(data)
	if err != nil {
		t.Fatal(err)
	}

	entries := []ManifestEntry{}
	for _, d := range story.Definitions() {
		entries = append(entries, ManifestEntry{Definition: d})
	}

	guide, err := projectComponentGuide(root, entries, graph)
	if err != nil {
		t.Fatal(err)
	}

	examples := guide.find("collection/data-grid").Examples
	if len(examples) != 1 || len(examples[0].Uses) != 2 {
		t.Fatal("repeated component roles duplicated the example or lost a role")
	}
}

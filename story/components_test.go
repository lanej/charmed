package story_test

import (
	"testing"

	"github.com/lanej/charmed/catalog"
	"github.com/lanej/charmed/story"
)

func TestComponentGraphSeparatesReplayControlsFromLibraryDependencies(t *testing.T) {
	graph, err := story.ComponentGraph()
	if err != nil {
		t.Fatal(err)
	}

	dependencies := graph.Dependencies("diagram/trace-replay")
	if len(dependencies) != 1 || dependencies[0].To != "diagram/state-diagram" {
		t.Fatal("fixture controls became component dependencies")
	}

	example, ok := graph.Example("diagram/trace-replay/demo")
	if !ok || example.Kind != catalog.InteractiveExample {
		t.Fatal("interactive example identity missing")
	}

	radio := false
	for _, use := range example.Uses {
		radio = radio || use.Component == "input/radio-group"
	}

	if !radio {
		t.Fatal("fixture usage lost its radio group")
	}

	consumer, ok := graph.Example("diagram/trace-replay/usage")
	if !ok || consumer.Kind != catalog.ConsumerExample || len(consumer.Uses) != 2 {
		t.Fatal("consumer and fixture sources conflated")
	}

	matrix := graph.Dependencies("collection/status-matrix")
	if matrix[0].Kind != catalog.Specializes || matrix[0].To != "collection/data-grid" {
		t.Fatal("grid specialization lost")
	}

	users := graph.Dependents("collection/data-grid")
	if len(users) != 2 {
		t.Fatal("shared grid duplicated or not indexed")
	}
}

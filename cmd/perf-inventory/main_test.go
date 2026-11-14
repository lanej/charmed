package main

import (
	"testing"

	"github.com/lanej/charmed/story"
)

func TestInventoryCoversPublicViews(t *testing.T) {
	t.Chdir("../..")

	result, err := collect()
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Fixtures) != len(story.Definitions()) {
		t.Fatal("fixture omitted")
	}

	views, types := 0, 0

	for _, surface := range result.Surfaces {
		if surface.Kind == "view" {
			views++

			if len(surface.Benchmarks) == 0 {
				t.Fatalf("unmapped native view: %s", surface.Symbol)
			}
		}

		if surface.Kind == "type" {
			types++
		}
	}

	if views == 0 || types == 0 {
		t.Fatal("public methods or types omitted from inventory")
	}
}

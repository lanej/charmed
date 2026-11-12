package workbench

import (
	"strings"
	"testing"

	"github.com/lanej/charmed/story"
)

func TestAtomicNavigationAndGrouping(t *testing.T) {
	t.Parallel()

	var entries []ManifestEntry
	for _, definition := range story.Definitions() {
		entries = append(entries, ManifestEntry{Definition: definition})
	}

	levels := atomicNavigation(entries)
	if len(levels) != len(story.Levels()) {
		t.Fatal("navigation must expose every atomic level")
	}

	seen := make(map[string]bool)
	files := make(map[string]bool)

	for i, name := range story.Levels() {
		level := levels[i]
		if level.Name != name || level.File != strings.ToLower(name)+".html" {
			t.Fatalf("level %d = %+v, want a dedicated %s page", i, level, name)
		}

		count := 0

		for _, group := range level.Groups {
			if files[group.File] {
				t.Errorf("duplicate group page %s", group.File)
			}

			files[group.File] = true

			for _, entry := range group.Entries {
				if entry.Level != name || entry.Family != group.Family || seen[entry.ID] {
					t.Errorf("%s grouped under the wrong level or repeated", entry.ID)
				}

				if storyPage(entry) != group.File {
					t.Errorf("%s links to a different group page", entry.ID)
				}

				seen[entry.ID] = true
				count++
			}
		}

		if level.Count != count {
			t.Errorf("%s count = %d, want %d", name, level.Count, count)
		}
	}

	if len(seen) != len(entries) {
		t.Errorf("grouped %d of %d entries", len(seen), len(entries))
	}
}

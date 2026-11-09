package workbench

import (
	"strings"
	"testing"

	"github.com/lanej/charmed/story"
)

func TestAtomicNavigationAndGrouping(t *testing.T) {
	t.Parallel()

	links := atomicNavigation()
	if len(links) != len(story.Levels()) {
		t.Fatal("navigation must expose every atomic level")
	}

	var entries []ManifestEntry

	for _, definition := range story.Definitions() {
		entries = append(entries, ManifestEntry{Definition: definition})
	}

	seen := make(map[string]bool)

	for i, name := range story.Levels() {
		if links[i].Name != name {
			t.Fatalf("level %d = %s, want %s", i, links[i].Name, name)
		}

		page, anchor, _ := strings.Cut(links[i].Href, "#")
		found := false

		for _, level := range groupLevels(page, entries) {
			if level.Name != name {
				continue
			}

			found = level.Anchor == anchor

			for _, group := range level.Groups {
				for _, entry := range group.Entries {
					if entry.Level != name || entry.Family != group.Name || seen[entry.ID] {
						t.Errorf("%s grouped under the wrong level or repeated", entry.ID)
					}

					seen[entry.ID] = true
				}
			}
		}

		if !found {
			t.Errorf("navigation target %s is missing", links[i].Href)
		}
	}

	if len(seen) != len(entries) {
		t.Errorf("grouped %d of %d entries", len(seen), len(entries))
	}
}

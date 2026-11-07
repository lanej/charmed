package workbench_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lanej/charmed/workbench"
)

func TestNestedPagesRetainRoutesAndReplaceStaleExports(t *testing.T) {
	out := t.TempDir()
	directory := filepath.Join(out, "generated")
	stale := []string{"chart_timeseries.mdx", "atoms/old-name.mdx"}
	for _, name := range stale {
		path := filepath.Join(directory, name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("old page"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	if err := workbench.Generate(out); err != nil {
		t.Fatal(err)
	}
	for _, name := range stale {
		if _, err := os.Stat(filepath.Join(directory, name)); !os.IsNotExist(err) {
			t.Fatalf("stale page remains: %s (%v)", name, err)
		}
	}
	data, err := os.ReadFile(filepath.Join(directory, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var entries []workbench.ManifestEntry
	if err = json.Unmarshal(data, &entries); err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		path := filepath.Join(directory, filepath.FromSlash(entry.Page))
		page, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(page), "<Composition storyId=\""+entry.ID+"\" />") {
			t.Fatalf("%s has no composition section", entry.ID)
		}
		// Every generated import must resolve from its nested directory.
		for _, line := range strings.Split(string(page), "\n") {
			if !strings.HasPrefix(line, "import ") {
				continue
			}
			parts := strings.Split(line, "'")
			if len(parts) < 3 {
				t.Fatalf("invalid generated import: %s", line)
			}
			// The caller chooses the output directory; components sit beside generated.
			got := filepath.Clean(filepath.Join(filepath.Dir(path), filepath.FromSlash(parts[1])))
			want := filepath.Join(out, "components", filepath.Base(parts[1]))
			if got != want {
				t.Fatalf("%s import resolves to %s, want %s", entry.ID, got, want)
			}
		}
	}
}

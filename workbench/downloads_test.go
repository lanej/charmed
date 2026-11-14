package workbench_test

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertGuideDownloads(t *testing.T, directory string, pages map[string][]byte) map[string][]byte {
	t.Helper()

	index := readDownload(t, directory, "llms.txt")
	downloads := map[string][]byte{"llms.txt": index}

	for file, page := range pages {
		name := strings.TrimSuffix(file, ".html") + ".md"

		markdown := readDownload(t, directory, name)
		downloads[name] = markdown

		if len(bytes.TrimSpace(markdown)) == 0 {
			t.Errorf("%s has no Markdown content", file)
		}

		if !bytes.Contains(index, []byte("https://lanej.io/charmed/"+name)) {
			t.Errorf("index omits %s", name)
		}

		if !bytes.Contains(page, []byte(`type="text/markdown" href="`+name+`"`)) {
			t.Errorf("%s omits alternate Markdown link", file)
		}
	}

	chart := string(readDownload(t, directory, "atoms-chart.md"))
	for _, want := range []string{"```go", "Run this example", "Updates, focus", "Native preview", "go run ./examples/catalog chart/timeseries"} {
		if !strings.Contains(chart, want) {
			t.Errorf("complete chart Markdown omits %q", want)
		}
	}

	for _, name := range []string{"charmed", "charmed-contributing"} {
		assertSkillArchive(t, directory, name)

		archive := "skills/" + name + ".tar.gz"
		downloads[archive] = readDownload(t, directory, archive)
	}

	packages, err := os.ReadDir(filepath.Join(directory, "skills"))
	if err != nil {
		t.Fatal(err)
	}

	if len(packages) != 4 {
		t.Errorf("expected only two skill directories and two archives; got %v", packages)
	}

	for _, name := range []string{"builder", "go"} {
		for _, path := range []string{"skills/" + name, "skills/" + name + ".tar.gz"} {
			if _, statErr := os.Stat(filepath.Join(directory, path)); !errors.Is(statErr, os.ErrNotExist) {
				t.Errorf("internal skill published at %s: %v", path, statErr)
			}
		}

		if bytes.Contains(index, []byte("/skills/"+name+"/")) {
			t.Errorf("index advertises internal skill %s", name)
		}
	}

	return downloads
}

func readDownload(t *testing.T, directory, name string) []byte {
	t.Helper()

	root, err := os.OpenRoot(directory)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if closeErr := root.Close(); closeErr != nil {
			t.Error(closeErr)
		}
	})

	data, err := root.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}

	return data
}

func assertSkillArchive(t *testing.T, directory, name string) {
	t.Helper()

	data := readDownload(t, directory, "skills/"+name+".tar.gz")

	compressed, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if closeErr := compressed.Close(); closeErr != nil {
			t.Error(closeErr)
		}
	})

	archive := tar.NewReader(compressed)
	entries := map[string]bool{}

	for {
		header, nextErr := archive.Next()
		if errors.Is(nextErr, io.EOF) {
			break
		}

		if nextErr != nil {
			t.Fatal(nextErr)
		}

		if !strings.HasPrefix(header.Name, name+"/") || strings.Contains(header.Name, "..") || header.Typeflag != tar.TypeReg || entries[header.Name] {
			t.Fatalf("unsafe or duplicate skill entry: %+v", header)
		}

		entries[header.Name] = true

		contents, readErr := io.ReadAll(archive)
		if readErr != nil {
			t.Fatal(readErr)
		}

		source := readDownload(t, filepath.Join("..", ".agents", "skills"), header.Name)

		raw := readDownload(t, filepath.Join(directory, "skills"), header.Name)
		if !bytes.Equal(contents, source) || !bytes.Equal(contents, raw) {
			t.Errorf("archive/raw skill differs from source: %s", header.Name)
		}
	}

	if !entries[name+"/SKILL.md"] || len(entries) != 1 {
		t.Errorf("%s must contain its standalone SKILL.md only: %v", name, entries)
	}
}

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

	for _, name := range []string{"charmed", "builder", "go"} {
		assertSkillArchive(t, directory, name)

		archive := "skills/" + name + ".tar.gz"
		downloads[archive] = readDownload(t, directory, archive)
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

	for _, path := range map[string][]string{
		"charmed": {"SKILL.md"},
		"builder": {"SKILL.md", "references/standard.md", "references/construction.md", "references/instructions.md", "scripts/check_examples.py", "scripts/operations.py", "scripts/decisions.py", "scripts/utility.py"},
		"go":      {"SKILL.md", "references/correctness.md", "assets/check-go.sh", "assets/go-quality.py", "assets/go-quality.json", "assets/golangci.yml", "assets/go-quality.yml"},
	}[name] {
		if !entries[name+"/"+path] {
			t.Errorf("incomplete %s archive: missing %s", name, path)
		}
	}
}

package workbench

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lanej/charmed/story"
)

// exampleCode reads curated documentation snippets inside the module root.
// Root-relative access prevents a fixture definition from escaping the repository.
func exampleCode(root string, d story.Definition) (code string, err error) {
	if d.ExampleSource == "" {
		return sourceSymbol(filepath.Join(root, d.Source), d.Symbol)
	}

	if !filepath.IsLocal(d.ExampleSource) || !strings.HasPrefix(d.ExampleSource, "docs/snippets/") {
		return "", fmt.Errorf("invalid example source %q", d.ExampleSource)
	}

	directory, err := os.OpenRoot(root)
	if err != nil {
		return "", err
	}

	defer func() { err = errors.Join(err, directory.Close()) }()

	file, err := directory.Open(d.ExampleSource)
	if err != nil {
		return "", fmt.Errorf("open example %s: %w", d.ExampleSource, err)
	}

	defer func() { err = errors.Join(err, file.Close()) }()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("read example %s: %w", d.ExampleSource, err)
	}

	return string(data), nil
}

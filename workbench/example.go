package workbench

import (
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lanej/charmed/story"
)

type usageImport struct{ Name, Path string }

type usageExample struct {
	Code    string
	Imports []usageImport
}

// exampleCode reads compiled consumer examples inside the module root.
// Root-relative access prevents a fixture definition from escaping the repository.
func exampleCode(root string, d story.Definition) (code usageExample, err error) {
	if d.ExampleSource == "" {
		return code, fmt.Errorf("%s needs a consumer usage example", d.ID)
	}

	if !filepath.IsLocal(d.ExampleSource) || !strings.HasPrefix(d.ExampleSource, "docs/snippets/") {
		return code, fmt.Errorf("invalid example source %q", d.ExampleSource)
	}

	directory, err := os.OpenRoot(root)
	if err != nil {
		return code, err
	}

	defer func() { err = errors.Join(err, directory.Close()) }()

	file, err := directory.Open(d.ExampleSource)
	if err != nil {
		return code, fmt.Errorf("open example %s: %w", d.ExampleSource, err)
	}

	defer func() { err = errors.Join(err, file.Close()) }()

	data, err := io.ReadAll(file)
	if err != nil {
		return code, fmt.Errorf("read example %s: %w", d.ExampleSource, err)
	}

	return parseExample(d.ExampleSource, data)
}

func parseExample(filename string, data []byte) (usageExample, error) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filename, data, parser.ParseComments)
	if err != nil {
		return usageExample{}, err
	}

	var example usageExample

	var parts []string

	found := false

	for _, declaration := range file.Decls {
		if imports, ok := declaration.(*ast.GenDecl); ok && imports.Tok == token.IMPORT {
			for _, spec := range file.Imports {
				name, err := strconv.Unquote(spec.Path.Value)
				if err != nil {
					return example, err
				}

				alias := path.Base(strings.TrimSuffix(name, "/v2"))
				if spec.Name != nil {
					alias = spec.Name.Name
				}

				example.Imports = append(example.Imports, usageImport{Name: alias, Path: name})
			}

			continue
		}

		start, end := declaration.Pos(), declaration.End()
		if fn, ok := declaration.(*ast.FuncDecl); ok && strings.HasPrefix(fn.Name.Name, "Example") {
			if found || fn.Body == nil {
				return example, fmt.Errorf("%s must have one example function", filename)
			}

			found = true
			start, end = fn.Body.Lbrace+1, fn.Body.Rbrace
		}

		part := string(data[fset.Position(start).Offset:fset.Position(end).Offset])

		formatted, err := format.Source([]byte(strings.TrimSpace(part)))
		if err != nil {
			return example, fmt.Errorf("format usage example %s: %w", filename, err)
		}

		parts = append(parts, strings.TrimSpace(string(formatted)))
	}

	if !found {
		return example, fmt.Errorf("%s has no example function", filename)
	}

	example.Code = strings.Join(parts, "\n\n") + "\n"

	return example, nil
}

package snippets

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionFiles() (frame string, err error) {
	directory, err := os.MkdirTemp("", "charmed-example-")
	if err != nil {
		return "", err
	}
	defer func() { err = errors.Join(err, os.RemoveAll(directory)) }()

	for _, name := range []string{"main.go", "README.md", "theme.go"} {
		if err := os.WriteFile(filepath.Join(directory, name), []byte("// Sample file\n"), 0o600); err != nil {
			return "", err
		}
	}

	picker := collection.NewFilePicker(directory, theme.New(theme.Dark))
	picker.AllowedTypes = []string{".go", ".md"}
	picker.ShowPermissions, picker.ShowSize = false, false
	picker.SetHeight(6)
	// A running Bubble Tea program executes this command and sends its message.
	picker, _ = picker.Update(picker.Init()())

	return picker.View(), nil
}

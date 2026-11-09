package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionFiles() string {
	picker := collection.NewFilePicker(".", theme.New(theme.Dark))
	picker.AllowedTypes = []string{".go", ".md"}
	// Return picker.Init() from Init to load the directory. Until then, the view is empty.
	// Retain picker.Update(msg)'s model and command; use picker.DidSelectFile(msg) on Enter.
	return picker.View()
}

package snippets

import (
	"charm.land/bubbles/v2/tree"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionTree() string {
	root := tree.Root("charmed").Child(
		tree.Root("components").Child("input", "collection", "chart"),
		tree.Root("workbench").Child("native", "browser", "MDX"),
	)
	treeView := collection.NewTree(root, 48, 12, theme.New(theme.Dark))

	return treeView.View()
}

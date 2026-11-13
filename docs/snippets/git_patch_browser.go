package snippets

import (
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleGitPatchBrowser() (string, error) {
	patch := gitui.PatchObservation{Available: true, Files: []gitui.FileDiff{
		{FileID: "config", Section: gitui.Staged, Path: "config.yaml", State: gitui.TextDiff, Hunks: []gitui.Hunk{
			{Header: "@@ -1 +1 @@", Lines: []gitui.DiffLine{{Kind: gitui.DeletedLine, OldLine: 1, Text: "replicas: 2"}, {Kind: gitui.AddedLine, NewLine: 1, Text: "replicas: 3"}}},
		}},
		{FileID: "logo", Section: gitui.Unstaged, Path: "logo.png", State: gitui.BinaryDiff},
	}}

	browser, err := gitui.NewPatchBrowser(patch, layout.Size{Width: 80, Height: 14}, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}
	// FocusDetail controls routing; SetPatch deep-copies and preserves surviving identity.
	// Available=true with no files is empty; the zero observation is unavailable.
	return browser.View(), nil
}

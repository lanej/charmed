package gitui_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestPatchBrowserInventorySelectionAndOwnership(t *testing.T) {
	m, err := gitui.NewPatchBrowser(gitui.PatchObservation{Reason: "not captured"}, layout.Size{Width: 100, Height: 12}, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(m.View()), "Patch unavailable: not captured") {
		t.Fatal("unknown inventory presented as empty")
	}

	if err := m.SetPatch(gitui.PatchObservation{Available: true}); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(m.View()), "No files changed") {
		t.Fatal("known empty patch lost")
	}

	patch := gitui.PatchObservation{Available: true, Files: []gitui.FileDiff{
		{FileID: "a", Path: "a.go", Section: gitui.Staged, State: gitui.TextDiff, Hunks: []gitui.Hunk{{Header: "@@", Lines: []gitui.DiffLine{{Kind: gitui.AddedLine, Text: "original"}}}}},
		{FileID: "a", Path: "a.go", Section: gitui.Unstaged, State: gitui.BinaryDiff},
	}}

	patch.Files[0].OriginalPath = "old.go"
	if err := m.SetPatch(patch); err != nil {
		t.Fatal(err)
	}

	if strings.Contains(ansi.Strip(m.View()), "Enter select") {
		t.Fatal("instructions were not opt-in")
	}

	m.ShowHelp = true
	m.Resize(layout.Size{Width: 100, Height: 12})

	if !strings.Contains(ansi.Strip(m.View()), "Enter select") {
		t.Fatal("opt-in instructions missing")
	}

	patch.Files[0].Hunks[0].Lines[0].Text = "mutated"

	m.Select("a", gitui.Staged)

	if m.View() != ansi.Strip(m.View()) || strings.Contains(m.View(), "→") {
		t.Fatal("native Git portable presentation bypassed")
	}

	if !strings.Contains(ansi.Strip(m.View()), "original") || strings.Contains(ansi.Strip(m.View()), "mutated") {
		t.Fatal("nested diff alias retained")
	}

	if !m.Select("a", gitui.Unstaged) {
		t.Fatal("section identity lost")
	}

	if selected, ok := m.Selected(); !ok || selected.Section != gitui.Unstaged || !strings.Contains(ansi.Strip(m.View()), "Binary content changed") {
		t.Fatal("wrong section diff")
	}

	patch.Files = append(patch.Files, patch.Files[0])

	before := m.View()
	if err := m.SetPatch(patch); err == nil || m.View() != before {
		t.Fatal("invalid replacement mutated browser")
	}

	m.Update(tea.KeyPressMsg{Code: '/', Text: "/"})
	m.Update(tea.KeyPressMsg{Code: 'z', Text: "z"})

	if strings.Contains(ansi.Strip(m.View()), "No files changed") {
		t.Fatal("filtering hid known patch inventory")
	}
}

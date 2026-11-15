package main

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func TestConsumerReviewRoutesFocusMarksAndResize(t *testing.T) {
	m, err := newModel(false)
	if err != nil {
		t.Fatal(err)
	}

	m.document.Select("tests")
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m.Update(tea.KeyPressMsg{Code: 'r'})

	if !m.document.DocumentFocused() || !m.review.Reviewed("tests") || m.review.Reviewed("routing") {
		t.Fatal("consumer did not review the highlighted file after switching focus")
	}

	m.Update(tea.WindowSizeMsg{Width: 40, Height: 18})
	view := m.View()

	if !view.AltScreen || view.MouseMode != tea.MouseModeCellMotion || lipgloss.Width(view.Content) != 40 || lipgloss.Height(view.Content) != 18 {
		t.Fatal("consumer lost terminal capabilities or resize allocation")
	}

	if m.document.SelectedID() != "tests" {
		t.Fatal("resize lost the selected review file")
	}
}

func TestConsumerNonFileDocument(t *testing.T) {
	m, err := newModel(true)
	if err != nil {
		t.Fatal(err)
	}

	m.document.Select("evidence")
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if m.review != nil || !m.document.Collapsed("evidence") {
		t.Fatal("non-file consumer did not fold arbitrary content")
	}

	_, cmd := m.Update(tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl})
	if cmd == nil {
		t.Fatal("consumer did not handle quit")
	}
}

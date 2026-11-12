package main

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func TestQuickstartRetainsFocusEditsAndResize(t *testing.T) {
	app := &model{field: input.NewText(input.TextOptions{Value: "prod", Width: 36, Theme: theme.New(theme.Dark)})}
	if app.Init() == nil || !app.field.Focused() {
		t.Fatal("initialization must focus the retained field and return its command")
	}

	app.field.CursorEnd()
	_, _ = app.Update(tea.KeyPressMsg{Code: 'x', Text: "x"})
	_, _ = app.Update(tea.WindowSizeMsg{Width: 24, Height: 10})

	if app.field.Value() != "prodx" || !app.field.Focused() {
		t.Fatal("editing and resizing must preserve the field value and focus")
	}

	if app.field.Width() >= 24 {
		t.Fatal("the input must reserve room for its prompt and cursor")
	}
}

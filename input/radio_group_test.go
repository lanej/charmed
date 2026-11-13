package input_test

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func TestRadioGroupFocusIsSeparateFromSelection(t *testing.T) {
	group, err := input.NewRadioGroup([]input.RadioOption{{ID: "a", Label: "A"}, {ID: "b", Label: "B", DisabledReason: "Unavailable"}, {ID: "c", Label: "C"}}, "a", theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	group.Focus()

	if cmd := group.Update(tea.KeyPressMsg{Code: tea.KeyDown}); cmd != nil {
		t.Fatal("arrow emitted activation")
	}

	if id, _ := group.Selected(); id != "a" {
		t.Fatal("focus committed selection")
	}

	cmd := group.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("selection command missing")
	}

	event, ok := cmd().(input.RadioSelected)
	if !ok || event.ID != "c" {
		t.Fatal("disabled option was not skipped")
	}

	if cmd := group.Update(tea.KeyPressMsg{Code: tea.KeySpace}); cmd != nil {
		t.Fatal("already selected option activated")
	}

	group.Blur()
	group.Update(tea.KeyPressMsg{Code: tea.KeyUp})

	if !group.Select("a") || group.Select("b") {
		t.Fatal("programmatic selection invalid")
	}

	if _, duplicateErr := input.NewRadioGroup([]input.RadioOption{{ID: "x"}, {ID: "x"}}, "", theme.New(theme.Mono)); duplicateErr == nil {
		t.Fatal("duplicate IDs accepted")
	}

	empty, err := input.NewRadioGroup(nil, "", theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	empty.Focus()

	if cmd := empty.Update(tea.KeyPressMsg{Code: tea.KeyEnter}); cmd != nil {
		t.Fatal("empty group activated")
	}
}

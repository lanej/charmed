package input_test

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func TestCheckboxActivationRetainsIntentAndFocusBoundary(t *testing.T) {
	tmode := theme.New(theme.Mono)
	m := input.NewCheckbox("all", "All services", input.Mixed, tmode)

	press := tea.KeyPressMsg{Code: tea.KeySpace}
	if _, cmd := m.Update(press); cmd != nil {
		t.Fatal("unfocused control activated")
	}

	m.Focus()

	next, cmd := m.Update(press)
	if cmd == nil || next.State != input.Checked || m.State != input.Mixed {
		t.Fatal("mixed-to-checked transition mutated original or failed")
	}

	if event, ok := cmd().(input.CheckChanged); !ok || event.ID != "all" || event.State != input.Checked {
		t.Fatal("wrong check event")
	}

	next, cmd = next.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil || next.State != input.Unchecked {
		t.Fatal("checked did not clear")
	}

	next.DisabledReason = "Read-only"
	if _, cmd = next.Update(press); cmd != nil || next.Focused() {
		t.Fatal("disabled control activated or claimed focus")
	}

	if !strings.Contains(ansi.Strip(next.View()), "[DISABLED] Read-only") {
		t.Fatal("disabled reason lost")
	}

	next.DisabledReason = ""
	next.Blur()

	if _, cmd = next.Update(press); cmd != nil {
		t.Fatal("blur did not stop activation")
	}

	next.Focus()

	next.Keys = key.NewBinding(key.WithKeys("x"))
	if _, cmd = next.Update(press); cmd != nil {
		t.Fatal("custom binding lost")
	}

	if _, cmd = next.Update(tea.KeyPressMsg{Code: 'x', Text: "x"}); cmd == nil {
		t.Fatal("custom binding failed")
	}

	next.State = input.CheckState(99)
	if _, cmd = next.Update(tea.KeyPressMsg{Code: 'x', Text: "x"}); cmd != nil || !strings.Contains(ansi.Strip(next.View()), "[?]") {
		t.Fatal("invalid state silently normalized")
	}
}

func TestRadioToggleAndButtonEmitTypedMessages(t *testing.T) {
	palette := theme.New(theme.Dark)
	enter := tea.KeyPressMsg{Code: tea.KeyEnter}
	radio := input.NewRadio("prod", "Production", false, palette)
	radio.Focus()

	radio, cmd := radio.Update(enter)
	if cmd == nil || !radio.Selected {
		t.Fatal("radio did not select")
	}

	if event, ok := cmd().(input.RadioSelected); !ok || event.ID != "prod" {
		t.Fatal("radio identity lost")
	}

	if _, cmd = radio.Update(enter); cmd != nil {
		t.Fatal("selected radio emitted redundant selection")
	}

	toggle := input.NewToggle("live", "Live", false, palette)
	toggle.Focus()

	toggle, cmd = toggle.Update(enter)
	if cmd == nil || !toggle.On {
		t.Fatal("toggle did not turn on")
	}

	if event, ok := cmd().(input.Toggled); !ok || !event.On || event.ID != "live" {
		t.Fatal("toggle event lost")
	}

	toggle, cmd = toggle.Update(enter)
	if cmd == nil || toggle.On {
		t.Fatal("toggle did not turn off")
	}

	button := input.NewButton("retry", "Retry", palette)
	button.Focus()

	button, cmd = button.Update(enter)
	if cmd == nil || button.Busy {
		t.Fatal("button lost intent or guessed operation state")
	}

	if event, ok := cmd().(input.Activated); !ok || event.ID != "retry" {
		t.Fatal("wrong action identity")
	}

	button.Busy = true
	if _, cmd = button.Update(enter); cmd != nil || !strings.Contains(ansi.Strip(button.View()), "BUSY") {
		t.Fatal("busy button activated or lost state")
	}

	button.Busy = false

	button.Disabled = true
	if _, cmd = button.Update(enter); cmd != nil {
		t.Fatal("disabled button activated")
	}
}

func TestControlMarksKeepGeometryAndEscapeLabels(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		widths := map[int]bool{}

		for _, state := range []input.CheckState{input.Unchecked, input.Checked, input.Mixed} {
			m := input.NewCheckbox("one", "界e\u0301", state, palette)
			widths[ansi.StringWidth(m.View())] = true
			m.Focus()
			widths[ansi.StringWidth(m.View())] = true
		}

		if len(widths) != 1 {
			t.Fatalf("%s focus or check state changed geometry", mode)
		}

		off, on := input.NewToggle("t", "Live", false, palette), input.NewToggle("t", "Live", true, palette)
		if ansi.StringWidth(off.View()) != ansi.StringWidth(on.View()) {
			t.Fatal("toggle state reflowed")
		}

		button := input.NewButton("b", "Retry\x1b[2J", palette)
		if !strings.Contains(ansi.Strip(button.View()), `Retry\u001b[2J`) {
			t.Fatal("untrusted label was interpreted")
		}
	}
}

func TestControlsUseUnicodeIndicatorsWithExplicitASCIIFallbacks(t *testing.T) {
	palette := theme.New(theme.Dark)
	for _, tc := range []struct {
		state input.CheckState
		glyph string
	}{{input.Unchecked, "☐"}, {input.Checked, "☑"}, {input.Mixed, "▣"}} {
		control := input.NewCheckbox("id", "Services", tc.state, palette)
		if !strings.Contains(ansi.Strip(control.View()), tc.glyph) {
			t.Fatal("checkbox lost its Unicode state")
		}
	}

	radio := input.NewRadio("id", "Production", true, palette)
	if !strings.Contains(ansi.Strip(radio.View()), "◉") {
		t.Fatal("radio selection lost its distinct circle")
	}

	off := input.NewToggle("id", "Live", false, palette)

	on := input.NewToggle("id", "Live", true, palette)
	if !strings.Contains(ansi.Strip(off.View()), "○── OFF") || !strings.Contains(ansi.Strip(on.View()), "━━● ON") {
		t.Fatal("switch track lost its thumb position")
	}

	button := input.NewButton("id", "Retry", palette)

	button.Icon = theme.GlyphRefresh
	if !strings.Contains(ansi.Strip(button.View()), "⟨ ↻ Retry ⟩") {
		t.Fatal("action icon was lost")
	}

	button.DisabledReason = "Read-only"
	if !strings.Contains(ansi.Strip(button.View()), "⊘ disabled Read-only") {
		t.Fatal("disabled state lost its icon or explanation")
	}

	button.Theme.ASCII = true
	if !strings.Contains(ansi.Strip(button.View()), "[ Retry ] [DISABLED] Read-only") {
		t.Fatal("explicit ASCII policy lost its readable fallback")
	}
}

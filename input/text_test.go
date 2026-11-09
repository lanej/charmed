package input_test

import (
	"image/color"
	"reflect"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func TestDisabledInputDistinguishesStateWithoutMutatingNativeModel(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		model := input.NewText(input.TextOptions{Value: "production", Width: 30, Theme: palette})
		model.Focus()
		model.SetVirtualCursor(false)
		model.Disabled = true
		before := model.Styles()
		plain := ansi.Strip(model.View())

		if !strings.HasPrefix(plain, palette.Glyph(theme.GlyphDisabled)+" ") || model.Focused() || model.Cursor() != nil {
			t.Fatalf("%s disabled input lost its state cue: %q", mode, plain)
		}

		updated, _ := model.Update(tea.KeyPressMsg{Code: 'x', Text: "x"})
		if updated.Value() != "production" || model.Prompt != "> " || !reflect.DeepEqual(before, model.Styles()) {
			t.Fatal("disabled rendering or input changed native state")
		}

		model.Disabled = false
		if !model.Focused() || model.Cursor() == nil || !strings.HasPrefix(ansi.Strip(model.View()), "> ") {
			t.Fatal("reenabling lost native focus or prompt")
		}

		model.Prompt = "Find: "
		model.SetValue(strings.Repeat("wide value ", 10))
		model.Resize(30)

		model.Disabled = true
		if !strings.HasPrefix(ansi.Strip(model.View()), "Find: ") || ansi.StringWidth(model.View()) > 30 {
			t.Fatal("disabled view lost custom prompt or width")
		}
	}
}

func TestTextareaActiveLineUsesThemeSurface(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		model := input.NewArea(40, 6, palette)
		model.SetValue("First\nSecond")
		model.Focus()

		found := false

		for _, line := range uv.NewStyledString(model.View()).Lines(ansi.GraphemeWidth) {
			for _, cell := range line {
				if cell.Content == "S" {
					found = true

					if cell.Style.Bg == nil || color.NRGBAModel.Convert(cell.Style.Bg) != color.NRGBAModel.Convert(palette.Surface) {
						t.Fatalf("%s active textarea line lost the theme surface", mode)
					}
				}
			}
		}

		if !found {
			t.Fatal("textarea lost active-line content")
		}
	}
}

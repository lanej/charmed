package input_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func TestFieldWrapsFeedbackToTheNativeInputWidth(t *testing.T) {
	t.Parallel()

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		palette := theme.New(mode)
		help := "Use a unique environment name with at least three characters."
		field := input.Field{
			Label: "Environment name", Help: help, Theme: palette,
			Input: input.NewText(input.TextOptions{Value: "production", Width: 24, Theme: palette}),
		}

		for _, width := range []int{24, 18} {
			field.Input.Resize(width)
			view := ansi.Strip(field.View())

			for _, line := range strings.Split(view, "\n") {
				if ansi.StringWidth(line) > width {
					t.Fatalf("field escaped %d columns: %q", width, line)
				}
			}

			if !strings.Contains(strings.Join(strings.Fields(view), " "), help) {
				t.Fatal("wrapping discarded help text")
			}
		}

		message := "Enter at least three characters before submitting."
		field.Input.Err = errors.New(message)
		view := ansi.Strip(field.View())

		if !strings.Contains(view, palette.Glyph(theme.GlyphError)) || strings.Contains(view, help) {
			t.Fatal("validation did not use the error glyph and replace help")
		}

		if !strings.Contains(strings.Join(strings.Fields(view), " "), message) || field.Input.Value() != "production" {
			t.Fatal("validation rendering discarded the error or changed the input")
		}

		for _, line := range strings.Split(view, "\n") {
			if ansi.StringWidth(line) > 18 {
				t.Fatalf("validation escaped the resized input: %q", line)
			}
		}
	}
}

func TestUnlabeledFieldDoesNotAddAnEmptyHeading(t *testing.T) {
	t.Parallel()

	field := input.Field{Theme: theme.New(theme.Dark), Input: input.NewText(input.TextOptions{Width: 20, Theme: theme.New(theme.Dark)})}
	if field.View() != field.Input.View() {
		t.Fatal("an empty field label added decoration")
	}
}

func TestTextResizePreservesNativeValueCursorAndCustomPrompt(t *testing.T) {
	t.Parallel()

	field := input.NewText(input.TextOptions{Value: "production-environment", Width: 30, Theme: theme.New(theme.Dark)})
	field.Prompt = "env> "

	for _, position := range []int{0, 5, len(field.Value())} {
		field.SetCursor(position)

		for _, width := range []int{30, 12, 24, 8} {
			field.Resize(width)

			if ansi.StringWidth(field.View()) != width {
				t.Fatalf("cursor %d escaped %d columns: %q", position, width, ansi.Strip(field.View()))
			}

			if field.Value() != "production-environment" || field.Position() != position {
				t.Fatal("resize changed native editing state")
			}
		}
	}
}

package theme_test

import (
	"image/color"
	"math"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

func luminance(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	linear := func(v uint32) float64 {
		x := float64(v) / 65535
		if x <= .04045 {
			return x / 12.92
		}

		return math.Pow((x+.055)/1.055, 2.4)
	}

	return .2126*linear(r) + .7152*linear(g) + .0722*linear(b)
}

func TestDimTokensRemainReadableAndSeriesDoNotAliasSemanticRoles(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast} {
		palette := theme.New(mode)

		background := luminance(palette.Background)
		selection, foreground := luminance(palette.SelectionBackground), luminance(palette.Foreground)

		if (max(selection, foreground)+.05)/(min(selection, foreground)+.05) < 4.5 {
			t.Fatalf("%s selection surface makes nested plain text unreadable", mode)
		}

		for _, dim := range []color.Color{palette.DimForeground, palette.DimSuccess, palette.DimInfo, palette.DimWarning, palette.DimDanger, palette.DimUnknown, palette.DimInactive} {
			value := luminance(dim)

			ratio := (max(value, background) + .05) / (min(value, background) + .05)
			if ratio < 4.5 {
				t.Fatalf("%s dim token contrast %.2f is below 4.5:1", mode, ratio)
			}
		}

		for i, series := range palette.Series {
			for _, semantic := range []color.Color{palette.Accent, palette.Success, palette.Info, palette.Warning, palette.Danger, palette.Unknown, palette.Inactive} {
				if color.NRGBAModel.Convert(series) == color.NRGBAModel.Convert(semantic) {
					t.Fatalf("%s series %d aliases a semantic token", mode, i)
				}
			}
		}
	}
}

func TestFocusKeepsControlGeometry(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		if mode == theme.Mono && !palette.Focus().GetReverse() {
			t.Fatal("monochrome focus lost its inverse highlight")
		}

		for _, value := range []string{"Value", "界", ""} {
			normal, focus := palette.Control().Render(value), palette.Focus().Render(value)
			if ansi.Strip(normal) != " "+value+" " || ansi.Strip(focus) != ansi.Strip(normal) {
				t.Fatalf("%s focus changed the reserved inset: %q / %q", mode, normal, focus)
			}

			if ansi.Strip(palette.Selection().Render(value)) != ansi.Strip(normal) {
				t.Fatal("collection selection changed control geometry")
			}
		}
	}
}

func TestCustomThemeSelectionWithoutNewTokensKeepsAVisibleCue(t *testing.T) {
	custom := theme.Theme{Mode: theme.Dark, Foreground: lipgloss.Color("#eeeeee"), Background: lipgloss.Color("#101722")}
	if !custom.Selection().GetUnderline() || ansi.Strip(custom.Selection().Render("Value")) != " Value " {
		t.Fatal("custom theme selection lost its cue or reserved inset")
	}
}

func TestNamedGlyphFallbacksKeepStateAndWidth(t *testing.T) {
	unicode, ascii := theme.New(theme.Dark), theme.New(theme.Mono)
	for _, pair := range []struct {
		glyph          theme.Glyph
		unicode, ascii string
	}{
		{theme.GlyphOK, "✓", "+"},
		{theme.GlyphWarning, "!", "!"},
		{theme.GlyphError, "×", "x"},
		{theme.GlyphUnknown, "?", "?"},
		{theme.GlyphOff, "○", "-"},
		{theme.GlyphInfo, "i", "i"},
		{theme.GlyphSelected, "●", "*"},
		{theme.GlyphUnselected, "○", "o"},
		{theme.GlyphExpanded, "▾", "v"},
		{theme.GlyphCollapsed, "▸", ">"},
		{theme.GlyphUp, "↑", "^"},
		{theme.GlyphDown, "↓", "v"},
		{theme.GlyphLeft, "←", "<"},
		{theme.GlyphRight, "→", ">"},
		{theme.GlyphLoading, "…", "."},
		{theme.GlyphUnavailable, "—", "-"},
	} {
		if unicode.Glyph(pair.glyph) != pair.unicode || ascii.Glyph(pair.glyph) != pair.ascii {
			t.Fatalf("incorrect representations for %s", pair.glyph)
		}

		if lipgloss.Width(pair.unicode) != 1 || lipgloss.Width(pair.ascii) != 1 {
			t.Fatalf("%s does not fit the one-cell icon slot", pair.glyph)
		}
	}
}

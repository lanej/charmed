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

		for _, dim := range []color.Color{palette.DimForeground, palette.DimSuccess, palette.DimInfo, palette.DimWarning, palette.DimDanger, palette.DimUnknown, palette.DimInactive, palette.DimPending, palette.DimRunning, palette.DimPaused, palette.DimTrace, palette.DimDebug, palette.DimFatal} {
			value := luminance(dim)

			ratio := (max(value, background) + .05) / (min(value, background) + .05)
			if ratio < 4.5 {
				t.Fatalf("%s dim token contrast %.2f is below 4.5:1", mode, ratio)
			}
		}

		for i, series := range palette.Series {
			for _, semantic := range []color.Color{palette.Accent, palette.Success, palette.Info, palette.Warning, palette.Danger, palette.Unknown, palette.Inactive, palette.Pending, palette.Running, palette.Paused} {
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
		{theme.GlyphPending, "◷", "p"},
		{theme.GlyphRunning, "▶", ">"},
		{theme.GlyphPaused, "‖", "="},
		{theme.GlyphDisabled, "⊘", "/"},
		{theme.GlyphConnected, "●", "+"},
		{theme.GlyphDisconnected, "○", "-"},
		{theme.GlyphSearch, "⌕", "/"},
		{theme.GlyphFilter, "▽", "f"},
		{theme.GlyphSort, "↕", "s"},
		{theme.GlyphRefresh, "↻", "r"},
		{theme.GlyphLink, "↗", "@"},
		{theme.GlyphExpand, "⤢", "+"},
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
		{theme.GlyphLog, "≡", "L"},
		{theme.GlyphTrace, "⋈", "T"},
		{theme.GlyphSpan, "─", "S"},
		{theme.GlyphUpstream, "↑", "^"},
		{theme.GlyphDownstream, "↓", "v"},
		{theme.GlyphScrollTrack, "│", "|"},
		{theme.GlyphScrollThumb, "┃", "#"},
	} {
		if unicode.Glyph(pair.glyph) != pair.unicode || ascii.Glyph(pair.glyph) != pair.ascii {
			t.Fatalf("incorrect representations for %s", pair.glyph)
		}

		if lipgloss.Width(pair.unicode) != 1 || lipgloss.Width(pair.ascii) != 1 {
			t.Fatalf("%s does not fit the one-cell icon slot", pair.glyph)
		}
	}
}

func TestCustomBorderAndSelectionTokensRemainIndependent(t *testing.T) {
	palette := theme.New(theme.Dark)
	palette.FocusBorder = lipgloss.Color("#abcdef")
	palette.SelectionForeground = lipgloss.Color("#fedcba")

	if palette.BorderStyle(true).GetForeground() != palette.FocusBorder || palette.BorderStyle(false).GetForeground() != palette.Border {
		t.Fatal("frame colors lost their normal/focused roles")
	}

	if palette.Selection().GetForeground() != palette.SelectionForeground || palette.Focus().GetBackground() != palette.Accent {
		t.Fatal("collection selection and inverse focus tokens were conflated")
	}

	palette.FocusBorder = nil
	palette.SelectionForeground = nil

	if palette.BorderStyle(true).GetForeground() != palette.Accent || palette.Selection().GetForeground() != palette.Foreground {
		t.Fatal("partial custom theme lost accent/text fallbacks")
	}
}

func TestSequentialScalesOrderIntensityIndependentlyOfCategories(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		if len(palette.Sequential) < 2 {
			t.Fatal("heatmap needs an ordered scale")
		}

		for i := 1; i < len(palette.Sequential); i++ {
			before, after := luminance(palette.Sequential[i-1]), luminance(palette.Sequential[i])
			if (mode == theme.Light && after >= before) || (mode != theme.Light && after <= before) {
				t.Fatalf("%s heatmap intensity is not monotonic", mode)
			}
		}
	}
}

func TestPermissionColorsRemainReadableOnContentAndSelection(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast} {
		palette := theme.New(mode)
		roles := palette.Permissions

		for _, foreground := range []color.Color{roles.Read, roles.Write, roles.Execute, roles.WriteExecute, roles.None, roles.Unknown} {
			for _, background := range []color.Color{palette.Background, palette.Surface, palette.SelectionBackground} {
				fg, bg := luminance(foreground), luminance(background)
				if (max(fg, bg)+.05)/(min(fg, bg)+.05) < 4.5 {
					t.Fatalf("%s permission color is unreadable on a supported surface", mode)
				}
			}
		}
	}
}

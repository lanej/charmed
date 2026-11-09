package text_test

import (
	"image/color"
	"io/fs"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestPermissionColorsFollowCapabilitiesAndCustomTokens(t *testing.T) {
	palette := theme.New(theme.Dark)
	palette.Permissions.Write = lipgloss.Color("#abcdef")
	roles := palette.Permissions

	for _, example := range []struct {
		mode fs.FileMode
		want color.Color
	}{
		{0o000, roles.None},
		{0o100, roles.Execute},
		{0o200, roles.Write},
		{0o300, roles.WriteExecute},
		{0o400, roles.Read},
		{0o500, roles.Execute},
		{0o600, roles.Write},
		{0o700, roles.WriteExecute},
	} {
		p := text.Permissions{Mode: example.mode, Available: true}
		group := p.Groups(palette)[0]
		cells := uv.NewStyledString(group).Lines(ansi.GraphemeWidth)

		if len(cells) != 1 || len(cells[0]) != 1 || cells[0][0].Style.Fg == nil {
			t.Fatalf("mode %03o lost its colored one-cell glyph", example.mode)
		}

		if color.NRGBAModel.Convert(cells[0][0].Style.Fg) != color.NRGBAModel.Convert(example.want) {
			t.Fatalf("mode %03o has the wrong capability color", example.mode)
		}
	}

	palette.Permissions = theme.PermissionColors{}
	group := (text.Permissions{Mode: 0o400, Available: true}).Groups(palette)[0]
	cells := uv.NewStyledString(group).Lines(ansi.GraphemeWidth)

	if color.NRGBAModel.Convert(cells[0][0].Style.Fg) != color.NRGBAModel.Convert(palette.Foreground) {
		t.Fatal("custom theme without permission tokens lost its text fallback")
	}
}

func TestPermissionGroupsPreserveAllModeCombinations(t *testing.T) {
	labels := [8]string{"---", "--x", "-w-", "-wx", "r--", "r-x", "rw-", "rwx"}
	glyphs := [8]string{"∅", "▖", "▝", "▞", "▘", "▌", "▀", "▛"}
	unicode, ascii := theme.New(theme.Dark), theme.New(theme.Mono)

	for owner := range labels {
		for group := range labels {
			for other := range labels {
				p := text.Permissions{Mode: fs.FileMode(owner*64 + group*8 + other), Available: true}
				if p.Label() != labels[owner]+labels[group]+labels[other] {
					t.Fatalf("mode %03o lost a permission bit: %s", p.Mode, p.Label())
				}

				for palette, want := range map[bool][3]string{
					false: {glyphs[owner], glyphs[group], glyphs[other]},
					true:  {labels[owner], labels[group], labels[other]},
				} {
					themeValue := unicode
					if palette {
						themeValue = ascii
					}

					got := p.Groups(themeValue)
					for i := range got {
						if ansi.Strip(got[i]) != want[i] {
							t.Fatalf("mode %03o group %d: %q != %q", p.Mode, i, got[i], want[i])
						}

						width := 1
						if palette {
							width = 3
						}

						if ansi.StringWidth(got[i]) != width {
							t.Fatalf("permission group changed its %d-cell allocation", width)
						}
					}
				}
			}
		}
	}
}

func TestPermissionSpecialBitsKeepExecutionDistinction(t *testing.T) {
	for _, example := range []struct {
		mode     fs.FileMode
		label    string
		modifier string
	}{
		{0o755 | fs.ModeSetuid, "rwsr-xr-x", "setuid"},
		{0o644 | fs.ModeSetuid, "rwSr--r--", "setuid"},
		{0o755 | fs.ModeSetgid, "rwxr-sr-x", "setgid"},
		{0o644 | fs.ModeSetgid, "rw-r-Sr--", "setgid"},
		{0o777 | fs.ModeSticky, "rwxrwxrwt", "sticky"},
		{0o666 | fs.ModeSticky, "rw-rw-rwT", "sticky"},
		{0o777 | fs.ModeSetuid | fs.ModeSetgid | fs.ModeSticky, "rwsrwsrwt", "setuid setgid sticky"},
	} {
		p := text.Permissions{Mode: example.mode, Available: true}
		if p.Label() != example.label {
			t.Fatalf("special mode %v: %q != %q", p.Mode, p.Label(), example.label)
		}

		for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
			compact := ansi.Strip(p.Compact(theme.New(mode)))
			if !strings.HasSuffix(compact, example.modifier) {
				t.Fatalf("%s compact view hides special bits: %q", mode, compact)
			}
		}
	}
}

func TestPermissionUnknownNeverLooksLikeNoneOrInventsSpecialBits(t *testing.T) {
	unknown := text.Permissions{Mode: 0o777 | fs.ModeSetuid}

	none := text.Permissions{Available: true}
	if unknown.Label() != "?????????" || none.Label() != "---------" {
		t.Fatal("known empty mode and unknown mode were conflated")
	}

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)

		wantUnknown, wantNone := "? ? ?", "∅ ∅ ∅"
		if palette.ASCII {
			wantUnknown, wantNone = "??? ??? ???", "--- --- ---"
		}

		if ansi.Strip(unknown.Compact(palette)) != wantUnknown || ansi.Strip(none.Compact(palette)) != wantNone {
			t.Fatalf("%s conflates missing data and absent bits", mode)
		}
	}
}

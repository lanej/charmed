package text

import (
	"image/color"
	"io/fs"
	"strings"

	"github.com/lanej/charmed/theme"
)

// Permissions displays observed mode bits, not effective access. Mode may describe
// a link or its target; the caller owns that observation and its provenance.
// The zero value is unknown. Available distinguishes known 000 from missing data.
type Permissions struct {
	Mode      fs.FileMode
	Available bool
}

const (
	permissionGroupBits = 3
	permissionMask      = 0o7
	permissionRead      = 0o4
	permissionWrite     = 0o2
	permissionExecute   = 0o1
)

// Label retains exact rwx bits and the conventional s/S/t/T special-bit spelling.
// Unknown data is nine question marks, never nine unset bits.
func (p Permissions) Label() string {
	if !p.Available {
		return "?????????"
	}

	label := []byte("---------")

	for i, bit := range []fs.FileMode{0o400, 0o200, 0o100, 0o040, 0o020, 0o010, 0o004, 0o002, 0o001} {
		if p.Mode&bit != 0 {
			label[i] = "rwx"[i%permissionGroupBits]
		}
	}

	for i, special := range []struct {
		bit   fs.FileMode
		lower byte
		upper byte
	}{{fs.ModeSetuid, 's', 'S'}, {fs.ModeSetgid, 's', 'S'}, {fs.ModeSticky, 't', 'T'}} {
		if p.Mode&special.bit != 0 {
			position := (i+1)*permissionGroupBits - 1

			label[position] = special.upper
			if p.Mode&(0o100>>uint(i*permissionGroupBits)) != 0 {
				label[position] = special.lower
			}
		}
	}

	return string(label)
}

// Groups returns owner, group and other in that order. Unicode uses one solid
// segment glyph per triplet; ASCII uses the exact three-character spelling.
// The Unicode glyphs encode ordinary rwx bits; Compact adds special-bit labels.
func (p Permissions) Groups(t theme.Theme) [3]string {
	var groups [3]string

	label := p.Label()
	for i := range groups {
		bits := (p.Mode >> uint((len(groups)-1-i)*permissionGroupBits)) & permissionMask

		glyph, foreground := permissionGlyph(bits, t.Permissions)
		if !p.Available {
			glyph, foreground = "?", t.Permissions.Unknown
		}

		if t.ASCII {
			glyph = label[i*permissionGroupBits : (i+1)*permissionGroupBits]
		}

		style := t.Text()
		if foreground != nil && t.Mode != theme.Mono {
			style = style.Foreground(foreground)
		}

		groups[i] = style.Render(glyph)
	}

	return groups
}

// Compact keeps explicit special-bit labels alongside the three ordinary groups.
func (p Permissions) Compact(t theme.Theme) string {
	groups := p.Groups(t)
	parts := []string{strings.Join(groups[:], " ")}

	if p.Available {
		for _, special := range []struct {
			bit   fs.FileMode
			label string
		}{{fs.ModeSetuid, "setuid"}, {fs.ModeSetgid, "setgid"}, {fs.ModeSticky, "sticky"}} {
			if p.Mode&special.bit != 0 {
				parts = append(parts, t.Text().Underline(true).Render(special.label))
			}
		}
	}

	return strings.Join(parts, " ")
}

func permissionGlyph(bits fs.FileMode, palette theme.PermissionColors) (string, color.Color) {
	glyphs := [8]string{"∅", "▖", "▝", "▞", "▘", "▌", "▀", "▛"}
	foreground := palette.None

	switch uint32(bits & (permissionWrite | permissionExecute)) {
	case permissionWrite | permissionExecute:
		foreground = palette.WriteExecute
	case permissionWrite:
		foreground = palette.Write
	case permissionExecute:
		foreground = palette.Execute
	default:
		if bits&permissionRead != 0 {
			foreground = palette.Read
		}
	}

	return glyphs[bits&permissionMask], foreground
}

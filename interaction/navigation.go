package interaction

import "charm.land/bubbles/v2/key"

type KeyLayout string

const (
	DefaultKeys  KeyLayout = ""
	StandardKeys KeyLayout = "standard"
	VimKeys      KeyLayout = "vim"
)

type NavigationKeys struct {
	Up, Down, Left, Right                      key.Binding
	PageUp, PageDown, HalfPageUp, HalfPageDown key.Binding
	Top, Bottom                                key.Binding
}

// Navigation supplies shared presets. Apply a preset before individual native
// key-map overrides; resizing must not reapply it over those customizations.
func Navigation(preset KeyLayout) (NavigationKeys, bool) {
	vim := false

	switch preset {
	case DefaultKeys, StandardKeys:
	case VimKeys:
		vim = true
	default:
		return NavigationKeys{}, false
	}

	binding := func(standard, vi, label string) key.Binding {
		keys := []string{standard}

		help := standard
		if glyph, ok := map[string]string{"up": "↑", "down": "↓", "left": "←", "right": "→"}[standard]; ok {
			help = glyph
		}

		if vim && vi != standard {
			keys = append(keys, vi)
			help = vi + "/" + help
		}

		return key.NewBinding(key.WithKeys(keys...), key.WithHelp(help, label))
	}

	return NavigationKeys{Up: binding("up", "k", "up"), Down: binding("down", "j", "down"), Left: binding("left", "h", "left"), Right: binding("right", "l", "right"), PageUp: binding("pgup", "ctrl+b", "page up"), PageDown: binding("pgdown", "ctrl+f", "page down"), HalfPageUp: binding("ctrl+u", "ctrl+u", "half page up"), HalfPageDown: binding("ctrl+d", "ctrl+d", "half page down"), Top: binding("home", "home", "first"), Bottom: binding("end", "G", "last")}, true
}

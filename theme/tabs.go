package theme

import "image/color"

// TabColors identifies the active and inactive tab independently of keyboard focus.
type TabColors struct {
	ActiveForeground, ActiveBackground     color.Color
	InactiveForeground, InactiveBackground color.Color
}

func sageTabs() TabColors {
	return TabColors{
		ActiveForeground: hex(roleTabActiveFg), ActiveBackground: hex(roleTabActiveBg),
		InactiveForeground: hex(roleTabInactiveFg), InactiveBackground: hex(roleTabInactiveBg),
	}
}

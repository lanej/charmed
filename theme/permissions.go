package theme

import "image/color"

// PermissionColors group capabilities independently of status and health colors.
// Write and Execute include combinations with Read; WriteExecute includes both.
type PermissionColors struct {
	Read, Write, Execute, WriteExecute color.Color
	None, Unknown                      color.Color
}

func permissionColors(mode Mode) PermissionColors {
	switch mode {
	case Light:
		return PermissionColors{
			Read: hex("#365e70"), Write: hex("#725a20"),
			Execute: hex("#65517f"), WriteExecute: hex("#7b4252"),
			None: hex("#566273"), Unknown: hex("#475e86"),
		}
	case Contrast:
		return PermissionColors{
			Read: hex("#b7dfea"), Write: hex("#f4dfa6"),
			Execute: hex("#d9c8f0"), WriteExecute: hex("#f3c6d0"),
			None: hex("#b0bdcf"), Unknown: hex("#c1d9f5"),
		}
	case Mono:
		return PermissionColors{}
	case Dark:
		return PermissionColors{
			Read: hex("#99c8d4"), Write: hex("#e1cd91"),
			Execute: hex("#c3b2df"), WriteExecute: hex("#e6b6bf"),
			None: hex("#a1acbd"), Unknown: hex("#a9c4e5"),
		}
	default:
		return PermissionColors{}
	}
}

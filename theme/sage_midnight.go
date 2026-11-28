package theme

// Shared swatches from lanej/dotfiles, design-system/themes/sage.toml.
const (
	sage      = "#acc8b7"
	teal      = "#8fb9a8"
	blue      = "#95b5d3"
	cyan      = "#8cc5cc"
	lichen    = "#b7cb91"
	lavender  = "#c4b2d9"
	dustyRose = "#d89dae"
	copper    = "#dda783"
	amber     = "#e8c88b"
	muted     = "#a3b1a7"
	comment   = "#9cab9f"
)

// Sage Midnight supplies the built-in dark profile. Charmed extends the shared
// palette with lifecycle shades, an achromatic OFF and an ordered sage ramp.
func sageMidnight() Theme {
	http, syntax := httpColors(Dark)

	return Theme{
		Mode:                Dark,
		Background:          hex("#0d1513"),
		Surface:             hex("#17231e"),
		Foreground:          hex("#d7dfd9"),
		Muted:               hex(muted),
		Border:              hex("#6b8575"),
		Accent:              hex(sage),
		FocusBorder:         hex(lavender),
		SelectionForeground: hex("#eef1ee"),
		SelectionBackground: hex("#26352d"),
		Link:                hex(sage),
		DimLink:             hex(teal),
		Success:             hex("#a6cd82"),
		Info:                hex(cyan),
		Warning:             hex(amber),
		Danger:              hex("#f18c96"),
		Unknown:             hex("#a2b3c1"),
		Inactive:            hex("#b1b1b1"),
		DimUnknown:          hex("#8c9daa"),
		DimInactive:         hex("#9b9b9b"),
		Pending:             hex("#b7bde3"),
		Running:             hex(teal),
		Paused:              hex("#c0ac95"),
		DimPending:          hex("#959bb9"),
		DimRunning:          hex("#7da393"),
		DimPaused:           hex("#aa9783"),
		DimForeground:       hex(muted),
		DimSuccess:          hex("#87a86c"),
		DimInfo:             hex("#75a5ab"),
		DimWarning:          hex("#bca473"),
		DimDanger:           hex("#df828c"),
		Trace:               hex(comment),
		Debug:               hex("#93accc"),
		Fatal:               hex("#d4a6ce"),
		DimTrace:            hex("#8f9e93"),
		DimDebug:            hex("#869cb9"),
		DimFatal:            hex("#b48fb1"),
		Series:              colors("#a798c9", blue, dustyRose, copper, lichen),
		Sequential:          colors("#26352d", "#40574a", "#617b69", "#849f8d", sage),
		Permissions:         permissionColors(Dark),
		HTTP:                http,
		Syntax:              syntax,
		Diff:                diffColors(Dark),
	}
}

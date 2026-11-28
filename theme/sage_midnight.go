package theme

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Dotfiles supplies the named roles and state aliases. Charmed owns the dim
// emphasis, categorical series and ordered intensity scale for native widgets.
func sageMidnight() Theme {
	http, syntax := httpColors(Dark)

	return Theme{
		Mode:         Dark,
		Background:   hex(roleSurface),
		Surface:      hex(roleSurfaceRaised),
		Foreground:   hex(roleText),
		Muted:        hex(roleTextMuted),
		Border:       hex(roleBorder),
		Accent:       hex(roleAccent),
		HeadingColor: overrideColor(roleHeading, roleAccent),
		KeyHintColor: overrideColor(roleKeyhint, roleHeading),
		ConfigOn:     overrideColor(stateConfigurationOn, roleHeading),
		ConfigOff:    overrideColor(stateConfigurationOff, roleTextMuted),
		Canceled:     hex(stateExecutionCanceled), DimCanceled: sageDim(stateExecutionCanceled),
		UnknownOutcome: hex(roleTextMuted), DimUnknownOutcome: sageDim(roleTextMuted),
		FocusBorder:         hex(roleFocus),
		FocusBackground:     overrideColor(roleFocus, roleAccent),
		FocusForeground:     overrideColor(roleTextInverse, roleSurface),
		SelectionForeground: hex(roleSelectionFg),
		SelectionBackground: hex(roleSelectionBg),
		Link:                hex(roleLink),
		DimLink:             sageDim(roleLink),
		Match:               hex(roleMatch),
		Success:             hex(roleStatusSuccess),
		Info:                hex(roleStatusInfo),
		Warning:             hex(roleStatusWarning),
		Danger:              hex(roleStatusError),
		Unknown:             hex(stateHealthUnknown),
		DimUnknown:          sageDim(stateHealthUnknown),
		Inactive:            hex(stateConfigurationOff),
		DimInactive:         hex(stateConfigurationOff),
		Pending:             hex(stateExecutionQueued),
		DimPending:          sageDim(stateExecutionQueued),
		Running:             hex(stateExecutionRunning),
		DimRunning:          sageDim(stateExecutionRunning),
		Paused:              hex(stateExecutionPaused),
		DimPaused:           sageDim(stateExecutionPaused),
		Connected:           hex(stateConnectivityConnected),
		DimConnected:        sageDim(stateConnectivityConnected),
		Disconnected:        overrideColor(stateConnectivityDisconnected, stateConfigurationOff),
		DimDisconnected:     overrideColor(stateConnectivityDisconnected, stateConfigurationOff),
		Disabled:            overrideColor(stateInteractionDisabled, stateConfigurationOff),
		DimDisabled:         overrideColor(stateInteractionDisabled, stateConfigurationOff),
		Fresh:               hex(stateDataFresh),
		DimFresh:            sageDim(stateDataFresh),
		Missing:             hex(stateDataMissing),
		DimMissing:          sageDim(stateDataMissing),
		Stale:               overrideColor(stateDataStale, roleStatusWarning),
		DimForeground:       hex(roleTextMuted),
		DimSuccess:          sageDim(roleStatusSuccess),
		DimInfo:             sageDim(roleStatusInfo),
		DimWarning:          sageDim(roleStatusWarning),
		DimDanger:           sageDim(roleStatusError),
		Trace:               hex(stateSeverityTrace),
		DimTrace:            hex(stateSeverityTrace),
		Debug:               hex(stateSeverityDebug),
		DimDebug:            sageDim(stateSeverityDebug),
		Fatal:               hex(stateSeverityFatal),
		DimFatal:            sageDim(stateSeverityFatal),
		FatalBold:           stateSeverityFatalBold,
		Series:              colors("#a798c9", swatchBlue, swatchDustyRose, swatchCopper, swatchLichen),
		Sequential:          colors(roleSelectionBg, "#40574a", "#617b69", "#849f8d", roleAccent),
		Permissions:         permissionColors(Dark),
		HTTP:                http,
		Syntax:              syntax,
		Diff:                diffColors(Dark),
		Git:                 sageGitColors(),
		Tabs:                sageTabs(),
		Lexical: LexicalColors{
			Function: hex(roleSyntaxFunction), Type: hex(roleSyntaxType), Operator: hex(roleSyntaxOperator),
			BooleanTrue: hex(roleSyntaxBooleanTrue), BooleanFalse: hex(roleSyntaxBooleanFalse), Null: hex(roleSyntaxNull),
			Special: hex(roleSyntaxSpecial), Variable: hex(roleSyntaxVariable), Receiver: hex(roleSyntaxReceiver),
			Constant: hex(roleSyntaxConstant), VariableReference: hex(roleSyntaxVariableReference), Section: hex(roleSyntaxSection),
		},
	}
}

// Leave equivalent roles inherited so changing a caller's base color still updates its dependents.
func overrideColor(role, inherited string) color.Color {
	if role == inherited {
		return nil
	}

	return hex(role)
}

// This reduction preserves 4.5:1 on selected rows for the regular Sage inks.
// States using subtle ink keep the original color at dim emphasis.
func sageDim(role string) color.Color {
	return lipgloss.Darken(hex(role), 0.08)
}

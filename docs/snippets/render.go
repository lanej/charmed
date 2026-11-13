// Package snippets contains compiled consumer examples used by the documentation.
package snippets

import "fmt"

// Render runs only bounded rendering and construction, never a terminal program or initialization command.
func Render(id string) (string, error) {
	switch id {
	case "collection/scroll-pane":
		return ExampleCollectionScrollPane()
	case "collection/grid-inspector":
		return ExampleCollectionGridInspector()
	case "collection/keyed-table":
		return ExampleCollectionKeyedTable()
	case "input/radio-group":
		return ExampleInputRadioGroup()
	case "diagram/trace-replay":
		return ExampleDiagramTraceReplay()
	case "pattern/measurement-history":
		return ExamplePatternMeasurementHistory()
	case "git/patch-browser":
		return ExampleGitPatchBrowser()
	case "layout/flex":
		return ExampleLayoutFlex()
	case "layout/grid":
		return ExampleLayoutGrid()
	case "layout/flow":
		return ExampleLayoutFlow()
	case "layout/dock":
		return ExampleLayoutDock()
	case "layout/split-pane":
		return ExampleLayoutSplitPane()
	case "layout/responsive":
		return ExampleLayoutResponsive()
	case "tui/screen":
		return ExampleTuiScreen()
	case "tui/regions":
		return ExampleTuiRegions()
	case "tui/validation":
		return ExampleTuiValidation()
	case "tui/static":
		return ExampleTuiStatic()

	case "collection/data-grid":
		return ExampleCollectionDataGrid()
	case "collection/status-matrix":
		return ExampleCollectionStatusMatrix()
	case "feedback/status-cell":
		return ExampleFeedbackStatusCell(), nil
	case "cli/help":
		return ExampleCLIHelp(), nil
	case "git/branch-row":
		return ExampleGitBranchRow(), nil
	case "git/branches":
		return ExampleGitBranches()
	case "git/browser":
		return ExampleGitBrowser()
	case "git/changes":
		return ExampleGitChanges()
	case "git/commit-id":
		return ExampleGitCommitId(), nil
	case "git/commit-row":
		return ExampleGitCommitRow(), nil
	case "git/conflict-state":
		return ExampleGitConflictState(), nil
	case "git/conflicts":
		return ExampleGitConflicts()
	case "git/diff":
		return ExampleGitDiff(), nil
	case "git/diff-summary":
		return ExampleGitDiffSummary(), nil
	case "git/divergence":
		return ExampleGitDivergence(), nil
	case "git/file-row":
		return ExampleGitFileRow(), nil
	case "git/file-state":
		return ExampleGitFileState(), nil
	case "git/head":
		return ExampleGitHead(), nil
	case "git/history":
		return ExampleGitHistory()
	case "git/hunk":
		return ExampleGitHunk(), nil
	case "git/ref":
		return ExampleGitRef(), nil
	case "git/staged":
		return ExampleGitStaged()
	case "git/staged-state":
		return ExampleGitStagedState(), nil
	case "git/stash-row":
		return ExampleGitStashRow(), nil
	case "git/stashes":
		return ExampleGitStashes()
	case "git/status":
		return ExampleGitStatus()
	case "git/summary":
		return ExampleGitSummary(), nil
	case "git/unstaged":
		return ExampleGitUnstaged()
	case "git/unstaged-state":
		return ExampleGitUnstagedState(), nil
	case "git/untracked":
		return ExampleGitUntracked()
	case "git/untracked-state":
		return ExampleGitUntrackedState(), nil
	case "git/worktree-row":
		return ExampleGitWorktreeRow(), nil
	case "git/worktree-state":
		return ExampleGitWorktreeState(), nil
	case "git/worktrees":
		return ExampleGitWorktrees()

	case "input/checkbox":
		return ExampleInputCheckbox(), nil
	case "input/radio":
		return ExampleInputRadio(), nil
	case "input/toggle":
		return ExampleInputToggle(), nil
	case "input/button":
		return ExampleInputButton(), nil
	case "feedback/freshness":
		return ExampleFeedbackFreshness(), nil
	case "feedback/evidence":
		return ExampleFeedbackEvidence(), nil
	case "feedback/evidence-summary":
		return ExampleFeedbackEvidenceSummary(), nil
	case "feedback/evidence-footer":
		return ExampleFeedbackEvidenceFooter(), nil
	case "template/report-page":
		return ExampleTemplateReportPage(), nil
	case "collection/static-table":
		return ExampleCollectionStaticTable(), nil
	case "layout/key-value":
		return ExampleLayoutKeyValue(), nil
	case "feedback/outcome":
		return ExampleFeedbackOutcome(), nil
	case "text/value-state":
		return ExampleTextValueState(), nil
	case "text/quantity":
		return ExampleTextQuantity(), nil
	case "text/identifier":
		return ExampleTextIdentifier(), nil
	case "text/match":
		return ExampleTextMatch(), nil
	case "text/sort":
		return ExampleTextSort(), nil
	case "text/change":
		return ExampleTextChange(), nil
	case "text/omission":
		return ExampleTextOmission(), nil
	case "foundations/theme":
		return ExampleFoundationsTheme(), nil
	case "foundations/layout":
		return ExampleFoundationsLayout(), nil
	case "foundations/glyphs":
		return ExampleFoundationsGlyphs(), nil
	case "foundations/permissions":
		return ExampleFoundationsPermissions(), nil
	case "foundations/tokens":
		return ExampleFoundationsTokens(), nil
	case "foundations/focus":
		return ExampleFoundationsFocus(), nil
	case "feedback/motion":
		return ExampleFeedbackMotion(), nil
	case "text/primitives":
		return ExampleTextPrimitives(), nil
	case "input/text":
		return ExampleInputText(), nil
	case "input/textarea":
		return ExampleInputTextarea(), nil
	case "feedback/spinner":
		return ExampleFeedbackSpinner(), nil
	case "feedback/progress":
		return ExampleFeedbackProgress(), nil
	case "feedback/time":
		return ExampleFeedbackTime(), nil
	case "collection/json-viewer":
		return ExampleCollectionJSONViewer()
	case "collection/viewport":
		return ExampleCollectionViewport(), nil
	case "chart/sparkline":
		return ExampleChartSparkline(), nil
	case "chart/timeseries":
		return ExampleChartTimeseries(), nil
	case "chart/line":
		return ExampleChartLine(), nil
	case "chart/scatter":
		return ExampleChartScatter(), nil
	case "chart/bars":
		return ExampleChartBars(), nil
	case "chart/heatmap":
		return ExampleChartHeatmap(), nil
	case "chart/stream":
		return ExampleChartStream(), nil
	case "chart/wave":
		return ExampleChartWave(), nil
	case "chart/ohlc":
		return ExampleChartOhlc()
	case "chart/canvas":
		return ExampleChartCanvas(), nil
	case "chart/waffle":
		return ExampleChartWaffle(), nil
	case "text/signed-number":
		return ExampleTextSignedNumber(), nil
	case "text/pill":
		return ExampleTextPill(), nil
	case "feedback/log-level":
		return ExampleFeedbackLogLevel(), nil
	case "feedback/status":
		return ExampleFeedbackStatus(), nil
	case "chart/picture":
		return ExampleChartPicture(), nil
	case "chart/interval-band":
		return ExampleChartIntervalBand(), nil
	case "diagram/state-transitions":
		return ExampleDiagramStateTransitions()
	case "input/field":
		return ExampleInputField(), nil
	case "feedback/paginator":
		return ExampleFeedbackPaginator(), nil
	case "feedback/help":
		return ExampleFeedbackHelp(), nil
	case "chart/waffle-header":
		return ExampleChartWaffleHeader(), nil
	case "chart/band-matrix":
		return ExampleChartBandMatrix(), nil
	case "chart/metric":
		return ExampleChartMetric(), nil
	case "chart/panel":
		return ExampleChartPanel(), nil
	case "feedback/attention":
		return ExampleFeedbackAttention(), nil
	case "chart/distributions":
		return ExampleChartDistributions(), nil
	case "chart/delta":
		return ExampleChartDelta(), nil
	case "feedback/percent":
		return ExampleFeedbackPercent(), nil
	case "feedback/connection":
		return ExampleFeedbackConnection(), nil
	case "navigation/tabs":
		return ExampleNavigationTabs(), nil
	case "layout/link-group":
		return ExampleLayoutLinkGroup()
	case "layout/header":
		return ExampleLayoutHeader(), nil
	case "layout/identity-header":
		return ExampleLayoutIdentityHeader(), nil
	case "layout/section":
		return ExampleLayoutSection(), nil
	case "layout/footer":
		return ExampleLayoutFooter(), nil
	case "template/application":
		return ExampleTemplateApplication(), nil
	case "collection/table":
		return ExampleCollectionTable(), nil
	case "collection/list":
		return ExampleCollectionList(), nil
	case "collection/tree":
		return ExampleCollectionTree(), nil
	case "collection/files":
		return ExampleCollectionFiles()
	case "text/code":
		return ExampleTextCode()
	case "text/markdown":
		return ExampleTextMarkdown()
	case "form/settings":
		return ExampleFormSettings(), nil
	case "collection/logs":
		return ExampleCollectionLogs(), nil
	case "chart/comparison-panel":
		return ExampleChartComparisonPanel(), nil
	case "screen/services":
		return ExampleScreenServices(), nil
	case "screen/cohorts":
		return ExampleScreenCohorts()
	case "screen/deployment":
		return ExampleScreenDeployment()
	case "text/presentation":
		return ExampleTextPresentation(), nil
	case "text/link":
		return ExampleTextLink()
	case "text/url":
		return ExampleTextUrl(), nil
	case "feedback/follow":
		return ExampleFeedbackFollow(), nil
	case "feedback/scrollbar":
		return ExampleFeedbackScrollbar(), nil
	case "text/disclosure":
		return ExampleTextDisclosure(), nil
	case "text/key-hint":
		return ExampleTextKeyHint(), nil
	case "text/pending":
		return ExampleTextPending(), nil
	case "text/observation":
		return ExampleTextObservation(), nil
	case "text/time":
		return ExampleTextTime(), nil
	case "text/http-method":
		return ExampleTextHttpMethod(), nil
	case "text/http-status":
		return ExampleTextHttpStatus(), nil
	case "text/syntax":
		return ExampleTextSyntax(), nil
	case "collection/log-tree":
		return ExampleCollectionLogTree()
	case "collection/http-log":
		return ExampleCollectionHttpLog()
	case "pattern/log-inspector":
		return ExamplePatternLogInspector()
	case "collection/command-palette":
		return ExampleCollectionCommandPalette()
	default:
		return "", fmt.Errorf("missing usage example for %s", id)
	}
}

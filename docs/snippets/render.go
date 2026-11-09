// Package snippets contains compiled consumer examples used by the documentation.
package snippets

import "fmt"

// Render runs only bounded rendering and construction, never a terminal program or initialization command.
func Render(id string) (string, error) {
	switch id {
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
		return ExampleCollectionFiles(), nil
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
	case "screen/deployment":
		return ExampleScreenDeployment()
	case "text/presentation":
		return ExampleTextPresentation(), nil
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

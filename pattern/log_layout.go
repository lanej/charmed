package pattern

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

type InspectorLayout string

const (
	InspectorAuto     InspectorLayout = ""
	InspectorSidebar  InspectorLayout = "sidebar"
	InspectorTimeline InspectorLayout = "timeline"
	InspectorColumns  InspectorLayout = "columns"
	InspectorFocused  InspectorLayout = "focused"
)

const (
	inspectorMinLogs           = 40
	inspectorTimelineMinHeight = 28
	inspectorRequestsShare     = .34
	inspectorDetailsShare      = .60
	inspectorTimelineShare     = .40
)

type inspectorRegions struct{ requests, details, logs layout.Size }

// SetLayout changes presentation without rebuilding captured data or correlation.
// Resizing preserves native styles/key maps; scroll offsets clamp to new sizes.
func (m *LogInspector) SetLayout(preset InspectorLayout) bool {
	switch preset {
	case InspectorAuto, InspectorSidebar, InspectorTimeline, InspectorColumns, InspectorFocused:
	default:
		return false
	}

	m.preferred = preset
	m.Resize(m.size)

	return true
}

func (m *LogInspector) Layout() InspectorLayout       { return m.preferred }
func (m *LogInspector) ActiveLayout() InspectorLayout { return m.active }

func (m *LogInspector) cycleLayout() {
	switch m.preferred {
	case InspectorAuto, InspectorSidebar:
		m.SetLayout(InspectorTimeline)
	case InspectorTimeline:
		m.SetLayout(InspectorColumns)
	case InspectorColumns:
		m.SetLayout(InspectorFocused)
	case InspectorFocused:
		m.SetLayout(InspectorSidebar)
	default:
		return
	}
}

func planInspector(size layout.Size, preset InspectorLayout) (InspectorLayout, inspectorRegions) {
	height := max(1, size.Height-2)
	full := layout.Size{Width: size.Width, Height: height}

	focused := inspectorRegions{requests: full, details: full, logs: full}
	if preset == InspectorFocused {
		return InspectorFocused, focused
	}

	if preset == InspectorTimeline && size.Width >= inspectorMinDetails+inspectorMinLogs+theme.SpaceSmall && size.Height >= inspectorTimelineMinHeight {
		top := max(10, int(float64(height)*inspectorTimelineShare))
		bottom := height - top
		detailsWidth := max(inspectorMinDetails, (size.Width-theme.SpaceSmall)/2)
		logsWidth := size.Width - theme.SpaceSmall - detailsWidth

		return InspectorTimeline, inspectorRegions{requests: layout.Size{Width: size.Width, Height: top}, details: layout.Size{Width: detailsWidth, Height: bottom}, logs: layout.Size{Width: logsWidth, Height: bottom}}
	}

	if preset == InspectorColumns && size.Width >= inspectorMinRequests+inspectorMinDetails+inspectorMinLogs+2*theme.SpaceSmall && size.Height >= inspectorMinHeight {
		extra := size.Width - inspectorMinRequests - inspectorMinDetails - inspectorMinLogs - 2*theme.SpaceSmall
		requestsWidth := inspectorMinRequests + extra/4
		detailsWidth := inspectorMinDetails + extra/2
		logsWidth := size.Width - requestsWidth - detailsWidth - 2*theme.SpaceSmall

		return InspectorColumns, inspectorRegions{requests: layout.Size{Width: requestsWidth, Height: height}, details: layout.Size{Width: detailsWidth, Height: height}, logs: layout.Size{Width: logsWidth, Height: height}}
	}

	left, right, wide := layout.Split(size.Width, theme.SpaceSmall, inspectorMinRequests, inspectorMinDetails, inspectorRequestsShare)
	if wide && size.Height >= inspectorMinHeight {
		detailHeight := max(7, int(float64(height)*inspectorDetailsShare))
		return InspectorSidebar, inspectorRegions{requests: layout.Size{Width: left, Height: height}, details: layout.Size{Width: right, Height: detailHeight}, logs: layout.Size{Width: right, Height: height - detailHeight}}
	}

	return InspectorFocused, focused
}

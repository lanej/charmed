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
	inspectorMinLogs               = 40
	inspectorTimelineMinHeight     = 18
	inspectorTimelinePreferredLogs = 40
	inspectorRequestsShare         = .34
	inspectorDetailsShare          = .60
	inspectorTimelineShare         = .30
	inspectorTimelineDetailsShare  = .75
)

type inspectorRegions struct {
	requests, details, logs      layout.Size
	requestsAbove, bottomStacked bool
}

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
	preset := m.preferred
	if preset == InspectorAuto {
		preset = m.active
	}

	switch preset {
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

func planInspector(size layout.Size, preset InspectorLayout, requestHeight int) (InspectorLayout, inspectorRegions) {
	height := max(1, size.Height)
	full := layout.Size{Width: size.Width, Height: height}

	focused := inspectorRegions{requests: full, details: full, logs: full}
	if preset == InspectorFocused {
		return InspectorFocused, focused
	}

	if (preset == InspectorTimeline || preset == InspectorAuto) && size.Width >= inspectorMinDetails+inspectorMinLogs+theme.SpaceSmall && size.Height >= inspectorTimelineMinHeight {
		top := min(max(3, requestHeight+2), max(4, int(float64(height)*inspectorTimelineShare)))
		bottom := height - top
		logsMinimum := min(inspectorTimelinePreferredLogs, size.Width-theme.SpaceSmall-inspectorMinDetails)
		detailsWidth := min(size.Width-theme.SpaceSmall-logsMinimum, max(inspectorMinDetails, int(float64(size.Width-theme.SpaceSmall)*inspectorTimelineDetailsShare)))

		logsWidth := size.Width - theme.SpaceSmall - detailsWidth
		if detailsWidth < inspectorDetailColumnsWidth {
			logsHeight := min(7, max(3, bottom/3))
			return InspectorTimeline, inspectorRegions{requests: layout.Size{Width: size.Width, Height: top}, details: layout.Size{Width: size.Width, Height: bottom - logsHeight}, logs: layout.Size{Width: size.Width, Height: logsHeight}, bottomStacked: true}
		}

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

// foldInspector returns freed rows and columns to expanded siblings. Folded
// headings remain keyboard targets, so focus and restoration stay predictable.
func foldInspector(size layout.Size, active InspectorLayout, regions inspectorRegions, collapsed map[string]bool) inspectorRegions {
	height := max(1, size.Height)

	requestClosed, detailsClosed, logsClosed := collapsed[requestPane], collapsed[httpPane], collapsed[logPane]
	if !requestClosed && !detailsClosed && !logsClosed {
		return regions
	}

	switch active {
	case InspectorTimeline:
		logsHeight := regions.logs.Height
		if requestClosed {
			regions.requests.Height = 1
		} else if detailsClosed && logsClosed {
			regions.requests.Height = max(1, height-3)
		}

		bottom := max(1, height-regions.requests.Height)

		regions.details.Height, regions.logs.Height = bottom, bottom
		if detailsClosed || logsClosed || regions.bottomStacked {
			regions.bottomStacked = true
			regions.details.Width, regions.logs.Width = size.Width, size.Width

			regions.details.Height, regions.logs.Height = splitFoldedHeight(bottom, detailsClosed, logsClosed)
			if !detailsClosed && !logsClosed {
				regions.logs.Height = min(logsHeight, max(1, bottom-3))
				regions.details.Height = bottom - regions.logs.Height
			}
		}
	case InspectorAuto, InspectorSidebar:
		if requestClosed || detailsClosed && logsClosed {
			regions.requestsAbove = true

			regions.requests = layout.Size{Width: size.Width, Height: 1}
			if !requestClosed {
				regions.requests.Height = max(1, height-3)
			}

			regions.details.Width, regions.logs.Width = size.Width, size.Width
			regions.details.Height, regions.logs.Height = splitFoldedHeight(max(1, height-regions.requests.Height), detailsClosed, logsClosed)
		} else {
			regions.details.Height, regions.logs.Height = splitFoldedHeight(height, detailsClosed, logsClosed)
		}
	case InspectorColumns:
		panes := []struct {
			id      string
			minimum int
			size    *layout.Size
		}{{requestPane, inspectorMinRequests, &regions.requests}, {httpPane, inspectorMinDetails, &regions.details}, {logPane, inspectorMinLogs, &regions.logs}}
		count, minimum := 0, 0

		for _, pane := range panes {
			if !collapsed[pane.id] {
				count++
				minimum += pane.minimum
			}
		}

		remaining := size.Width - max(0, count-1)*theme.SpaceSmall - minimum

		folded := 3 - count
		if collapsed[httpPane] {
			folded++
		}

		for _, pane := range panes {
			if collapsed[pane.id] {
				*pane.size = layout.Size{Width: size.Width, Height: 1}
				if pane.id == httpPane {
					pane.size.Height = 2
				}

				continue
			}

			extra := remaining / count
			*pane.size = layout.Size{Width: pane.minimum + extra, Height: max(1, height-folded)}
			remaining -= extra
			count--
		}
	case InspectorFocused:
		if requestClosed {
			regions.requests.Height = 1
		}

		if detailsClosed {
			regions.details.Height = 2
		}

		if logsClosed {
			regions.logs.Height = 1
		}
	default:
	}

	return regions
}

func splitFoldedHeight(height int, firstClosed, secondClosed bool) (int, int) {
	if firstClosed && secondClosed {
		return 2, 1
	}

	if firstClosed {
		return 2, max(1, height-2)
	}

	if secondClosed {
		return max(1, height-1), 1
	}

	first := max(3, int(float64(height)*inspectorDetailsShare))

	return first, max(1, height-first)
}

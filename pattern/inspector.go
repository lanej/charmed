package pattern

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
)

func (m *CohortComparison) inspectorContext() string {
	if m.selected.MemberID == "" {
		return "No member observations"
	}
	var requests uint64
	for _, p := range m.visible {
		if p.MemberKey == m.selected {
			requests = p.Requests
			break
		}
	}

	return ansi.Hardwrap("Version "+m.selected.CohortID, max(1, layout.PanelContentSize(m.sideSize).Width), true) + fmt.Sprintf("\n%d requests", requests)
}

func (m *CohortComparison) updateInspector() {
	showAttention := false
	for _, policy := range m.boundaries {
		showAttention = showAttention || policy.Hosts.Configured()
	}

	width := max(1, layout.PanelContentSize(m.sideSize).Width)

	const codeWidth, minRateWidth, defaultRateWidth, distanceWidth = 5, 5, 16, 6

	const rateColumn = 1

	columns := []table.Column{{Title: "CODE", Width: codeWidth}, {Title: "RATE", Width: defaultRateWidth}, {Title: "z(ref)", Width: distanceWidth}}
	if showAttention {
		columns = append(columns, table.Column{Title: "ATTN", Width: feedback.MaxAttentionWidth(m.theme)})
		occupied := collection.DenseTablePaddingWidth(len(columns))

		for i, column := range columns {
			if i != rateColumn {
				occupied += column.Width
			}
		}

		columns[rateColumn].Width = max(minRateWidth, width-occupied)
	}

	m.Inspector.SetRows(nil)
	m.Inspector.SetColumns(columns)
	m.Inspector.SetWidth(width)
	m.Inspector.SetHeight(max(1, m.inspectorHeight-3-lipgloss.Height(m.inspectorContext())))

	var rows []table.Row

	if m.selected.MemberID != "" {
		for i := range m.data.Metrics {
			for _, summary := range m.summaries[i] {
				for _, p := range summary.Points {
					if p.MemberKey != m.selected {
						continue
					}
					z := "n/a"
					if value, ok := cohort.Distance(p, m.summaries[i][0]); ok {
						z = m.signedNumber(i, value, 2, "", true).View(m.theme)
					}
					palette := m.theme
					switch m.metricTones[m.data.Metrics[i].ID] {
					case text.Success:
						palette.Accent = palette.Success
					case text.Info:
						palette.Accent = palette.Info
					case text.Warning:
						palette.Accent = palette.Warning
					case text.Danger:
						palette.Accent = palette.Danger
					}

					row := table.Row{m.metricLabel(i, true), feedback.PercentBar(p.Rate*100, p.HasRate, columns[rateColumn].Width, palette), z}
					if showAttention {
						row = append(row, feedback.Attention(m.assessments[i].Members[p.MemberKey].Level, m.theme))
					}

					rows = append(rows, row)
				}
			}
		}
	}
	m.Inspector.SetRows(rows)
	if len(rows) > 0 {
		m.Inspector.SetCursor(m.metric)
	}
}

func (m *CohortComparison) inspectorView() string {
	title := "Member inspector"
	if m.selected.MemberID != "" {
		title = "Member " + m.selected.MemberID
	}
	return layout.Panel(title, strings.Join([]string{m.inspectorContext(), m.Inspector.View()}, "\n"), layout.Size{Width: m.sideSize.Width, Height: m.inspectorHeight}, m.theme, m.focus.Current() == "inspector")
}

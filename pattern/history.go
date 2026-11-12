package pattern

import (
	"fmt"
	"strings"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"

	"github.com/charmbracelet/x/ansi"
)

type historyGap struct{ start, end time.Time }

func appendHistoryGap(gaps []historyGap, start, end time.Time) []historyGap {
	if len(gaps) > 0 && gaps[len(gaps)-1].end.Equal(start) {
		gaps[len(gaps)-1].end = end
		return gaps
	}

	return append(gaps, historyGap{start: start, end: end})
}

func historyGapLabel(gaps []historyGap) string {
	labels := make([]string, len(gaps))
	for i, gap := range gaps {
		labels[i] = gap.start.UTC().Format(time.RFC3339Nano) + "–" + gap.end.UTC().Format(time.RFC3339Nano)
	}

	return strings.Join(labels, ", ")
}

func (m *CohortComparison) historyRows(index, width int) []string {
	if len(m.historyNotes[index]) == 0 {
		return nil
	}

	rows := []string{ansi.Wrap("History lines bridge missing intervals:", max(1, width), "")}
	for _, note := range m.historyNotes[index] {
		rows = append(rows, ansi.Wrap(note, max(1, width), ""))
	}

	return rows
}

func (m *CohortComparison) historySeries(metric cohort.Metric) ([]chart.TimeSeriesData, []string) {
	series := make([]chart.TimeSeriesData, len(m.data.Cohorts)+1)

	selected := len(m.data.Cohorts)
	for c, group := range m.data.Cohorts {
		series[c] = chart.TimeSeriesData{ID: fmt.Sprintf("cohort:%d", c), Label: group.Label, Unit: metric.DisplayUnit()}
	}

	series[selected] = chart.TimeSeriesData{ID: "selected", Label: "member", Unit: metric.DisplayUnit()}
	gaps := make([][]historyGap, len(series))

	var previousEnd time.Time
	for _, window := range m.data.Windows {
		if !previousEnd.IsZero() && window.Start.After(previousEnd) {
			for c := range gaps {
				gaps[c] = appendHistoryGap(gaps[c], previousEnd, window.Start)
			}
		}

		previousEnd = window.End
		selectedAvailable := false

		for c, group := range m.data.Cohorts {
			summary := cohort.AnalyzeMetric(window.Samples, group.ID, metric, m.minRequests)
			if summary.HasValue {
				series[c].Points = append(series[c].Points, timeserieslinechart.TimePoint{Time: window.End, Value: summary.Value * metric.Scale()})
			} else {
				gaps[c] = appendHistoryGap(gaps[c], window.Start, window.End)
			}

			for _, point := range summary.Points {
				if point.MemberKey == m.selected && point.HasValue {
					series[selected].Points = append(series[selected].Points, timeserieslinechart.TimePoint{Time: window.End, Value: point.Value * metric.Scale()})
					selectedAvailable = true
				}
			}
		}

		if !selectedAvailable && m.selected.MemberID != "" {
			gaps[selected] = appendHistoryGap(gaps[selected], window.Start, window.End)
		}
	}

	var notes []string

	for c, missing := range gaps {
		if len(missing) == 0 || (c == selected && m.selected.MemberID == "") {
			continue
		}

		label := series[c].Label
		if c == selected {
			label = "Member " + m.selected.CohortID + "/" + m.selected.MemberID
		}

		notes = append(notes, label+" missing "+historyGapLabel(missing))
	}

	return series, notes
}

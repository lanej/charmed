package chart

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type (
	DeltaRow struct {
		Label    string
		Property DeltaProperty
	}
	BandRow struct {
		Title        string
		Comparisons  []DeltaRow
		Distribution DistributionComparison
	}
)

// BandMatrix shares one physical plot rectangle across deltas and distributions.
// Each metric retains its own numeric domain. Width hints reserve minimum label
// and value columns; when those cannot fit, labels wrap above full-width tracks.
type BandMatrix struct {
	Title                  string
	Rows                   []BandRow
	LabelWidth, ValueWidth int
}

type matrixColumns struct {
	label, plot, value int
	inline             bool
}

func (b BandMatrix) columns(width int) matrixColumns {
	label, value := b.LabelWidth, b.ValueWidth
	for _, row := range b.Rows {
		for _, comparison := range row.Comparisons {
			label = max(label, lipgloss.Width(comparison.Label))
			value = max(value, lipgloss.Width(comparison.Property.Label()))
		}

		for _, series := range row.Distribution.Series {
			label = max(label, lipgloss.Width(series.Label))
			value = max(value, lipgloss.Width(text.AbbreviatedNumber(series.Mean, 3, row.Distribution.Unit, series.Members > 0)))
		}
	}

	label = min(label, max(0, width-value-5))
	plot := width - label - value - 2

	inline := label > 0 && plot >= 3
	if !inline {
		label, value, plot = 0, 0, width
	}

	if plot%2 == 0 && plot > 1 {
		plot--
	}

	return matrixColumns{label: label, plot: max(1, plot), value: value, inline: inline}
}

func matrixTrack(label, value, track string, columns matrixColumns, width int) string {
	if !columns.inline {
		return ansi.Wrap(label+" "+value, width, "") + "\n" + layout.Fit(track, layout.Size{Width: width, Height: 1})
	}

	left := layout.Fit(ansi.Truncate(label, columns.label, "~"), layout.Size{Width: columns.label, Height: 1})
	right := strings.Repeat(" ", max(0, columns.value-lipgloss.Width(value))) + value

	return layout.Fit(left+" "+track+" "+right, layout.Size{Width: width, Height: 1})
}

func (b BandMatrix) renderRows(width int, t theme.Theme) []string {
	columns := b.columns(width)

	var rows []string

	for _, row := range b.Rows {
		lines := []string{row.Title}
		for _, comparison := range row.Comparisons {
			lines = append(lines, matrixTrack(comparison.Label, comparison.Property.number().View(t), comparison.Property.Plot(columns.plot, t), columns, width))
		}

		distribution := row.Distribution

		low, high := distribution.domain()
		if !finite(distribution.Min) || !finite(distribution.Max) || distribution.Min >= distribution.Max || !finite(distribution.Max-distribution.Min) {
			lines = append(lines, ansi.Wrap("Distribution range unavailable", width, ""))
		} else {
			var legend []legendEntry

			for index, series := range distribution.Series {
				style := t.Text().Foreground(t.SeriesColor(index))
				value := text.AbbreviatedNumber(series.Mean, 3, distribution.Unit, series.Members > 0)

				track := strings.Repeat(" ", columns.plot)

				if series.Members > 0 {
					options := distributionOptions(distribution, series, index, low, high)

					track = IntervalBand(options, columns.plot, t)
					if distribution.ShowLegend {
						legend = append(legend, options.legendEntries(series.Label+" ", t)...)
					}
				}

				lines = append(lines, matrixTrack(style.Render(series.Label), style.Render(value), track, columns, width))
			}

			if finite(low) && finite(high) {
				left := text.AbbreviatedNumber(low, 3, distribution.Unit, true)
				right := text.AbbreviatedNumber(high, 3, distribution.Unit, true)

				axis := left + strings.Repeat(" ", max(1, columns.plot-lipgloss.Width(left)-lipgloss.Width(right))) + right
				for _, line := range strings.Split(ansi.Wrap(axis, columns.plot, ""), "\n") {
					lines = append(lines, strings.Repeat(" ", columns.offset())+line)
				}
			} else {
				lines = append(lines, ansi.Wrap("No eligible member spread", width, ""))
			}

			if distribution.ShowLegend {
				if key := renderLegend(legend, width, t); key != "" {
					lines = append(lines, key)
				}
			}
		}

		rows = append(rows, strings.Join(lines, "\n"))
	}

	if len(rows) == 0 {
		rows = []string{"No observations"}
	}

	return rows
}

func (c matrixColumns) offset() int {
	if c.inline {
		return c.label + 1
	}

	return 0
}

// Height measures exactly the columns and wrapping used by View, including its frame.
func (b BandMatrix) Height(width int, t theme.Theme) int {
	contentWidth := max(1, layout.PanelContentSize(layout.Size{Width: width}).Width)
	return lipgloss.Height(strings.Join(b.renderRows(contentWidth, t), "\n\n")) + 2*theme.BorderCells
}

func (b BandMatrix) View(size layout.Size, t theme.Theme, focused bool) string {
	rows := b.renderRows(max(1, layout.PanelContentSize(size).Width), t)
	return layout.Panel(b.Title, strings.Join(rows, "\n\n"), size, t, focused)
}

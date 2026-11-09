package chart

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

type legendEntry struct {
	glyph, label string
	style        lipgloss.Style
}

func renderLegend(entries []legendEntry, width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}
	var lines []string
	line := ""
	seen := map[string]bool{}
	for _, entry := range entries {
		token := entry.style.Render(entry.glyph) + " " + t.Text().Render(entry.label)
		if seen[token] {
			continue
		}
		seen[token] = true
		if strings.Contains(token, "\n") || lipgloss.Width(token) > width {
			if line != "" {
				lines = append(lines, line)
				line = ""
			}

			lines = append(lines, wrapLegendEntry(entry, width, t)...)

			continue
		}

		if line != "" && lipgloss.Width(line)+2+lipgloss.Width(token) > width {
			lines = append(lines, line)
			line = ""
		}
		if line != "" {
			line += "  "
		}
		line += layout.Fit(token, layout.Size{Width: min(width, lipgloss.Width(token)), Height: 1})
	}
	if line != "" {
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func wrapLegendEntry(entry legendEntry, width int, t theme.Theme) []string {
	prefixWidth := lipgloss.Width(entry.glyph) + 1
	if prefixWidth >= width {
		lines := []string{layout.Fit(entry.style.Render(entry.glyph), layout.Size{Width: width, Height: 1})}
		for _, label := range strings.Split(ansi.Wrap(entry.label, width, ""), "\n") {
			lines = append(lines, t.Text().Render(label))
		}

		return lines
	}

	labels := strings.Split(ansi.Wrap(entry.label, width-prefixWidth, ""), "\n")

	lines := make([]string, len(labels))
	for i, label := range labels {
		prefix := strings.Repeat(" ", prefixWidth)
		if i == 0 {
			prefix = entry.style.Render(entry.glyph) + " "
		}

		lines[i] = prefix + t.Text().Render(label)
	}

	return lines
}

func bandStyle(o BandOptions, t theme.Theme) lipgloss.Style {
	if o.Style != nil {
		return *o.Style
	}
	return lipgloss.NewStyle().Foreground(t.SeriesColor(o.Series))
}
func selectedStyle(t theme.Theme) lipgloss.Style { return t.Text().Bold(true) }

func (o BandOptions) legendEntries(prefix string, t theme.Theme) []legendEntry {
	var entries []legendEntry
	if !validBand(o) {
		return entries
	}
	if finite(o.Center) && o.Center >= o.Min && o.Center <= o.Max {
		entries = append(entries, legendEntry{"|", prefix + "mean", bandStyle(o, t).Bold(true)})
	}
	for _, marker := range o.Markers {
		if marker.Label != "" && finite(marker.Value) && marker.Value >= o.Min && marker.Value <= o.Max {
			entries = append(entries, legendEntry{string(marker.Glyph), marker.Label, marker.Style})
		}
	}
	if o.Selected != nil && finite(*o.Selected) && *o.Selected >= o.Min && *o.Selected <= o.Max {
		entries = append(entries, legendEntry{"X", "selected observation", selectedStyle(t)})
	}
	return entries
}

// Legend is an optional companion to the one-row interval atom. Marker labels
// and styles come from the same options used to draw the band.
func (o BandOptions) Legend(width int, t theme.Theme) string {
	return renderLegend(o.legendEntries("", t), width, t)
}

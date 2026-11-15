package chart_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/theme"
)

func TestDeltaPropertyKeepsAStableSymmetricScale(t *testing.T) {
	p := chart.DeltaProperty{Value: 0.2, Extent: 0.5, Available: true}
	positive := ansi.Strip(p.View(27, theme.New(theme.Mono)))
	p.Value = -0.2
	negative := ansi.Strip(p.View(27, theme.New(theme.Mono)))
	if len(positive) != 27 || len(negative) != 27 || !strings.Contains(positive, "+0.2pp") || !strings.Contains(negative, "-0.2pp") {
		t.Fatal("exact signed delta or allocation was lost")
	}
	left, right := strings.Index(positive, "|"), strings.Index(negative, "|")
	if left != right || strings.Count(positive, "#") != strings.Count(negative, "#") {
		t.Fatal("matching deltas do not have symmetric geometry")
	}
	p.Available = false
	if !strings.Contains(p.View(27, theme.New(theme.Mono)), "n/a") {
		t.Fatal("missing pair rendered as zero")
	}
	p.Available = true
	p.Value = 0.0001
	if !strings.Contains(p.View(27, theme.New(theme.Mono)), "0.0001pp") {
		t.Fatal("small nonzero delta was labeled as zero")
	}
	if chart.SymmetricExtent(0.3) != 0.5 || chart.SymmetricExtent(0.51) != 1 {
		t.Fatal("unreadable shared scale")
	}
}

func TestDeltaPropertiesShareCellColumnsAcrossUnitsAndMissingValues(t *testing.T) {
	palette := theme.New(theme.Mono)
	properties := []chart.DeltaProperty{
		{Value: 8, Extent: 20, Available: true, Unit: "C", LabelWidth: 9},
		{Value: 40, Extent: 50, Available: true, Unit: "kWh", LabelWidth: 9},
		{Value: -1.5, Extent: 5, Available: true, Unit: "µs", LabelWidth: 9},
	}
	center, rightEdge := -1, -1

	for _, property := range properties {
		frame := ansi.Strip(property.View(60, palette))

		column := ansi.StringWidth(frame[:strings.Index(frame, "|")])
		if center < 0 {
			center = column
		} else if column != center {
			t.Fatal("unit/value length moved the zero column")
		}

		label := property.Label()
		edge := ansi.StringWidth(frame[:strings.Index(frame, label)]) + ansi.StringWidth(label)

		if rightEdge < 0 {
			rightEdge = edge
		} else if edge != rightEdge {
			t.Fatal("Unicode unit moved the numeric column edge")
		}
	}

	missing := properties[0]
	missing.Available = false

	frame := ansi.Strip(missing.View(60, palette))
	if edge := ansi.StringWidth(frame[:strings.Index(frame, "n/a")]) + 3; edge != rightEdge {
		t.Fatal("unavailable value left the shared numeric column")
	}

	if !strings.HasSuffix(strings.TrimRight(frame, " "), "n/a") {
		t.Fatal("missing label did not occupy the shared numeric column")
	}
}

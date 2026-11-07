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
	if len(positive) != 27 || len(negative) != 27 || !strings.Contains(positive, "+0.200pp") || !strings.Contains(negative, "-0.200pp") {
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

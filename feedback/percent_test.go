package feedback_test

import (
	"math"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func TestPercentCellsPreserveValueAndAbsoluteScale(t *testing.T) {
	for _, value := range []float64{0, 1.968, 50, 100} {
		view := ansi.Strip(feedback.PercentBar(value, true, 16, theme.New(theme.Mono)))
		if len(view) != 16 {
			t.Fatal("percentage cell escaped its column")
		}
		if value == 100 && !strings.HasPrefix(view, "#######") {
			t.Fatal("100% did not fill the fixed-width track")
		}
		if value == 0 && strings.Contains(view, "#") {
			t.Fatal("zero has a filled segment")
		}
		if value == 1.968 && !strings.Contains(view, "1.968%") {
			t.Fatal("bar quantization lost the precise value")
		}
	}
	for _, value := range []float64{math.NaN(), 101, -1} {
		if !strings.Contains(feedback.PercentBar(value, true, 16, theme.New(theme.Mono)), "n/a") {
			t.Fatal("invalid percentage rendered as data")
		}
	}
	if !strings.Contains(feedback.PercentBar(0, false, 16, theme.New(theme.Mono)), "n/a") {
		t.Fatal("missing data rendered as zero")
	}
}

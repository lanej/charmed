package text

import (
	"math"
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

type QuantityScale string

const (
	Unscaled     QuantityScale = ""
	DecimalScale QuantityScale = "decimal"
	BinaryScale  QuantityScale = "binary"
)

// Quantity retains the supplied observation. Scaling and rounding affect its label only.
// Unit is literal text (e.g. B or req/s); scaling is opt-in and precision is clamped to 0..9.
type Quantity struct {
	Value     float64
	State     ValueState
	Unit      string
	Scale     QuantityScale
	Precision int
	Emphasis  Emphasis
	Width     int // Minimum right-aligned cell width; long values never wrap or truncate.
}

func (q Quantity) Label() string {
	if q.State != ValuePresent {
		return q.State.Label("")
	}

	if math.IsNaN(q.Value) || math.IsInf(q.Value, 0) {
		return ValueInvalid.Label("")
	}

	base, prefixes := 1.0, []string{""}

	switch q.Scale {
	case Unscaled:
	case DecimalScale:
		base, prefixes = 1000, []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y"}
	case BinaryScale:
		base, prefixes = 1024, []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi", "Yi"}
	default:
		return "invalid scale"
	}

	precision := max(0, min(9, q.Precision))
	factor := math.Pow10(precision)

	value, index := q.Value, 0
	for index+1 < len(prefixes) && math.Abs(value) >= base {
		value /= base
		index++
	}

	if index+1 < len(prefixes) && math.Round(math.Abs(value)*factor)/factor >= base {
		value /= base
		index++
	}

	if value == 0 {
		value = 0
	}

	label := strconv.FormatFloat(value, 'f', precision, 64)
	if value != 0 && strings.Trim(label, "-0.") == "" {
		label = strconv.FormatFloat(value, 'g', 3, 64)
	}

	unit := prefixes[index] + Inline(q.Unit)
	if unit != "" {
		if q.Unit != "" {
			label += " "
		}

		label += unit
	}

	return label
}

func (q Quantity) View(t theme.Theme) string {
	label := q.Label()
	return ToneStyle(Neutral, q.Emphasis, t).Width(max(q.Width, ansi.StringWidth(label))).AlignHorizontal(lipgloss.Right).Render(label)
}

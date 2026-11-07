package text

import (
	"fmt"
	"math"

	"github.com/lanej/charmed/theme"
)

// SignTones is application policy: positive can mean improvement or regression.
// Unspecified positive/negative tones default to success/danger respectively.
type SignTones struct{ Positive, Negative, Zero Tone }

// SignedNumber keeps exact sign and units when color is unavailable. Zero defaults to
// neutral; missing and non-finite observations are unavailable, never zero.
type SignedNumber struct {
	Value     float64
	Available bool
	Precision int
	Unit      string
	Tones     SignTones
	Emphasis  Emphasis
}

func (n SignedNumber) valid() bool {
	return n.Available && !math.IsNaN(n.Value) && !math.IsInf(n.Value, 0)
}
func (n SignedNumber) Label() string {
	if !n.valid() {
		return "n/a"
	}
	value := n.Value
	if value == 0 {
		value = 0
	} // Canonicalize IEEE negative zero.
	precision := max(0, min(9, n.Precision))
	if value != 0 && math.Abs(value) < .5*math.Pow10(-precision) {
		return fmt.Sprintf("%+.3g%s", value, n.Unit)
	}
	return fmt.Sprintf("%+.*f%s", precision, value, n.Unit)
}
func (n SignedNumber) View(t theme.Theme) string {
	tone := n.Tones.Zero
	if !n.valid() {
		tone = Neutral
	}
	if n.valid() && n.Value > 0 {
		tone = n.Tones.Positive
		if tone == "" {
			tone = Success
		}
	} else if n.valid() && n.Value < 0 {
		tone = n.Tones.Negative
		if tone == "" {
			tone = Danger
		}
	}
	return ToneStyle(tone, n.Emphasis, t).Render(n.Label())
}

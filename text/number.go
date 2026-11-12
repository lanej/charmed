package text

import (
	"fmt"
	"math"
	"strings"
)

// NumberLabel formats an observation without a leading plus sign. Tiny nonzero
// values and large magnitudes use scientific notation instead of zero or unbounded labels.
func NumberLabel(value float64, precision int, unit string, available bool) string {
	number := SignedNumber{Value: value, Precision: precision, Unit: unit, Available: available}
	if number.valid() && math.Abs(value) >= 1e12 {
		return fmt.Sprintf("%.3g%s", value, unit)
	}

	return strings.TrimPrefix(number.Label(), "+")
}

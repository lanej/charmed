package text

import (
	"fmt"
	"math"
	"strings"
)

// NumberLabel formats an observation without a leading plus sign or redundant
// decimal zeros. Tiny nonzero and very large values retain their magnitude.
func NumberLabel(value float64, precision int, unit string, available bool) string {
	if !available || math.IsNaN(value) || math.IsInf(value, 0) {
		return "n/a"
	}

	return numberDigits(value, precision, false) + unit
}

// AbbreviatedNumber scales base byte units (B and B/s) in binary, and counts,
// rates and base SI units in decimal. Percentages, temperatures, already-prefixed
// and unknown units stay literal. Unavailable values never become zero.
func AbbreviatedNumber(value float64, precision int, unit string, available bool) string {
	if !available || math.IsNaN(value) || math.IsInf(value, 0) {
		return "n/a"
	}

	scale := Unscaled

	switch unit {
	case "B", "B/s":
		scale = BinaryScale
	case "", "req", "req/s", "ops", "ops/s", "Hz", "W", "Wh", "J", "Pa":
		scale = DecimalScale
	}

	if scale == Unscaled {
		return NumberLabel(value, precision, unit, true)
	}

	return (Quantity{Value: value, Precision: precision, Unit: unit, Scale: scale}).Label()
}

func numberDigits(value float64, precision int, signed bool) string {
	if value == 0 {
		value = 0
	}

	precision = max(0, min(9, precision))

	format := "%.*f"
	if signed {
		format = "%+.*f"
	}

	label := fmt.Sprintf(format, precision, value)
	if value != 0 && (math.Abs(value) < .5*math.Pow10(-precision) || math.Abs(value) >= 1e12 || strings.Trim(label, "+-0.") == "") {
		format = "%.3g"
		if signed {
			format = "%+.3g"
		}

		label = fmt.Sprintf(format, value)
	}

	return trimDecimalZeros(label)
}

func trimDecimalZeros(label string) string {
	exponent := ""
	if index := strings.IndexAny(label, "eE"); index >= 0 {
		label, exponent = label[:index], label[index:]
	}

	if strings.Contains(label, ".") {
		label = strings.TrimSuffix(strings.TrimRight(label, "0"), ".")
	}

	return label + exponent
}

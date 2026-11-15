package text_test

import (
	"math"
	"testing"

	"github.com/lanej/charmed/text"
)

func TestAbbreviatedNumbersKeepUnitsAndUsefulDigits(t *testing.T) {
	for _, tc := range []struct {
		value      float64
		precision  int
		unit, want string
	}{
		{20, 3, "kWh", "20kWh"},
		{1.5, 3, "C", "1.5C"},
		{88, 3, "%", "88%"},
		{1258291, 1, "B", "1.2 MiB"},
		{2048, 3, "B/s", "2 KiB/s"},
		{1250, 2, "req/s", "1.25 kreq/s"},
		{-1250, 2, "ops/s", "-1.25 kops/s"},
		{5000, 3, "Hz", "5 kHz"},
		{1200, 2, "", "1.2k"},
		{1200, 3, "°C", "1200°C"},
		{1000, 3, "%", "1000%"},
		{1023.96, 1, "B", "1 KiB"},
		{0.000001, 3, "ms", "1e-06ms"},
		{1e-100, 3, "CO2", "1e-100CO2"},
		{0, 3, "pp", "0pp"},
	} {
		if got := text.AbbreviatedNumber(tc.value, tc.precision, tc.unit, true); got != tc.want {
			t.Fatalf("%g %s: got %q, want %q", tc.value, tc.unit, got, tc.want)
		}
	}

	for _, value := range []float64{math.NaN(), math.Inf(1)} {
		if got := text.AbbreviatedNumber(value, 3, "B", true); got != "n/a" {
			t.Fatalf("nonfinite quantity: %q", got)
		}
	}

	if text.AbbreviatedNumber(0, 3, "%", false) != "n/a" {
		t.Fatal("missing observation became zero")
	}

	signed := text.SignedNumber{Value: 2048, Precision: 3, Unit: "B", Available: true, Abbreviate: true}
	if signed.Label() != "+2 KiB" {
		t.Fatal("abbreviation lost explicit sign")
	}

	if text.NumberLabel(1200, 3, "B", true) != "1200B" {
		t.Fatal("unscaled helper changed unit scale")
	}
}

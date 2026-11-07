package attention_test

import (
	"math"
	"testing"

	"github.com/lanej/charmed/attention"
)

func TestBoundariesAreOptionalAndInclusive(t *testing.T) {
	rule := attention.Rule{Warn: &attention.Bounds{Lower: attention.Limit(-2), Upper: attention.Limit(2)}, Error: &attention.Bounds{Lower: attention.Limit(-3), Upper: attention.Limit(3)}}
	if err := rule.Validate(); err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		value float64
		level attention.Level
	}{{0, attention.Normal}, {2, attention.Warn}, {-2, attention.Warn}, {3, attention.Error}, {-3, attention.Error}, {8, attention.Error}} {
		if got := rule.Evaluate(test.value, true); got.Level != test.level || !got.Available {
			t.Fatalf("%v: %+v", test.value, got)
		}
	}
	for _, value := range []float64{math.NaN(), math.Inf(1)} {
		if rule.Evaluate(value, true).Level != attention.Unknown {
			t.Fatal("nonfinite value classified as normal")
		}
	}
	if rule.Evaluate(0, false).Level != attention.Unknown {
		t.Fatal("missing measurement classified as normal")
	}
	if (attention.Rule{}).Evaluate(999, false).Level != attention.Off {
		t.Fatal("no rule implied attention or health")
	}
	upper := attention.Rule{Error: &attention.Bounds{Upper: attention.Limit(0)}}
	if upper.Evaluate(-1, true).Level != attention.Normal || upper.Evaluate(0, true).Level != attention.Error {
		t.Fatal("one-sided zero boundary was lost")
	}
	same := attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(3)}, Error: &attention.Bounds{Upper: attention.Limit(3)}}
	if same.Evaluate(3, true).Level != attention.Error {
		t.Fatal("error did not take precedence")
	}
}
func TestInvalidBoundariesAndOwnership(t *testing.T) {
	for _, r := range []attention.Rule{
		{Warn: &attention.Bounds{Upper: attention.Limit(math.Inf(1))}},
		{Warn: &attention.Bounds{Lower: attention.Limit(2), Upper: attention.Limit(1)}},
		{Warn: &attention.Bounds{Upper: attention.Limit(3)}, Error: &attention.Bounds{Upper: attention.Limit(2)}},
		{Warn: &attention.Bounds{Lower: attention.Limit(-3)}, Error: &attention.Bounds{Lower: attention.Limit(-2)}},
	} {
		if r.Validate() == nil {
			t.Fatal("invalid rule accepted")
		}
	}
	rule := attention.Rule{Warn: &attention.Bounds{Upper: attention.Limit(2)}}
	copy := rule.Clone()
	*rule.Warn.Upper = 50
	if copy.Evaluate(3, true).Level != attention.Warn {
		t.Fatal("cloned policy aliases caller data")
	}
	if attention.Worst(attention.Off, attention.Normal, attention.Unknown, attention.Warn, attention.Error) != attention.Error {
		t.Fatal("wrong attention priority")
	}
	if attention.Worst(attention.Normal, attention.Unknown) != attention.Unknown {
		t.Fatal("unknown was masked by normal")
	}
}

// Package attention evaluates optional numeric boundaries without UI or I/O.
package attention

import (
	"fmt"
	"math"
	"strings"
)

type Level string

const (
	Off     Level = "off"
	Normal  Level = "normal"
	Unknown Level = "unknown"
	Warn    Level = "warn"
	Error   Level = "error"
)

// Bounds trigger at or beyond a configured limit. Either side may be omitted.
type Bounds struct{ Lower, Upper *float64 }
type Rule struct{ Warn, Error *Bounds }
type Result struct {
	Level     Level   `json:"level"`
	Value     float64 `json:"value"`
	Available bool    `json:"available"`
	Threshold float64 `json:"threshold"`
	Side      string  `json:"side,omitempty"`
}
type Threshold struct {
	Level Level
	Value float64
	Side  string
}
type Changed struct {
	Type      string `json:"type"`
	Metric    string `json:"metric"`
	Scope     string `json:"scope"`
	Reference string `json:"reference"`
	Candidate string `json:"candidate"`
	Previous  Level  `json:"previous"`
	Level     Level  `json:"level"`
}

func Limit(value float64) *float64 { return &value }
func has(b *Bounds) bool           { return b != nil && (b.Lower != nil || b.Upper != nil) }
func (r Rule) Configured() bool    { return has(r.Warn) || has(r.Error) }
func (r Rule) Validate() error {
	for _, b := range []*Bounds{r.Warn, r.Error} {
		if b == nil {
			continue
		}
		for _, v := range []*float64{b.Lower, b.Upper} {
			if v != nil && (math.IsNaN(*v) || math.IsInf(*v, 0)) {
				return fmt.Errorf("attention limits must be finite")
			}
		}
		if b.Lower != nil && b.Upper != nil && *b.Lower >= *b.Upper {
			return fmt.Errorf("lower attention limit must be below upper limit")
		}
	}
	if has(r.Warn) && has(r.Error) {
		if r.Warn.Lower != nil && r.Error.Lower != nil && *r.Error.Lower > *r.Warn.Lower {
			return fmt.Errorf("lower error limit must be at least as extreme as warning")
		}
		if r.Warn.Upper != nil && r.Error.Upper != nil && *r.Error.Upper < *r.Warn.Upper {
			return fmt.Errorf("upper error limit must be at least as extreme as warning")
		}
	}
	return nil
}
func (r Rule) Clone() Rule {
	clone := func(b *Bounds) *Bounds {
		if b == nil {
			return nil
		}
		copy := &Bounds{}
		if b.Lower != nil {
			copy.Lower = Limit(*b.Lower)
		}
		if b.Upper != nil {
			copy.Upper = Limit(*b.Upper)
		}
		return copy
	}
	return Rule{Warn: clone(r.Warn), Error: clone(r.Error)}
}
func (r Rule) Thresholds() []Threshold {
	var limits []Threshold
	for _, item := range []struct {
		level  Level
		bounds *Bounds
	}{{Warn, r.Warn}, {Error, r.Error}} {
		if item.bounds == nil {
			continue
		}
		if item.bounds.Lower != nil {
			limits = append(limits, Threshold{Level: item.level, Value: *item.bounds.Lower, Side: "lower"})
		}
		if item.bounds.Upper != nil {
			limits = append(limits, Threshold{Level: item.level, Value: *item.bounds.Upper, Side: "upper"})
		}
	}
	return limits
}

// Evaluate uses a validated rule. No rule is off; unavailable measurements are
// unknown. Error takes precedence, including when warn and error limits coincide.
func (r Rule) Evaluate(value float64, available bool) Result {
	if !r.Configured() {
		return Result{Level: Off}
	}
	if !available || math.IsNaN(value) || math.IsInf(value, 0) {
		return Result{Level: Unknown}
	}
	result := Result{Level: Normal, Value: value, Available: true}
	for _, level := range []Level{Error, Warn} {
		for _, t := range r.Thresholds() {
			if t.Level == level && ((t.Side == "lower" && value <= t.Value) || (t.Side == "upper" && value >= t.Value)) {
				result.Level, result.Threshold, result.Side = level, t.Value, t.Side
				return result
			}
		}
	}
	return result
}
func Priority(level Level) int {
	switch level {
	case Error:
		return 4
	case Warn:
		return 3
	case Unknown:
		return 2
	case Normal:
		return 1
	}
	return 0
}
func Worst(levels ...Level) Level {
	result := Off
	for _, level := range levels {
		if Priority(level) > Priority(result) {
			result = level
		}
	}
	return result
}
func (r Rule) Summary(unit string) string {
	if !r.Configured() {
		return "off"
	}
	var parts []string
	for _, t := range r.Thresholds() {
		label, sign := "W", ">="
		if t.Level == Error {
			label = "E"
		}
		if t.Side == "lower" {
			sign = "<="
		}
		parts = append(parts, fmt.Sprintf("%s%s%+g", label, sign, t.Value))
	}
	return strings.Join(parts, " ") + " " + unit
}

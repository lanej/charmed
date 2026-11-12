package text

import (
	"image/color"
	"strconv"
	"strings"
	"time"

	"github.com/lanej/charmed/theme"
)

// Pending renders a stable glyph by default. Supplying a frame opts into a
// three-cell ellipsis; the caller owns ticks, pause policy, and reduced motion.
func Pending(t theme.Theme, frame ...uint) string {
	return PendingWithPresentation(GlyphOnly, t, frame...)
}

// PendingWithPresentation offers labeled variants without changing the compact default.
func PendingWithPresentation(mode Presentation, t theme.Theme, frame ...uint) string {
	label := t.Glyph(theme.GlyphPending)
	if len(frame) > 0 && !t.ReducedMotion {
		label = strings.Repeat(".", int(frame[0]%3)+1)
	}

	style := t.Control()
	if t.Pending != nil && t.Mode != theme.Mono {
		style = style.Foreground(t.Pending)
	}

	if mode == GlyphOnly {
		return style.Width(3 + 2*theme.InlinePadding).AlignHorizontal(0.5).Render(label)
	}

	return style.Render(mode.Format(label, "PENDING"))
}

func Log(label string, t theme.Theme, options ...Presentation) string {
	return observation(theme.GlyphLog, label, t.HTTP.Read, t, options...)
}

func Trace(label string, t theme.Theme, options ...Presentation) string {
	return observation(theme.GlyphTrace, label, t.HTTP.Change, t, options...)
}

func Span(label string, t theme.Theme, options ...Presentation) string {
	return observation(theme.GlyphSpan, label, t.HTTP.Create, t, options...)
}

// Upstream and Downstream describe direction relative to the caller's service.
func Upstream(label string, t theme.Theme, options ...Presentation) string {
	return observation(theme.GlyphUpstream, label, t.Info, t, options...)
}

func Downstream(label string, t theme.Theme, options ...Presentation) string {
	return observation(theme.GlyphDownstream, label, t.Info, t, options...)
}

func observation(glyph theme.Glyph, label string, foreground color.Color, t theme.Theme, options ...Presentation) string {
	style := t.Text()
	if t.Mode != theme.Mono && foreground != nil {
		style = style.Foreground(foreground)
	}

	return style.Render(presentation(options).Format(t.Glyph(glyph), label))
}

// TimeOfDay retains milliseconds and the supplied timestamp's location.
// It is a display abbreviation; retain the full timestamp in inspection details.
func TimeOfDay(at time.Time) string {
	if at.IsZero() {
		return "--:--:--.---"
	}

	return at.Format("15:04:05.000")
}

// ShortDuration rounds for compact display, without changing the observation.
// Values under one microsecond retain nanoseconds; negative values retain sign.
func ShortDuration(value time.Duration) string {
	var resolution time.Duration

	switch {
	case value > -time.Microsecond && value < time.Microsecond:
		resolution = time.Nanosecond
	case value > -time.Millisecond && value < time.Millisecond:
		resolution = time.Microsecond
	case value > -time.Second && value < time.Second:
		resolution = time.Millisecond
	case value > -time.Minute && value < time.Minute:
		resolution = 10 * time.Millisecond
	default:
		resolution = time.Second
	}

	return value.Round(resolution).String()
}

// Age is a compact elapsed time with one unit: 14s, 18m, 3h or 2d, truncated rather than rounded so an age never
// reads fresher than it is. A negative duration is a clock disagreement; callers that can see both times label it
// as future (as Freshness does) rather than passing it here.
func Age(value time.Duration) string {
	value = max(value, 0)

	switch {
	case value < time.Minute:
		return strconv.Itoa(int(value/time.Second)) + "s"
	case value < time.Hour:
		return strconv.Itoa(int(value/time.Minute)) + "m"
	case value < 24*time.Hour:
		return strconv.Itoa(int(value/time.Hour)) + "h"
	default:
		return strconv.Itoa(int(value/(24*time.Hour))) + "d"
	}
}

package chart

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// BucketState describes interval coverage, independently of the observed count.
// The zero value is unknown; only explicitly complete buckets receive a bar.
type BucketState string

const (
	BucketUnknown  BucketState = ""
	BucketComplete BucketState = "complete"
	BucketPartial  BucketState = "partial"
	BucketMissing  BucketState = "missing"
)

// CountBucket is an application-supplied interval. Label should identify its
// bounds/timezone when ambiguous; the chart never infers coverage from a clock.
type CountBucket struct {
	Label string
	Count uint64
	State BucketState
}

// BucketBars compares counts in equal-duration complete intervals. Partial
// intervals retain their exact count as "so far" without a bar or extrapolation;
// unknown and missing intervals never plot as zero. Order is supplied by the
// caller. Use separate charts for different units or interval durations.
type BucketBars struct{ Buckets []CountBucket }

// View scales complete bars to the largest complete count in this view. Exact
// counts remain visible, including zero and values too small to receive a cell.
// Below the space needed for four bar cells, rows wrap without bars; no bucket
// is dropped. A one-column allocation returns "!" because wide graphemes cannot
// fit. Height follows content, so the parent supplies scrolling.
func (b BucketBars) View(width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	if width == 1 {
		return "!"
	}

	if len(b.Buckets) == 0 {
		return t.Subtle().Render(ansi.Hardwrap("No bucket observations", width, true))
	}

	labelWidth, countWidth := 0, 1

	var peak uint64

	for _, bucket := range b.Buckets {
		labelWidth = max(labelWidth, ansi.StringWidth(text.Inline(bucket.Label)))

		countWidth = max(countWidth, len(strconv.FormatUint(bucket.Count, 10)))
		if bucket.State == BucketComplete {
			peak = max(peak, bucket.Count)
		}
	}

	barWidth := width - labelWidth - countWidth - 4

	rows := make([]string, 0, len(b.Buckets))
	for _, bucket := range b.Buckets {
		label := text.Inline(bucket.Label)

		value := strconv.FormatUint(bucket.Count, 10)
		switch bucket.State {
		case BucketComplete:
			if barWidth >= 4 {
				cells := 0
				if peak > 0 {
					cells = min(barWidth, int(float64(bucket.Count)/float64(peak)*float64(barWidth)))
				}

				bar := t.SeriesStyle(0).Render(strings.Repeat(t.Glyph(theme.GlyphBarFill), cells))
				label += strings.Repeat(" ", labelWidth-ansi.StringWidth(label))
				value = strings.Repeat(" ", countWidth-len(value)) + value
				rows = append(rows, label+"  "+bar+strings.Repeat(" ", barWidth-cells)+"  "+value)

				continue
			}
		case BucketPartial:
			value += " so far (partial)"
		case BucketMissing:
			value = "missing"
		case BucketUnknown:
			value = "unknown"
		default:
			value = "invalid coverage"
		}

		rows = append(rows, ansi.Hardwrap(ansi.Wrap(label+"  "+value, width, ""), width, true))
	}

	return strings.Join(rows, "\n")
}

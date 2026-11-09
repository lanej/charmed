package collection

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"mime"
	"strings"
	"unicode/utf8"

	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type CaptureState string

const (
	NotCaptured CaptureState = ""
	Captured    CaptureState = "captured"
	Truncated   CaptureState = "truncated"
	Redacted    CaptureState = "redacted"
)

// HTTPBody owns copied captured bytes. Encoding describes these bytes, which may
// differ from the wire Content-Encoding if the caller supplied decoded content.
type HTTPBody struct {
	State                 CaptureState
	Bytes                 []byte
	ContentType, Encoding string
}

// JSONFoldable permits structural parsing only for complete, declared JSON in
// supported text encodings within the preview budget. Syntax is validated by
// JSONViewer; malformed JSON therefore keeps a visible parsing diagnostic.
func (b HTTPBody) JSONFoldable(limit int) bool {
	if b.State != Captured || len(b.Bytes) == 0 || len(b.Bytes) > max(1, limit) || !utf8.Valid(b.Bytes) {
		return false
	}

	if b.Encoding != "" && !strings.EqualFold(b.Encoding, "identity") {
		return false
	}

	_, params, err := mime.ParseMediaType(b.ContentType)
	charset := strings.ToLower(params["charset"])

	return err == nil && text.ContentLanguage(b.ContentType) == "json" && (charset == "" || charset == "utf-8" || charset == "us-ascii")
}

func (b HTTPBody) Validate() error {
	switch b.State {
	case NotCaptured, Captured, Truncated, Redacted:
		return nil
	default:
		return fmt.Errorf("invalid body capture state %q", b.State)
	}
}

func (b HTTPBody) View(raw bool, limit int, t theme.Theme) string {
	switch b.State {
	case NotCaptured:
		return t.Subtle().Render("Body not captured")
	case Redacted:
		return t.Subtle().Render("Body redacted by capture source")
	case Captured, Truncated:
	default:
		return t.Subtle().Render("Unknown body capture state")
	}

	if len(b.Bytes) == 0 {
		if b.State == Truncated {
			return "Empty captured prefix [capture truncated]"
		}

		return t.Subtle().Render("Empty body (0 bytes)")
	}

	preview := b.Bytes[:min(len(b.Bytes), max(1, limit))]
	if utf8.Valid(b.Bytes) && len(preview) < len(b.Bytes) {
		for len(preview) > 0 && !utf8.Valid(preview) {
			preview = preview[:len(preview)-1]
		}
	}

	notices := []string{}
	if b.State == Truncated {
		notices = append(notices, "[capture truncated]")
	}

	if len(preview) < len(b.Bytes) {
		notices = append(notices, fmt.Sprintf("[preview truncated: %d of %d bytes]", len(preview), len(b.Bytes)))
	}

	if b.Encoding != "" && !strings.EqualFold(b.Encoding, "identity") {
		return bodySummary(preview, len(b.Bytes), "encoding "+text.Inline(b.Encoding), notices, raw, t)
	}

	language := text.ContentLanguage(b.ContentType)

	_, params, parseErr := mime.ParseMediaType(b.ContentType)
	if charset := strings.ToLower(params["charset"]); charset != "" && charset != "utf-8" && charset != "us-ascii" {
		return bodySummary(preview, len(b.Bytes), "unsupported charset "+text.Inline(charset), notices, raw, t)
	}

	if language == "" || !utf8.Valid(preview) {
		label := "binary or unsupported content"
		if parseErr != nil {
			label = "missing or invalid content type"
		}

		return bodySummary(preview, len(b.Bytes), label, notices, raw, t)
	}

	source := string(preview)
	if !raw && language == "json" && b.State == Captured && len(preview) == len(b.Bytes) && !jsonIndentTooDeep(preview) {
		var pretty bytes.Buffer
		if err := json.Indent(&pretty, preview, "", "  "); err != nil {
			notices = append(notices, "[invalid JSON: showing captured source]")
		} else {
			formatted := pretty.Bytes()
			if len(formatted) > max(1, limit) {
				formatted = formatted[:max(1, limit)]
				for len(formatted) > 0 && !utf8.Valid(formatted) {
					formatted = formatted[:len(formatted)-1]
				}

				notices = append(notices, "[formatted preview truncated]")
			}

			source = string(formatted)
		}
	}

	if !raw && language == "json" && jsonIndentTooDeep(preview) {
		notices = append(notices, "[JSON formatting depth limit: showing captured source]")
	}

	content := text.Syntax(source, language, t)
	if len(notices) > 0 {
		content += "\n" + t.Subtle().Render(strings.Join(notices, " "))
	}

	return content
}

func bodySummary(preview []byte, total int, label string, notices []string, raw bool, t theme.Theme) string {
	summary := fmt.Sprintf("%d captured bytes · %s", total, label)
	if raw {
		summary += "\n" + t.Text().Render(hex.Dump(preview))
	}

	if len(notices) > 0 {
		summary += "\n" + t.Subtle().Render(strings.Join(notices, " "))
	}

	return summary
}

// WHY: Bound indentation amplification before allocating the pretty buffer.
func jsonIndentTooDeep(source []byte) bool {
	const maxIndentDepth = 64

	depth := 0

	quoted, escaped := false, false
	for _, b := range source {
		if quoted {
			if escaped {
				escaped = false
			} else if b == '\\' {
				escaped = true
			} else if b == '"' {
				quoted = false
			}

			continue
		}

		switch b {
		case '"':
			quoted = true
		case '{', '[':
			depth++
			if depth > maxIndentDepth {
				return true
			}
		case '}', ']':
			depth--
		}
	}

	return false
}

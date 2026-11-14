package collection

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type jsonLine struct {
	source string
	node   int
}

// Refresh reprojects the current document after native sizing or presentation
// changes. Raw/formatted scroll positions are separate. Disabling folds expands
// every container and discards the previous collapsed state.
func (m *JSONViewer) Refresh() {
	m.viewRevision++
	if !m.geometrySyntax {
		m.preparedSyntax, m.fallbackSyntax = nil, ""
	}

	if !m.Folding {
		clear(m.closed)
	}

	offset := m.Model.YOffset()
	if m.Raw != m.wasRaw {
		if m.Raw {
			m.formattedOffset, offset = offset, m.rawOffset
		} else {
			m.rawOffset, offset = offset, m.formattedOffset
		}

		m.wasRaw = m.Raw
	}

	m.renderErr, m.visible, m.rows = nil, nil, map[int]int{}

	body := ""
	if m.parseErr == nil && len(m.nodes) > 0 && !m.Raw {
		body, m.renderErr = m.project()
	}

	if m.Raw || m.Err() != nil || len(m.nodes) == 0 {
		body = m.fallback()
	}

	if m.Prefix != "" {
		body = m.Prefix + "\n" + body
	}

	m.Model.SetContent(body)
	m.Model.SetYOffset(offset)

	if m.lastWidth > 0 && m.lastWidth != m.Model.Width() && m.Interactive() && m.Focused && (m.selected != 0 || offset > 0) {
		m.reveal()
	}

	m.lastWidth = m.Model.Width()
}

func (m *JSONViewer) fallback() string {
	limit := m.Limits.normalized().RenderBytes
	source := text.Escape(string(m.source))

	truncated := len(source) > limit/2
	if truncated {
		source = source[:max(1, limit/2)]
		for len(source) > 0 && !utf8.ValidString(source) {
			source = source[:len(source)-1]
		}
	}

	highlighted := m.fallbackSyntax
	if highlighted == "" {
		m.syntaxBuilds++

		highlighted = text.Syntax(source, "json", m.Theme)
		if len(highlighted) <= limit {
			m.fallbackSyntax = highlighted
		}
	}

	content := ansi.Hardwrap(highlighted, max(1, m.Model.Width()), true)
	if len(content) > limit {
		content = ansi.Hardwrap(source, max(1, m.Model.Width()), true)
	}

	if err := m.Err(); err != nil {
		content += "\n" + m.Theme.Subtle().Render("["+text.Inline(err.Error())+"]")
	}

	if truncated {
		content += "\n" + m.Theme.Subtle().Render("[JSON source preview truncated]")
	}

	return content
}

func (m *JSONViewer) project() (string, error) {
	lines := []jsonLine{}

	budget := m.Limits.normalized().RenderBytes
	if !m.appendNode(0, false, &lines, &budget) {
		return "", errors.New("JSON rendered-output limit: showing captured source")
	}

	source := make([]string, len(lines))
	for i, line := range lines {
		source[i] = line.source
	}

	highlighted := m.preparedSyntax
	if highlighted == nil {
		m.syntaxBuilds++
		rendered := text.Syntax(strings.Join(source, "\n"), "json", m.Theme)

		highlighted = strings.Split(rendered, "\n")
		if len(rendered) <= m.Limits.normalized().RenderBytes {
			m.preparedSyntax = highlighted
		}
	}

	width, gutter := max(1, m.Model.Width()), 0
	if m.Folding {
		gutter = min(4, max(0, width-1))
	}

	row := 0
	if m.Prefix != "" {
		row = lipgloss.Height(m.Prefix)
	}

	output := []string{}
	used := 0

	for i, line := range lines {
		if line.node >= 0 {
			m.visible = append(m.visible, line.node)
			m.rows[line.node] = row
		}

		wrapped := strings.Split(ansi.Hardwrap(highlighted[i], max(1, width-gutter), true), "\n")
		for j, part := range wrapped {
			mark := strings.Repeat(" ", gutter)
			if j == 0 && gutter > 0 {
				mark = m.marker(line.node, gutter)
			}

			output = append(output, mark+part)
			used += len(mark) + len(part) + 1
			row++

			if used > m.Limits.normalized().RenderBytes {
				m.visible, m.rows = nil, nil
				return "", errors.New("JSON rendered-output limit: showing captured source")
			}
		}
	}

	return strings.Join(output, "\n"), nil
}

func (m *JSONViewer) marker(node, width int) string {
	cursor, disclosure := "  ", "  "
	if node >= 0 && m.Focused && node == m.selected {
		cursor = m.Theme.Heading().Render(m.Theme.Glyph(theme.GlyphRight)) + " "
	}

	if node >= 0 && len(m.nodes[node].children) > 0 {
		disclosure = text.Disclosure("", !m.closed[node], m.Theme, m.Symbols) + " "
	}

	return layout.Fit(cursor+disclosure, layout.Size{Width: width, Height: 1})
}

func (m *JSONViewer) appendNode(index int, comma bool, lines *[]jsonLine, budget *int) bool {
	node := m.nodes[index]

	line := strings.Repeat("  ", node.depth)
	if node.keyStart >= 0 {
		line += string(m.source[node.keyStart:node.keyEnd]) + ": "
	}

	suffix := ""
	if comma {
		suffix = ","
	}

	if node.kind != '{' && node.kind != '[' {
		return appendJSONLine(lines, line+string(m.source[node.start:node.end])+suffix, index, budget)
	}

	end, noun := "}", "fields"
	if node.kind == '[' {
		end, noun = "]", "items"
	}

	line += string(node.kind)
	if len(node.children) == 0 {
		return appendJSONLine(lines, line+end+suffix, index, budget)
	}

	if m.Collapsed(index) {
		if len(node.children) == 1 {
			noun = strings.TrimSuffix(noun, "s")
		}

		ellipsis := "…"
		if m.Theme.ASCII {
			ellipsis = "..."
		}

		return appendJSONLine(lines, line+fmt.Sprintf("%s %d %s", ellipsis, len(node.children), noun)+end+suffix, index, budget)
	}

	if !appendJSONLine(lines, line, index, budget) {
		return false
	}

	for position, child := range node.children {
		if !m.appendNode(child, position < len(node.children)-1, lines, budget) {
			return false
		}
	}

	return appendJSONLine(lines, strings.Repeat("  ", node.depth)+end+suffix, -1, budget)
}

func appendJSONLine(lines *[]jsonLine, source string, node int, budget *int) bool {
	*budget -= len(text.Escape(source)) + 1
	if *budget < 0 {
		return false
	}

	*lines = append(*lines, jsonLine{source: source, node: node})

	return true
}

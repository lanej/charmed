package pattern

import (
	"fmt"
	"slices"
	"strings"
	"unicode"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func (r InspectorRecord) hasDetail() bool { return r.Document != nil || r.Detail != nil }

func (r InspectorRecord) row(width int, selected bool, t theme.Theme) string {
	if r.Row != nil {
		return r.Row(width, selected)
	}

	if r.Summary != nil {
		return r.Summary.View(width, t, selected)
	}

	return ""
}

func validateInspectorDocument(record InspectorRecord) error {
	if record.Document == nil {
		return nil
	}

	seen := make(map[string]bool, len(record.Document.Sections))
	for _, section := range record.Document.Sections {
		if strings.TrimSpace(section.ID) == "" || strings.IndexFunc(section.ID, unicode.IsControl) >= 0 {
			return fmt.Errorf("inspector record %q has empty or unsafe section ID %q", record.ID, section.ID)
		}

		if seen[section.ID] {
			return fmt.Errorf("inspector record %q has duplicate section ID %q", record.ID, section.ID)
		}

		seen[section.ID] = true
	}

	return nil
}

func cloneInspectorSummary(summary layout.RecordSummary) layout.RecordSummary {
	summary.Fields = slices.Clone(summary.Fields)
	if summary.Next != nil {
		next := *summary.Next
		summary.Next = &next
	}

	return summary
}

func cloneInspectorRecords(records []InspectorRecord) []InspectorRecord {
	result := slices.Clone(records)
	for index := range result {
		record := &result[index]
		if record.Summary != nil {
			summary := cloneInspectorSummary(*record.Summary)
			record.Summary = &summary
		}

		if record.Document == nil {
			continue
		}

		document := *record.Document
		document.Summary = cloneInspectorSummary(document.Summary)
		document.Sections = slices.Clone(document.Sections)

		for index := range document.Sections {
			section := &document.Sections[index]
			section.Fields = slices.Clone(section.Fields)
			section.Statuses = slices.Clone(section.Statuses)
			section.Timeline = slices.Clone(section.Timeline)
			section.Paragraphs = slices.Clone(section.Paragraphs)
		}

		record.Document = &document
	}

	return result
}

type inspectorDetailAnchor struct {
	sections      []string
	row, fallback int
}

func (m *ListInspector) detailAnchor() inspectorDetailAnchor {
	offset := m.Detail.YOffset()
	anchor := inspectorDetailAnchor{fallback: offset}
	previousStart := -1

	for _, section := range m.detailSections {
		if section.Start > offset {
			break
		}

		// Paired sections occupy the same rows; the first supplies the anchor.
		if section.Start == previousStart {
			continue
		}

		previousStart = section.Start
		anchor.sections = append(anchor.sections, section.ID)
		anchor.row = offset - section.Start
	}

	return anchor
}

func (m *ListInspector) restoreDetail(anchor inspectorDetailAnchor) {
	for index := len(anchor.sections) - 1; index >= 0; index-- {
		for _, section := range m.detailSections {
			if section.ID != anchor.sections[index] {
				continue
			}

			row := 0
			if index == len(anchor.sections)-1 {
				// The separator after a section belongs to its scroll anchor too.
				row = min(anchor.row, max(0, section.End-section.Start))
			}

			m.Detail.SetYOffset(section.Start + row)

			return
		}
	}

	if len(anchor.sections) > 0 && len(m.detailSections) > 0 {
		m.Detail.GotoTop()
		return
	}

	m.Detail.SetYOffset(anchor.fallback)
}

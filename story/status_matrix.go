package story

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/theme"
)

func statusCellDefinition() Definition {
	return Definition{ID: "feedback/status-cell", Title: "Molecules/Feedback/Status cell", Level: "Molecules", Family: "feedback", Description: "A semantic health fill and glyph with version or refined-state text; focus edges preserve health color.", Source: "feedback/status_cell.go", Symbol: "StatusCell", ExampleSource: "docs/snippets/feedback_status_cell.go", Composes: []string{"feedback/status"}, Defaults: DefaultConfig(), Factory: statusCellStory}
}

func statusMatrixDefinition() Definition {
	return Definition{ID: "collection/status-matrix", Title: "Organisms/Collections/Status matrix", Level: "Organisms", Family: "collection", Description: "Stable row/column identities and sparse status cells compose a keyboard-navigable mapping grid with pinned headers. Arrows or h/j/k/l reveal off-screen hosts; absent mappings differ from unknown health.", Source: "collection/status_matrix.go", Symbol: "StatusMatrix", ExampleSource: "docs/snippets/collection_status_matrix.go", Controls: []string{"state"}, States: []string{"ready", "empty", "loading", "error"}, Composes: []string{"collection/data-grid", "feedback/status-cell"}, Defaults: DefaultConfig(), Factory: statusMatrixStory}
}

func statusCellStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()
		lines := []string{"Health + version or refined state", ""}

		for _, status := range feedback.Statuses() {
			label := "v2.8.1"
			if status == feedback.Warn {
				label = "draining"
			}

			cell := feedback.StatusCell{Status: status, Label: label}
			lines = append(lines, cell.View(18, t, false)+" "+cell.View(18, t, true)+" "+string(status))
		}

		return layout.Fit(strings.Join(lines, "\n"), size)
	}), nil
}

func statusMatrixStory(e Environment) (runtime.Scene, error) {
	palette := e.Config.Palette()
	rows := []collection.MatrixAxis{{ID: "public", Label: "lb-public"}, {ID: "internal", Label: "lb-internal"}, {ID: "canary", Label: "lb-canary"}, {ID: "batch", Label: "lb-batch"}}

	for i := range rows {
		rows[i].Label = palette.Glyph(theme.GlyphLoadBalancer) + " " + rows[i].Label
	}

	columns := make([]collection.MatrixAxis, 6)
	for i := range columns {
		id := fmt.Sprintf("host-%02d", i+1)
		columns[i] = collection.MatrixAxis{ID: id, Label: palette.Glyph(theme.GlyphHost) + " " + id}
	}

	entries := []collection.MatrixEntry{
		{RowID: "public", ColumnID: "host-01", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.1"}},
		{RowID: "public", ColumnID: "host-02", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.1"}},
		{RowID: "public", ColumnID: "host-03", Cell: feedback.StatusCell{Status: feedback.Warn, Label: "draining"}},
		{RowID: "public", ColumnID: "host-04", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.0"}},
		{RowID: "internal", ColumnID: "host-02", Cell: feedback.StatusCell{Status: feedback.Error, Label: "unhealthy"}},
		{RowID: "internal", ColumnID: "host-03", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.0"}},
		{RowID: "internal", ColumnID: "host-06", Cell: feedback.StatusCell{Status: feedback.Unknown, Label: "no probe"}},
		{RowID: "canary", ColumnID: "host-04", Cell: feedback.StatusCell{Status: feedback.Running, Label: "warming"}},
		{RowID: "canary", ColumnID: "host-05", Cell: feedback.StatusCell{Status: feedback.Pending, Label: "v2.9.0-rc1"}},
		{RowID: "batch", ColumnID: "host-05", Cell: feedback.StatusCell{Status: feedback.Disabled, Label: "maintenance"}},
		{RowID: "batch", ColumnID: "host-06", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.1"}},
	}
	if e.Config.State == "empty" {
		rows, columns, entries = nil, nil, nil
	}

	m, err := collection.NewStatusMatrix(rows, columns, entries, e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	m.Corner = "Balancer / Host"
	m.Focus()

	return &scene{resize: func(size layout.Size) {
		m.Resize(layout.Size{Width: size.Width, Height: min(len(rows)+2, max(0, size.Height-4))})
	}, update: func(msg tea.Msg) tea.Cmd {
		var cmd tea.Cmd

		m, cmd = m.Update(msg)

		return cmd
	}, view: func(size layout.Size) string {
		if message := stateMessage(e); message != "" && e.Config.State != "empty" {
			return title("Load balancer / host mapping", e) + message
		}

		selected, present := m.Selected()

		detail := "No selection"
		if selected.RowID != "" {
			detail = selected.RowID + " / " + selected.ColumnID + ": " + e.Config.Palette().Glyph(theme.GlyphUnavailable)
			if present {
				detail = fmt.Sprintf("%s / %s: %s | %s", selected.RowID, selected.ColumnID, selected.Cell.Status, selected.Cell.Label)
			}
		}

		return title("Load balancer / host mapping", e) + m.View() + "\n" + layout.Fit(detail, layout.Size{Width: size.Width, Height: 1})
	}}, nil
}

package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionStatusMatrix() (string, error) {
	palette := theme.New(theme.Dark)
	balancer, host := palette.Glyph(theme.GlyphLoadBalancer), palette.Glyph(theme.GlyphHost)
	rows := []collection.MatrixAxis{
		{ID: "public", Label: balancer + " lb-public"},
		{ID: "internal", Label: balancer + " lb-internal"},
	}
	hosts := []collection.MatrixAxis{
		{ID: "host-01", Label: host + " host-01"},
		{ID: "host-02", Label: host + " host-02"},
		{ID: "host-03", Label: host + " host-03"},
	}
	entries := []collection.MatrixEntry{
		{RowID: "public", ColumnID: "host-01", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.1"}},
		{RowID: "public", ColumnID: "host-02", Cell: feedback.StatusCell{Status: feedback.Warn, Label: "draining"}},
		{RowID: "public", ColumnID: "host-03", Cell: feedback.StatusCell{Status: feedback.OK, Label: "v2.8.0"}},
		{RowID: "internal", ColumnID: "host-02", Cell: feedback.StatusCell{Status: feedback.Error, Label: "unhealthy"}},
		{RowID: "internal", ColumnID: "host-03", Cell: feedback.StatusCell{Status: feedback.Unknown, Label: "no probe"}},
	}

	matrix, err := collection.NewStatusMatrix(rows, hosts, entries, layout.Size{Width: 70, Height: 4}, palette)
	if err != nil {
		return "", err
	}

	matrix.Corner = "Balancer / Host"
	matrix.Focus()
	// Forward Update(msg) while focused; the parent owns Tab and Resize(size).
	// SetData retains selection by ID. Selected distinguishes absent mappings.
	// Set ShowHelp to opt into instructions; the embedded DataGrid stays accessible.
	return matrix.View(), nil
}

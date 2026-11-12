package collection_test

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

// A styled cell that wraps must close its style at each line end, or the next column inherits it.
func TestStaticTableWrappedStyleDoesNotTintTheNextColumn(t *testing.T) {
	bg := lipgloss.NewStyle().Background(lipgloss.Color("#223355"))
	pill := "backend · " + bg.Render(" canary ") + " · " + bg.Render(" non-executing ")
	table := collection.StaticTable{
		Columns: []collection.TableColumn{{Title: "HOST"}, {Title: "ROUTE", Min: 20}, {Title: "NOTE"}},
		Rows: [][]collection.TableCell{
			{collection.TableText("web1"), {Styled: pill}, collection.TableText("after")},
		},
	}

	for _, line := range strings.Split(table.View(40, theme.New(theme.Dark)), "\n") {
		if !strings.Contains(line, "\x1b[") {
			continue
		}

		rest := line[strings.LastIndex(line, "\x1b["):]
		if !strings.HasPrefix(rest, "\x1b[m") && !strings.HasPrefix(rest, "\x1b[0m") {
			t.Fatalf("line leaves a style open: %q", line)
		}

		if ansi.StringWidth(line) > 40 {
			t.Fatalf("line is %d cells: %q", ansi.StringWidth(line), ansi.Strip(line))
		}
	}
}

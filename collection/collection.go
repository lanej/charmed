// Package collection provides themed constructors while preserving Bubbles APIs.
package collection

import (
	"charm.land/bubbles/v2/filepicker"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/tree"
	"charm.land/bubbles/v2/viewport"
	"github.com/lanej/charmed/theme"
)

func NewTable(columns []table.Column, rows []table.Row, width, height int, t theme.Theme) table.Model {
	s := table.DefaultStyles()
	s.Header = t.Heading().Padding(0, 1)
	s.Cell = t.Text().Padding(0, 1)
	s.Selected = t.Focus()
	return table.New(table.WithColumns(columns), table.WithRows(rows), table.WithWidth(max(1, width)), table.WithHeight(max(1, height)), table.WithStyles(s), table.WithFocused(true))
}

// NewDenseTable retains the native table API with one cell of column separation.
func NewDenseTable(columns []table.Column, rows []table.Row, width, height int, t theme.Theme) table.Model {
	m := NewTable(columns, rows, width, height, t)
	s := table.DefaultStyles()
	s.Header = t.Heading().Padding(0, 1, 0, 0)
	s.Cell = t.Text().Padding(0, 1, 0, 0)
	s.Selected = t.Focus()
	m.SetStyles(s)
	return m
}

func NewViewport(content string, width, height int) viewport.Model {
	m := viewport.New(viewport.WithWidth(max(1, width)), viewport.WithHeight(max(1, height)))
	m.SetContent(content)
	return m
}

func NewList(items []list.Item, width, height int, t theme.Theme) list.Model {
	d := list.NewDefaultDelegate()
	d.Styles.NormalTitle = t.Text()
	d.Styles.NormalDesc = t.Subtle()
	d.Styles.SelectedTitle = t.Focus().PaddingLeft(1)
	d.Styles.SelectedDesc = t.Subtle().PaddingLeft(1)
	m := list.New(items, d, width, height)
	m.Styles.Title = t.Heading()
	m.DisableQuitKeybindings()
	return m
}

func NewFilePicker(directory string, t theme.Theme) filepicker.Model {
	m := filepicker.New()
	m.CurrentDirectory = directory
	s := filepicker.DefaultStyles()
	s.File = t.Text()
	s.Directory = t.Heading()
	s.Selected = t.Focus()
	s.Cursor = t.Heading()
	m.Styles = s
	return m
}

func NewTree(root *tree.Node, width, height int, t theme.Theme) tree.Model {
	m := tree.New(root, width, height)
	s := tree.DefaultStyles(t.IsDark())
	s.NodeStyle = t.Text()
	s.SelectedNodeStyle = t.Focus()
	s.ParentNodeStyle = t.Heading()
	s.RootNodeStyle = t.Heading()
	s.CursorStyle = t.Heading()
	s.EnumeratorStyle = t.Subtle()
	s.IndenterStyle = t.Subtle()
	m.SetStyles(s)
	if t.ASCII {
		m.SetOpenCharacter("-")
		m.SetClosedCharacter("+")
		m.SetCursorCharacter(">")
	}
	return m
}

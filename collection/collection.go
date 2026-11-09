// Package collection provides themed constructors while preserving Bubbles APIs.
package collection

import (
	"charm.land/bubbles/v2/filepicker"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/tree"
	"charm.land/bubbles/v2/viewport"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func NewTable(columns []table.Column, rows []table.Row, width, height int, t theme.Theme) table.Model {
	s := table.DefaultStyles()
	s.Header = t.Heading().Padding(0, theme.InlinePadding)
	s.Cell = t.Text().Padding(0, theme.InlinePadding)
	// WHY: Table cells already reserve the inset; row focus must not add another.
	s.Selected = t.Selection().Padding(0)

	return table.New(table.WithColumns(columns), table.WithRows(rows), table.WithWidth(max(1, width)), table.WithHeight(max(1, height)), table.WithStyles(s), table.WithFocused(true))
}

// NewDenseTable retains the native table API with one cell of column separation.
func NewDenseTable(columns []table.Column, rows []table.Row, width, height int, t theme.Theme) table.Model {
	m := NewTable(columns, rows, width, height, t)
	s := table.DefaultStyles()
	s.Header = t.Heading().Padding(0, theme.SpaceSmall, 0, 0)
	s.Cell = t.Text().Padding(0, theme.SpaceSmall, 0, 0)
	s.Selected = t.Selection().Padding(0)
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
	d.Styles.NormalTitle = t.Control()
	d.Styles.NormalDesc = t.Subtle().Padding(0, theme.InlinePadding)
	d.Styles.SelectedTitle = t.Focus()
	d.Styles.SelectedDesc = d.Styles.NormalDesc
	m := list.New(items, d, width, height)
	m.Styles.Title = t.Heading()
	m.Styles.StatusBar = m.Styles.StatusBar.Foreground(t.Muted)
	m.Styles.StatusEmpty, m.Styles.StatusBarActiveFilter, m.Styles.StatusBarFilterCount = t.Subtle(), t.Text(), t.Subtle()
	m.Styles.NoItems = m.Styles.NoItems.Foreground(t.Muted)
	m.Styles.ActivePaginationDot = t.Heading().SetString(t.Glyph(theme.GlyphSelected))
	m.Styles.InactivePaginationDot = t.Subtle().SetString(t.Glyph(theme.GlyphUnselected))
	m.Paginator.ActiveDot = m.Styles.ActivePaginationDot.String()
	m.Paginator.InactiveDot = m.Styles.InactivePaginationDot.String()
	m.Styles.ArabicPagination = t.Subtle()
	m.Styles.DividerDot = m.Styles.DividerDot.Foreground(t.Muted)
	m.Help = feedback.NewHelp(width, t)
	m.DisableQuitKeybindings()

	return m
}

func NewFilePicker(directory string, t theme.Theme) filepicker.Model {
	m := filepicker.New()
	m.CurrentDirectory = directory
	s := filepicker.DefaultStyles()
	s.File = t.Control()
	s.Directory = t.Heading().Padding(0, theme.InlinePadding)
	s.Selected = t.Focus()
	s.Cursor = t.Heading()
	m.Styles = s

	return m
}

func NewTree(root *tree.Node, width, height int, t theme.Theme) tree.Model {
	m := tree.New(root, width, height)
	s := tree.DefaultStyles(t.IsDark())
	s.NodeStyle = t.Control()
	s.SelectedNodeStyle = t.Selection()
	s.ParentNodeStyle = t.Heading().Padding(0, theme.InlinePadding)
	s.RootNodeStyle = t.Heading().Padding(0, theme.InlinePadding)
	s.CursorStyle = t.Heading()
	s.EnumeratorStyle = t.Subtle()
	s.IndenterStyle = t.Subtle()
	m.SetStyles(s)
	m.Help = feedback.NewHelp(width, t)
	m.SetOpenCharacter(t.Glyph(theme.GlyphExpanded))
	m.SetClosedCharacter(t.Glyph(theme.GlyphCollapsed))
	m.SetCursorCharacter(t.Glyph(theme.GlyphRight))

	return m
}

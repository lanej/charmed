package pattern

import (
	"fmt"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/interaction"
)

// SetKeyLayout configures navigation across the composed models. Call it before
// per-model overrides; resizing and palette reopening retain native key maps.
func (m *LogInspector) SetKeyLayout(preset interaction.KeyLayout) (tea.Cmd, error) {
	nav, ok := interaction.Navigation(preset)
	if !ok {
		return nil, fmt.Errorf("unknown keyboard layout %q", preset)
	}

	cmd, err := m.Palette.SetKeyLayout(preset)
	if err != nil {
		return nil, err
	}

	tableKeys := &m.HTTP.Table.KeyMap
	tableKeys.LineUp, tableKeys.LineDown = nav.Up, nav.Down
	tableKeys.PageUp, tableKeys.PageDown = nav.PageUp, nav.PageDown
	tableKeys.HalfPageUp, tableKeys.HalfPageDown = nav.HalfPageUp, nav.HalfPageDown
	tableKeys.GotoTop, tableKeys.GotoBottom = nav.Top, nav.Bottom
	treeKeys := &m.Logs.Model.KeyMap
	treeKeys.Up, treeKeys.Down = nav.Up, nav.Down
	treeKeys.PageUp, treeKeys.PageDown = nav.PageUp, nav.PageDown
	treeKeys.HalfPageUp, treeKeys.HalfPageDown = nav.HalfPageUp, nav.HalfPageDown
	treeKeys.GoToTop, treeKeys.GoToBottom = nav.Top, nav.Bottom
	treeKeys.Open, treeKeys.Close = nav.Right, nav.Left
	setViewportNavigation(&m.HTTP.Details, nav)
	setViewportNavigation(&m.Request, nav)
	setViewportNavigation(&m.Response, nav)
	setViewportNavigation(&m.Logs.Details, nav)

	return cmd, nil
}

func setViewportNavigation(model *viewport.Model, nav interaction.NavigationKeys) {
	model.KeyMap.Up, model.KeyMap.Down = nav.Up, nav.Down
	model.KeyMap.Left, model.KeyMap.Right = nav.Left, nav.Right
	model.KeyMap.PageUp, model.KeyMap.PageDown = nav.PageUp, nav.PageDown
	model.KeyMap.HalfPageUp, model.KeyMap.HalfPageDown = nav.HalfPageUp, nav.HalfPageDown
}

package pattern

import (
	"charm.land/bubbles/v2/viewport"
	"github.com/lanej/charmed/collection"
)

type inspectorJSONState struct{ request, response collection.JSONState }

func (m *LogInspector) pruneJSONStates() {
	retained := map[string]bool{}
	for _, id := range m.HTTP.IDs() {
		retained[id] = true
	}

	for id := range m.jsonStates {
		if !retained[id] {
			delete(m.jsonStates, id)
		}
	}
}

func (m *LogInspector) prepareJSONSelection() bool {
	if m.detailID == m.HTTP.SelectedID() {
		return false
	}

	if m.detailID != "" {
		m.jsonStates[m.detailID] = inspectorJSONState{request: m.RequestJSON.Snapshot(), response: m.ResponseJSON.Snapshot()}
	}

	m.RequestJSON.Clear()
	m.ResponseJSON.Clear()
	m.detailID = m.HTTP.SelectedID()

	return true
}

func (m *LogInspector) refreshJSONPayload(viewer *collection.JSONViewer, model *viewport.Model, payload func(int) (string, collection.HTTPBody, bool), fallback func(int) string, reset bool, state collection.JSONState) {
	offset := model.YOffset()

	prefix, body, ok := payload(model.Width())
	if !ok || !body.JSONFoldable(m.HTTP.BodyPreviewBytes) {
		viewer.Clear()
		model.SetContent(fallback(model.Width()))

		if reset {
			model.GotoTop()
		} else {
			model.SetYOffset(offset)
		}

		return
	}

	viewer.Prefix, viewer.Raw, viewer.Folding = prefix, m.HTTP.Raw, m.HTTP.Folding
	viewer.Theme, viewer.Symbols = m.HTTP.Theme, m.HTTP.Symbols
	viewer.Focused = viewer == m.focusedJSON()
	// JSONFoldable bounds input by the HTTP preview budget without changing
	// the caller's independent JSON limits.
	if err := viewer.SetSource(body.Bytes); err != nil {
		// SetSource has already rendered its bounded source and diagnostic.
		return
	}

	if reset {
		viewer.RestoreState(state)
	}
}

func (m *LogInspector) focusedJSON() *collection.JSONViewer {
	switch m.focus.Current() {
	case requestDataPane:
		return m.RequestJSON
	case httpPane:
		return m.ResponseJSON
	default:
		return nil
	}
}

func (m *LogInspector) syncJSONFocus() {
	for _, pane := range []struct {
		id     string
		viewer *collection.JSONViewer
	}{{requestDataPane, m.RequestJSON}, {httpPane, m.ResponseJSON}} {
		focused := pane.id == m.focus.Current()
		if pane.viewer.Focused != focused {
			pane.viewer.Focused = focused
			if pane.viewer.Interactive() {
				pane.viewer.Refresh()
			}
		}
	}
}

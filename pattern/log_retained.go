package pattern

import (
	"strings"

	"github.com/lanej/charmed/layout"
)

// InspectorRevisions identify independent pane presentation changes. Advance the
// affected revision after changing data, focus, scroll, folds or native styles.
type InspectorRevisions struct{ Requests, Details, Logs uint64 }

type inspectorFrame struct {
	revision uint64
	size     layout.Size
	content  string
}

// ViewRetained opts into pane caching; View continues reading native state on
// every call. Revisions and calls are owned by the application's update loop.
func (m *LogInspector) ViewRetained(revisions InspectorRevisions) string {
	if m.retainedOwner != m {
		m.retainedPanes = nil
		m.retainedOwner = m
	}

	m.retainedUsed = map[string]bool{}

	m.retainedRevisions = &revisions
	defer func() {
		m.retainedRevisions = nil
		for id := range m.retainedPanes {
			if !m.retainedUsed[id] {
				delete(m.retainedPanes, id)
			}
		}

		m.retainedUsed = nil
	}()

	return m.View()
}

func (m *LogInspector) paneView(id string, size layout.Size, render func(layout.Size) string) string {
	if m.retainedRevisions == nil {
		return render(size)
	}

	if m.retainedUsed != nil {
		m.retainedUsed[id] = true
	}

	revision := m.retainedRevisions.Requests

	switch id {
	case httpPane:
		revision = m.retainedRevisions.Details
	case logPane:
		revision = m.retainedRevisions.Logs
	}

	if m.retainedPanes == nil {
		m.retainedPanes = map[string]inspectorFrame{}
	}

	frame, ok := m.retainedPanes[id]
	if !ok || frame.revision != revision || frame.size != size {
		frame = inspectorFrame{revision: revision, size: size, content: strings.Clone(render(size))}
		m.retainedPanes[id] = frame
	}

	return frame.content
}

package story

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/examples/cohorts/fixture"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
)

func composedDefinitions() []Definition {
	return []Definition{
		{ID: "collection/scroll-pane", Title: "Molecules/Collections/Scroll pane", Level: "Molecules", Family: "collection", Description: "Bind a native viewport to a stable scrollbar allocation. Content and size updates clamp scrolling while preserving native customization and pointer identity.", Source: "collection/scroll_pane.go", Symbol: "ScrollPane", Composes: []string{"collection/viewport", "feedback/scrollbar"}, Defaults: DefaultConfig(), Factory: scrollPaneStory},
		{ID: "collection/grid-inspector", Title: "Organisms/Collections/Grid inspector", Level: "Organisms", Family: "collection", Description: "Inspect a keyed grid coordinate alongside scrollable details. Narrow allocations stack both panes; very short allocations show the focused pane.", Source: "collection/grid_inspector.go", Symbol: "GridInspector", Composes: []string{"collection/data-grid", "collection/scroll-pane"}, Defaults: DefaultConfig(), Factory: gridInspectorStory},
		{ID: "collection/keyed-table", Title: "Organisms/Collections/Keyed table", Level: "Organisms", Family: "collection", Description: "Retain table selection through filtering and reordering using stable caller keys. Atomic replacement validates identities and row shapes before touching the native table.", Source: "collection/keyed_table.go", Symbol: "KeyedTable", Composes: []string{"collection/table"}, Defaults: DefaultConfig(), Factory: keyedTableStory},
		{ID: "input/radio-group", Title: "Molecules/Input/Radio group", Level: "Molecules", Family: "input", Description: "Choose one option from an ordered group. Arrow keys move focus without selecting, disabled options are skipped, and Enter or Space commits the choice.", Source: "input/radio_group.go", Symbol: "RadioGroup", Composes: []string{"input/radio"}, Defaults: DefaultConfig(), Factory: radioGroupStory},
		{ID: "diagram/trace-replay", Title: "Organisms/Diagrams/Trace replay", Level: "Organisms", Family: "diagram", Description: "Inspect a supplied transition trace by occurrence. Replay stops at either end and derives active state from the inspected prefix without executing transitions.", Source: "diagram/trace_replay.go", Symbol: "TraceReplay", Composes: []string{"diagram/state-transitions", "input/radio-group"}, Defaults: stateDiagramConfig(), Factory: stateDiagramStory},
		{ID: "pattern/measurement-history", Title: "Organisms/Patterns/Measurement history", Level: "Organisms", Family: "pattern", Description: "Inspect one member across metric rows and supplied time windows. Missing members, missing metrics, scalar zero and unavailable ratios retain distinct meaning.", Source: "pattern/measurement_history.go", Symbol: "MeasurementHistory", Composes: []string{"collection/grid-inspector"}, Defaults: DefaultConfig(), Factory: measurementHistoryStory},
		{ID: "git/patch-browser", Title: "Organisms/Git/Patch browser", Level: "Organisms", Family: "git", Description: "Browse every supplied file patch with independent list and diff focus. Unknown inventory, known empty patches, binary content and unavailable diffs remain distinct.", Source: "gitui/patch_browser.go", Symbol: "PatchBrowser", Composes: []string{"git/diff", "collection/scroll-pane"}, Defaults: DefaultConfig(), Factory: patchBrowserStory},
	}
}

func scrollPaneStory(e Environment) (runtime.Scene, error) {
	native := collection.NewViewport("", 1, 1)
	native.SoftWrap = true
	pane := collection.NewScrollPane(&native, e.Config.Size(), e.Config.Palette())
	pane.SetContent(strings.Repeat("Request observations remain scrollable.\n", 40))

	return &scene{resize: pane.Resize, update: pane.Update, view: func(_ layout.Size) string { return pane.View() }}, nil
}

func gridInspectorStory(e Environment) (runtime.Scene, error) {
	axes := []collection.GridAxis{{ID: "api", Label: "API"}, {ID: "worker", Label: "Workers"}}

	grid, err := collection.NewDataGrid(axes, axes, []collection.GridEntry[int]{{RowID: "api", ColumnID: "api", Value: 24}}, e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	m := collection.NewGridInspector(grid, func(entry collection.GridEntry[int], present bool, _ int) string {
		return fmt.Sprintf("%s / %s\nObserved: %t\nReplicas: %d", entry.RowID, entry.ColumnID, present, entry.Value)
	}, e.Config.Size(), e.Config.Palette())

	return &scene{resize: m.Resize, update: func(msg tea.Msg) tea.Cmd {
		if k, ok := msg.(tea.KeyPressMsg); ok && k.String() == "tab" {
			m.FocusDetail(!m.DetailFocused())
			return nil
		}

		return m.Update(msg)
	}, view: func(_ layout.Size) string { return m.View() }}, nil
}

func keyedTableStory(e Environment) (runtime.Scene, error) {
	native := collection.NewTable([]table.Column{{Title: "Service", Width: 24}, {Title: "Replicas", Width: 12}}, nil, 40, 6, e.Config.Palette())

	m, err := collection.NewKeyedTable(&native, []collection.KeyedRow[string]{{Key: "api", Row: table.Row{"API gateway", "24"}}, {Key: "worker", Row: table.Row{"Workers", "12"}}})
	if err != nil {
		return nil, err
	}

	native.Focus()

	return &scene{resize: func(size layout.Size) { native.SetWidth(size.Width); native.SetHeight(size.Height) }, update: m.Update, view: func(_ layout.Size) string { return m.View() }}, nil
}

func radioGroupStory(e Environment) (runtime.Scene, error) {
	m, err := input.NewRadioGroup([]input.RadioOption{{ID: "last", Label: "Last transition"}, {ID: "three", Label: "Last three transitions"}, {ID: "all", Label: "Complete trace"}, {ID: "live", Label: "Live execution", DisabledReason: "Replay is inspection only"}}, "last", e.Config.Palette())
	if err != nil {
		return nil, err
	}

	m.Focus()

	return &scene{update: m.Update, view: func(_ layout.Size) string { return m.View() }}, nil
}

func measurementHistoryStory(e Environment) (runtime.Scene, error) {
	m, err := pattern.NewMeasurementHistory(fixture.Data(), cohort.MemberKey{CohortID: "west", MemberID: "sensor-4"}, e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	return &scene{resize: m.Resize, update: func(msg tea.Msg) tea.Cmd {
		if k, ok := msg.(tea.KeyPressMsg); ok && k.String() == "tab" {
			m.Inspector.FocusDetail(!m.Inspector.DetailFocused())
			return nil
		}

		return m.Update(msg)
	}, view: func(_ layout.Size) string { return m.View() }}, nil
}

func patchBrowserStory(e Environment) (runtime.Scene, error) {
	m, err := gitui.NewPatchBrowser(gitui.PatchObservation{Available: true, Files: gitDiffs()}, e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	return &scene{resize: m.Resize, update: func(msg tea.Msg) tea.Cmd {
		if k, ok := msg.(tea.KeyPressMsg); ok && k.String() == "tab" {
			m.FocusDetail(!m.DetailFocused())
			return nil
		}

		return m.Update(msg)
	}, view: func(_ layout.Size) string { return m.View() }}, nil
}

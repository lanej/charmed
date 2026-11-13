package story

import (
	"strings"

	"github.com/lanej/charmed/catalog"
)

// ComponentGraph is an audited slice of reusable capabilities. Legacy Composes
// remains catalog metadata; entries outside this slice are not inferred to be leaves.
func ComponentGraph() (*catalog.Graph, error) {
	component := func(id, name, pkg, symbol, file string, level catalog.Level) catalog.Component {
		return catalog.Component{ID: id, Name: name, Package: pkg, Symbol: symbol, Level: level, Coverage: catalog.Partial, Source: catalog.Source{Path: file, Symbol: symbol}}
	}
	components := []catalog.Component{
		component("feedback/canary-membership", "Canary membership", "feedback", "CanaryMembership", "feedback/canary_membership.go", catalog.Atoms),
		component("text/indicator", "Indicator", "text", "Indicator", "text/text.go", catalog.Atoms),
		component("text/pill", "Color pill", "text", "Pill", "text/text.go", catalog.Atoms),
		component("feedback/availability", "Availability", "feedback", "Availability", "feedback/availability.go", catalog.Atoms),
		component("feedback/health", "Health", "feedback", "Health", "feedback/health.go", catalog.Atoms),
		component("feedback/readiness", "Readiness", "feedback", "Readiness", "feedback/readiness.go", catalog.Atoms),
		component("feedback/rollout", "Rollout", "feedback", "Rollout", "feedback/rollout.go", catalog.Atoms),
		component("collection/data-grid", "Data grid", "collection", "DataGrid", "collection/data_grid.go", catalog.Organisms),
		component("collection/status-matrix", "Status matrix", "collection", "StatusMatrix", "collection/status_matrix.go", catalog.Organisms),
		component("feedback/status-cell", "Status cell", "feedback", "StatusCell", "feedback/status_cell.go", catalog.Molecules),
		component("feedback/status", "Status", "feedback", "Status", "feedback/status.go", catalog.Atoms),
		component("collection/grid-inspector", "Grid inspector", "collection", "GridInspector", "collection/grid_inspector.go", catalog.Organisms),
		component("collection/scroll-pane", "Scroll pane", "collection", "ScrollPane", "collection/scroll_pane.go", catalog.Molecules),
		component("collection/viewport", "Viewport", "collection", "NewViewport", "collection/collection.go", catalog.Atoms),
		component("feedback/scrollbar", "Scrollbar", "feedback", "Scrollbar", "feedback/controls.go", catalog.Atoms),
		component("pattern/measurement-history", "Measurement history", "pattern", "MeasurementHistory", "pattern/measurement_history.go", catalog.Organisms),
		component("diagram/state-diagram", "State diagram", "diagram", "StateDiagram", "diagram/state.go", catalog.Molecules),
		component("diagram/trace-replay", "Trace replay", "diagram", "TraceReplay", "diagram/trace_replay.go", catalog.Organisms),
		component("input/radio-group", "Radio group", "input", "RadioGroup", "input/radio_group.go", catalog.Molecules),
		component("input/radio", "Radio", "input", "Radio", "input/controls.go", catalog.Atoms),
	}
	relation := func(from, to, role, file, symbol string, kind catalog.RelationKind) catalog.Relation {
		return catalog.Relation{From: from, To: to, Role: role, Kind: kind, Source: catalog.Source{Path: file, Symbol: symbol}}
	}
	relations := []catalog.Relation{
		relation("feedback/canary-membership", "text/indicator", "pool membership presentation", "feedback/canary_membership.go", "CanaryMembership.View", catalog.Composes),
		relation("text/indicator", "text/pill", "labeled pill presentation", "text/text.go", "Indicator", catalog.Composes),
		relation("feedback/status", "text/indicator", "state presentation", "feedback/status.go", "Status.View", catalog.Composes),
		relation("feedback/availability", "text/indicator", "state presentation", "feedback/availability.go", "Availability.View", catalog.Composes),
		relation("feedback/health", "text/indicator", "state presentation", "feedback/health.go", "Health.View", catalog.Composes),
		relation("feedback/readiness", "text/indicator", "state presentation", "feedback/readiness.go", "Readiness.View", catalog.Composes),
		relation("feedback/rollout", "text/indicator", "state presentation", "feedback/rollout.go", "Rollout.View", catalog.Composes),
		relation("collection/status-matrix", "collection/data-grid", "grid base", "collection/status_matrix.go", "StatusMatrix", catalog.Specializes),
		relation("collection/status-matrix", "feedback/status-cell", "cell presentation", "collection/status_matrix.go", "statusGridCell", catalog.Composes),
		relation("feedback/status-cell", "feedback/status", "health presentation", "feedback/status_cell.go", "StatusCell.View", catalog.Composes),
		relation("collection/grid-inspector", "collection/data-grid", "grid", "collection/grid_inspector.go", "GridInspector.Grid", catalog.Composes),
		relation("collection/grid-inspector", "collection/scroll-pane", "details", "collection/grid_inspector.go", "GridInspector.Details", catalog.Composes),
		relation("collection/scroll-pane", "collection/viewport", "content viewport", "collection/scroll_pane.go", "ScrollPane.Model", catalog.Composes),
		relation("collection/scroll-pane", "feedback/scrollbar", "scroll position", "collection/scroll_pane.go", "ScrollPane.View", catalog.Composes),
		relation("pattern/measurement-history", "collection/grid-inspector", "inspection", "pattern/measurement_history.go", "MeasurementHistory.Inspector", catalog.Composes),
		relation("diagram/trace-replay", "diagram/state-diagram", "diagram", "diagram/trace_replay.go", "TraceReplay.graph", catalog.Composes),
		relation("input/radio-group", "input/radio", "options", "input/radio_group.go", "RadioGroup.radios", catalog.Composes),
	}
	functions := map[string]string{
		"feedback/canary-membership": "ExampleFeedbackCanaryMembership",
		"text/indicator":             "ExampleTextIndicator", "text/pill": "ExampleTextPill",
		"feedback/availability": "ExampleFeedbackAvailability",
		"feedback/health":       "ExampleFeedbackHealth",
		"feedback/readiness":    "ExampleFeedbackReadiness",
		"feedback/rollout":      "ExampleFeedbackRollout",
		"collection/data-grid":  "ExampleCollectionDataGrid", "collection/status-matrix": "ExampleCollectionStatusMatrix",
		"feedback/status-cell": "ExampleFeedbackStatusCell", "feedback/status": "ExampleFeedbackStatus",
		"collection/grid-inspector": "ExampleCollectionGridInspector", "collection/scroll-pane": "ExampleCollectionScrollPane",
		"collection/viewport": "ExampleCollectionViewport", "feedback/scrollbar": "ExampleFeedbackScrollbar",
		"pattern/measurement-history": "ExamplePatternMeasurementHistory", "diagram/state-transitions": "ExampleDiagramStateTransitions",
		"diagram/trace-replay": "ExampleDiagramTraceReplay", "input/radio-group": "ExampleInputRadioGroup", "input/radio": "ExampleInputRadio",
	}

	examples := make([]catalog.Example, 0, len(components)+2)
	for _, c := range components {
		entry := c.ID
		if entry == "diagram/state-diagram" {
			entry = "diagram/state-transitions"
		}

		source := catalog.Source{Path: "docs/snippets/" + strings.NewReplacer("/", "_", "-", "_").Replace(entry) + ".go", Symbol: functions[entry]}

		example := catalog.Example{ID: entry + "/usage", EntryID: entry, Primary: c.ID, Kind: catalog.ConsumerExample, Source: source, Uses: []catalog.Use{{Component: c.ID, Role: "primary demonstration", Source: source}}}
		if c.ID == "diagram/trace-replay" {
			example.Uses = append(example.Uses, catalog.Use{Component: "diagram/state-diagram", Role: "graph data", Source: source})
		}

		examples = append(examples, example)
	}

	source := catalog.Source{Path: "story/diagrams.go", Symbol: "stateDiagramStory"}

	for _, entry := range []string{"diagram/trace-replay", "diagram/state-transitions"} {
		primary := "diagram/trace-replay"
		if entry == "diagram/state-transitions" {
			primary = "diagram/state-diagram"
		}

		examples = append(examples, catalog.Example{ID: entry + "/demo", EntryID: entry, Primary: primary, Kind: catalog.InteractiveExample, Source: source, Uses: []catalog.Use{
			{Component: "diagram/state-diagram", Role: "graph data", Source: source},
			{Component: "diagram/trace-replay", Role: "trace inspection", Source: source},
			{Component: "input/radio-group", Role: "history-window controls", Source: source},
		}})
	}

	return catalog.New(catalog.Data{Components: components, Relations: relations, Examples: examples})
}

package story

import (
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func dataAtomDefinitions() []Definition {
	definitions := []Definition{}

	for _, spec := range []struct{ id, title, description, source, symbol string }{
		{"input/checkbox", "Checkbox", "Unchecked, checked and mixed selection with focus, native activation bindings and typed change messages.", "input/controls.go", "NewCheckbox"},
		{"input/radio", "Radio", "Single-choice selection; the owning parent maintains group exclusivity.", "input/controls.go", "NewRadio"},
		{"input/toggle", "Toggle", "Explicit on/off setting with typed changes, disabled reasons and stable state geometry.", "input/controls.go", "NewToggle"},
		{"input/button", "Action button", "Application-owned actions with focus, disabled explanations and an explicit busy state.", "input/controls.go", "NewButton"},
		{"text/value-state", "Value state", "Present, missing, redacted and invalid values remain distinct from empty strings, zero and false.", "text/value.go", "ValueState"},
		{"feedback/freshness", "Freshness", "Explicit observation and current times with a caller-owned stale threshold; future timestamps stay visible.", "feedback/freshness.go", "Freshness"},
		{"text/quantity", "Quantity", "Decimal or binary scaling, literal units, right alignment and small nonzero preservation.", "text/quantity.go", "Quantity"},
		{"text/identifier", "Identifier", "Cell-bounded middle abbreviation keeps the full underlying identifier intact.", "text/identifier.go", "Identifier"},
		{"text/match", "Match highlighting", "Caller-supplied rune ranges highlight entire graphemes while source controls remain escaped data.", "text/match.go", "HighlightMatches"},
		{"text/sort", "Sort indicator", "Explicit unsorted, ascending and descending order with optional sort-key precedence.", "text/markers.go", "SortIndicator"},
		{"text/change", "Change marker", "Added, modified, removed and unchanged categories preserve meaning without color.", "text/markers.go", "ChangeKind"},
		{"text/omission", "Omission marker", "Display-hidden and source-truncated evidence retain known, unknown and zero omitted counts.", "text/markers.go", "Omission"},
		{"feedback/evidence", "Evidence state", "Healthy, degraded, down, unknown, stale, insufficient, truncated and attribution-gap evidence; only healthy reads as OK.", "feedback/evidence.go", "Evidence"},
		{"feedback/outcome", "Operation outcome", "Succeeded, failed, canceled, timed out and unknown results never infer the outcome of pending work.", "feedback/outcome.go", "Outcome"},
	} {
		config := atomConfig()
		config.Rows = 18
		id := spec.id
		family := strings.Split(id, "/")[0]
		familyTitle := map[string]string{"input": "Input", "text": "Text", "feedback": "Feedback"}[family]

		definition := Definition{ID: id, Title: "Atoms/" + familyTitle + "/" + spec.title, Level: "Atoms", Family: family, Description: spec.description, Source: spec.source, Symbol: spec.symbol, Controls: []string{"emphasis"}, Defaults: config, Factory: func(e Environment) (runtime.Scene, error) { return dataAtomStory(id, e) }}
		if family == "input" {
			definition.Factory = func(e Environment) (runtime.Scene, error) { return controlAtomStory(id, e) }
			definition.Controls = []string{"state"}
			definition.States = []string{"ready", "disabled"}
		}

		if id == "feedback/evidence" {
			definition.Composes = []string{"text/pill", "feedback/status"}
		}

		if id == "feedback/outcome" {
			definition.Composes = []string{"text/pill"}
		}

		definitions = append(definitions, definition)
	}

	return definitions
}

func dataAtomStory(id string, e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		palette := e.Config.Palette()
		emphasis := text.Emphasis(e.Config.Emphasis)

		rows := []string{palette.Heading().Render(id + " / " + e.Config.Emphasis), ""}
		switch id {
		case "text/value-state":
			for _, value := range []string{"", "0", "false", "api"} {
				rows = append(rows, text.ValuePresent.View(value, emphasis, palette))
			}

			for _, state := range []text.ValueState{text.ValueMissing, text.ValueRedacted, text.ValueInvalid} {
				rows = append(rows, state.View("hidden observation", emphasis, palette))
			}
		case "feedback/freshness":
			now := time.Date(2026, 11, 10, 12, 0, 0, 0, time.UTC)
			for _, age := range []time.Duration{0, 12 * time.Second, 8 * time.Minute, -time.Second} {
				rows = append(rows, (feedback.Freshness{ObservedAt: now.Add(-age), StaleAfter: time.Minute, Emphasis: emphasis}).View(now, palette))
			}

			rows = append(rows, (feedback.Freshness{Emphasis: emphasis}).View(now, palette))
		case "text/quantity":
			for _, q := range []text.Quantity{
				{Value: 1258291, Unit: "B", Scale: text.BinaryScale, Precision: 1},
				{Value: 842, Unit: "req/s"},
				{Value: 12500, Scale: text.DecimalScale, Precision: 1},
				{Value: .000001, Unit: "s", Precision: 2},
				{Value: 0, Unit: "requests"},
				{State: text.ValueMissing},
			} {
				q.Emphasis, q.Width = emphasis, 16
				rows = append(rows, q.View(palette))
			}
		case "text/identifier":
			identity := text.Identifier{Value: "trace-a17c9dd285e14766-92ef", Emphasis: emphasis}
			rows = append(rows, identity.View(min(size.Width, 30), palette), identity.View(min(size.Width, 16), palette), identity.View(min(size.Width, 8), palette))
		case "text/match":
			normal := text.ToneStyle(text.Neutral, emphasis, palette)
			rows = append(rows, text.HighlightMatches("Layout: timeline", []text.MatchRange{{Start: 0, End: 3}, {Start: 8, End: 9}}, normal, normal.Underline(true)))
			rows = append(rows, text.HighlightMatches("界é sample", []text.MatchRange{{Start: 2, End: 3}}, normal, normal.Underline(true)))
		case "text/sort":
			for _, direction := range []text.SortDirection{text.SortNone, text.SortAscending, text.SortDescending} {
				rows = append(rows, (text.SortIndicator{Direction: direction, Priority: 1, Emphasis: emphasis}).View(palette))
			}
		case "text/change":
			for _, kind := range []text.ChangeKind{text.Unchanged, text.Added, text.Modified, text.Removed} {
				rows = append(rows, kind.View(emphasis, palette))
			}
		case "text/omission":
			for _, o := range []text.Omission{{Count: 24, KnownCount: true}, {}, {Kind: text.SourceTruncated}, {Kind: text.SourceTruncated, Count: 42, KnownCount: true}} {
				o.Emphasis = emphasis
				rows = append(rows, o.View(palette))
			}
		case "feedback/evidence":
			for _, evidence := range feedback.Evidences() {
				rows = append(rows, evidence.View(emphasis, palette)+" "+text.ToneStyle(text.Neutral, text.Dim, palette).Render("status "+string(evidence.Status())))
			}
		case "feedback/outcome":
			for _, outcome := range []feedback.Outcome{feedback.OutcomeUnknown, feedback.Succeeded, feedback.Failed, feedback.Canceled, feedback.TimedOut} {
				rows = append(rows, outcome.View(emphasis, palette))
			}
		}

		return strings.Join(rows, "\n")
	}), nil
}

func controlAtomStory(id string, e Environment) (runtime.Scene, error) {
	palette := e.Config.Palette()
	checkbox := input.NewCheckbox("services", "All services", input.Mixed, palette)
	radio := input.NewRadio("production", "Production", false, palette)
	toggle := input.NewToggle("live", "Live updates", false, palette)
	button := input.NewButton("retry", "Retry", palette)
	button.Icon = theme.GlyphRefresh

	checkbox.Focus()
	radio.Focus()
	toggle.Focus()
	button.Focus()

	disabled := e.Config.State == "disabled"
	checkbox.Disabled, radio.Disabled, toggle.Disabled, button.Disabled = disabled, disabled, disabled, disabled
	message := "Space / Enter activates"

	return &scene{update: func(msg tea.Msg) tea.Cmd {
		var cmd tea.Cmd

		switch id {
		case "input/checkbox":
			checkbox, cmd = checkbox.Update(msg)
		case "input/radio":
			radio, cmd = radio.Update(msg)
		case "input/toggle":
			toggle, cmd = toggle.Update(msg)
		case "input/button":
			button, cmd = button.Update(msg)
		}

		switch event := msg.(type) {
		case input.CheckChanged:
			message = "Changed: " + event.ID
			emit(e, event)
		case input.RadioSelected:
			message = "Selected: " + event.ID
			emit(e, event)
		case input.Toggled:
			message = "Toggled: " + event.ID
			emit(e, event)
		case input.Activated:
			message = "Requested: " + event.ID
			emit(e, event)
		}

		return cmd
	}, view: func(layout.Size) string {
		rows := []string{palette.Heading().Render(id), ""}
		switch id {
		case "input/checkbox":
			rows = append(rows, checkbox.View(), "")
			for _, state := range []input.CheckState{input.Unchecked, input.Checked, input.Mixed} {
				rows = append(rows, input.NewCheckbox("sample", "Services", state, palette).View())
			}

			locked := input.NewCheckbox("locked", "Restricted", input.Checked, palette)
			locked.DisabledReason = "Read-only"
			rows = append(rows, locked.View())
		case "input/radio":
			rows = append(rows, radio.View(), input.NewRadio("staging", "Staging", !radio.Selected, palette).View())
			locked := input.NewRadio("locked", "Development", false, palette)
			locked.DisabledReason = "Unavailable"
			rows = append(rows, locked.View())
		case "input/toggle":
			rows = append(rows, toggle.View(), input.NewToggle("on", "Enabled setting", true, palette).View())
			locked := input.NewToggle("locked", "Locked", false, palette)
			locked.DisabledReason = "Read-only"
			rows = append(rows, locked.View())
		case "input/button":
			rows = append(rows, button.View())
			busy := input.NewButton("busy", "Refreshing", palette)
			busy.Icon = theme.GlyphRefresh
			busy.Busy = true
			rows = append(rows, busy.View())
			locked := input.NewButton("locked", "Delete", palette)
			locked.Icon = theme.GlyphDelete
			locked.DisabledReason = "Read-only"
			rows = append(rows, locked.View())
		}

		return strings.Join(rows, "\n") + footer(message, e)
	}}, nil
}

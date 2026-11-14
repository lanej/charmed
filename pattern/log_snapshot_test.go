package pattern_test

import (
	"testing"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func TestApplyInspectorSnapshotValidationAndFullRefresh(t *testing.T) {
	palette := theme.New(theme.Dark)
	records := []collection.HTTPExchange{{ID: "one", StatusCode: 200}, {ID: "two", StatusCode: 200}}
	logs := []collection.LogRecord{{ID: "log", RequestID: "one"}}

	model, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: palette, Requests: records, Logs: logs})
	if err != nil {
		t.Fatal(err)
	}

	model.Resize(layout.Size{Width: 120, Height: 40})
	model.SelectRequest("one")
	model.Request.SetContent("native presentation")

	changed, err := model.ApplyEntries(records, logs)
	if err != nil || changed || model.Request.GetContent() != "native presentation" {
		t.Fatalf("unchanged snapshot refreshed native presentation: %v", err)
	}

	records[0].StatusCode = 500

	invalid := []collection.LogRecord{{ID: "bad"}, {ID: "bad"}}
	if changed, err = model.ApplyEntries(records, invalid); err == nil || changed {
		t.Fatal("invalid logs applied")
	}

	if selected, ok := model.HTTP.Selected(); !ok || selected.StatusCode != 200 {
		t.Fatal("invalid logs partially applied changed requests")
	}

	changed, err = model.ApplyEntries(records, logs)
	if err != nil || !changed {
		t.Fatalf("changed observation not applied: %v", err)
	}

	fresh, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: palette, Requests: records, Logs: logs})
	if err != nil {
		t.Fatal(err)
	}

	fresh.Resize(layout.Size{Width: 120, Height: 40})
	fresh.SelectRequest("one")

	if model.View() != fresh.View() {
		t.Fatal("changed snapshot differs from fresh full preparation")
	}
}

package collection_test

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func TestApplyHTTPSnapshotValidationOwnershipAndSelection(t *testing.T) {
	m := collection.NewHTTPLog(80, 20, theme.New(theme.Dark))
	records := []collection.HTTPExchange{{ID: "one", RequestBody: collection.HTTPBody{State: collection.Captured, Bytes: []byte("old")}}, {ID: "two"}}

	changed, err := m.ApplyEntries(records)
	if err != nil || !changed {
		t.Fatalf("initial: %v %v", changed, err)
	}

	m.Follow = false
	m.Select("one")
	m.Details.SetContent("native customization")

	changed, err = m.ApplyEntries(records)
	if err != nil || changed || m.Details.GetContent() != "native customization" {
		t.Fatalf("unchanged snapshot refreshed native view: %v", err)
	}

	records[0].RequestBody.Bytes[0] = 'n'

	changed, err = m.ApplyEntries(records)
	if err != nil || !changed || m.SelectedID() != "one" {
		t.Fatalf("changed snapshot: %v", err)
	}

	m.Limit = 1

	invalid := []collection.HTTPExchange{{ID: "bad", RequestBody: collection.HTTPBody{State: collection.CaptureState("invalid")}}, records[1]}
	if changed, err = m.ApplyEntries(invalid); err == nil || changed {
		t.Fatal("evicted invalid input accepted")
	}

	if len(m.Entries()) != 2 {
		t.Fatal("invalid snapshot partially changed retained data")
	}
}

func TestApplyLogSnapshotRetainsImmutableObservations(t *testing.T) {
	m := collection.NewLogTree(80, 20, theme.New(theme.Dark))
	records := []collection.LogRecord{{ID: "one"}}

	changed, err := m.ApplyEntries(records)
	if err != nil || !changed {
		t.Fatalf("initial: %v", err)
	}

	changed, err = m.ApplyEntries(records)
	if err != nil || changed {
		t.Fatalf("unchanged: %v", err)
	}

	if changed, err = m.ApplyEntries([]collection.LogRecord{{ID: "one"}, {ID: "one"}}); err == nil || changed {
		t.Fatal("invalid log snapshot accepted")
	}
}

func TestApplyEmptySnapshotPreparesInitialProjection(t *testing.T) {
	m := collection.NewLogTree(80, 20, theme.New(theme.Dark))

	changed, err := m.ApplyEntries(nil)
	if err != nil || !changed {
		t.Fatalf("initial empty projection: %v", err)
	}

	fresh := collection.NewLogTree(80, 20, theme.New(theme.Dark))
	if err = fresh.SetEntries(nil); err != nil {
		t.Fatal(err)
	}

	if m.View() != fresh.View() {
		t.Fatal("empty snapshot differs from full refresh")
	}

	if changed, err = m.ApplyEntries(nil); changed || err != nil {
		t.Fatal("empty snapshot prepared again")
	}
}

func TestApplyHTTPEqualityCoversEveryObservationField(t *testing.T) {
	kind := reflect.TypeFor[collection.HTTPExchange]()
	for i := range kind.NumField() {
		field := kind.Field(i)
		t.Run(field.Name, func(t *testing.T) {
			model := collection.NewHTTPLog(80, 20, theme.New(theme.Dark))

			initial := collection.HTTPExchange{ID: "one"}
			if err := model.SetEntries([]collection.HTTPExchange{initial}); err != nil {
				t.Fatal(err)
			}

			changed := initial

			value := reflect.ValueOf(&changed).Elem().Field(i)
			switch value.Kind() {
			case reflect.String:
				value.SetString("changed")
			case reflect.Bool:
				value.SetBool(true)
			case reflect.Int, reflect.Int64:
				value.SetInt(1)
			case reflect.Map:
				value.Set(reflect.ValueOf(http.Header{"X-Test": {"changed"}}))
			case reflect.Struct:
				if field.Name == "Time" {
					value.Set(reflect.ValueOf(time.Unix(1, 0)))
				} else {
					value.Set(reflect.ValueOf(collection.HTTPBody{State: collection.Captured, Bytes: []byte("changed")}))
				}
			case reflect.Invalid, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
				t.Fatalf("add equality coverage for new field %s", field.Name)
			default:
				t.Fatalf("unknown reflect kind for %s", field.Name)
			}

			applied, err := model.ApplyEntries([]collection.HTTPExchange{changed})
			if err != nil || !applied {
				t.Fatalf("changed field %s skipped: %v", field.Name, err)
			}

			applied, err = model.ApplyEntries([]collection.HTTPExchange{changed})
			if err != nil || applied {
				t.Fatalf("equal field %s rebuilt: %v", field.Name, err)
			}
		})
	}
}

func TestUnchangedHTTPSnapshotDoesNotFormatRowsOrAllocate(t *testing.T) {
	model := collection.NewHTTPLog(80, 20, theme.New(theme.Dark))
	calls := 0

	model.SetURLFormatter(func(value string) string { calls++; return value })

	records := []collection.HTTPExchange{{ID: "one", URL: "https://example.test/one", RequestBody: collection.HTTPBody{State: collection.Captured, Bytes: []byte("body")}}}
	if err := model.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	calls = 0

	var failed bool

	allocations := testing.AllocsPerRun(10, func() { changed, err := model.ApplyEntries(records); failed = failed || changed || err != nil })

	if failed || calls != 0 {
		t.Fatal("unchanged data prepared row presentation")
	}

	if allocations != 0 {
		t.Fatalf("unchanged complete snapshot allocated %v", allocations)
	}
}

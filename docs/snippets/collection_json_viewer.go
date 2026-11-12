package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionJSONViewer() (string, error) {
	payload := []byte(`{
  "shipment": {"id": "shp_42", "options": {"insurance": true, "currency": "USD"}},
  "rates": [{"carrier": "USPS", "rate": 8.500}, {"carrier": "UPS", "rate": 12.50}],
  "attempt": 1,
  "attempt": 2,
  "sequence": 9007199254740993
}`)

	viewer := collection.NewJSONViewer(64, 30, theme.New(theme.Dark))
	if err := viewer.SetSource(payload); err != nil {
		// The viewer displays a bounded source fallback with the diagnostic.
		return "", err
	}

	// Forward Bubble Tea messages through viewer.Update(msg).
	// Enter folds objects/arrays; arrows navigate; paging scrolls wrapped values.

	// Native viewport customization stays available through viewer.Model.
	// Set viewer.Raw = true, then viewer.Refresh(), to inspect the original source.

	return viewer.View(), nil
}

package collection

import (
	"bytes"
	"maps"
	"reflect"
	"slices"
)

// ApplyEntries applies a changed retained snapshot. It validates all supplied
// records, including the evicted prefix, before changing state. A complete
// snapshot equal to previously validated data needs no repeated validation. True includes
// the first empty projection; later equal data does
// not refresh presentation or resume follow; use SetEntries after native edits.
func (m *HTTPLog) ApplyEntries(records []HTTPExchange) (bool, error) {
	if m.snapshotReady && len(records) <= max(1, m.Limit) && slices.EqualFunc(m.records, records, sameExchange) {
		return false, nil
	}

	if err := ValidateHTTPExchanges(records); err != nil {
		return false, err
	}

	retained := records[max(0, len(records)-max(1, m.Limit)):]
	if m.snapshotReady && (len(m.records) == 0 && len(retained) == 0 || slices.EqualFunc(m.records, retained, sameExchange)) {
		return false, nil
	}

	if err := m.SetEntries(records); err != nil {
		return false, err
	}

	return true, nil
}

// ApplyEntries skips rebuilding equal immutable observations. Fields and nested
// values retain LogRecord's caller-owned immutable contract. True includes the
// first empty projection. Use SetEntries to
// refresh presentation after native edits, even when observations are unchanged.
func (m *LogTree) ApplyEntries(records []LogRecord) (bool, error) {
	if m.snapshotReady && len(records) <= max(1, m.Limit) && (len(m.records) == 0 && len(records) == 0 || reflect.DeepEqual(m.records, records)) {
		return false, nil
	}

	if err := ValidateLogRecords(records); err != nil {
		return false, err
	}

	retained := records[max(0, len(records)-max(1, m.Limit)):]
	if m.snapshotReady && (len(m.records) == 0 && len(retained) == 0 || reflect.DeepEqual(m.records, retained)) {
		return false, nil
	}

	if err := m.SetEntries(records); err != nil {
		return false, err
	}

	return true, nil
}

func sameExchange(a, b HTTPExchange) bool {
	// Time location is observable in the RFC3339 detail timestamp.
	//nolint:staticcheck // QF1009: preserve representation, including the displayed zone, rather than instant equality.
	return a.ID == b.ID && a.RequestID == b.RequestID && a.TraceID == b.TraceID && a.Service == b.Service && a.Time == b.Time && a.Method == b.Method && a.URL == b.URL && a.StatusCode == b.StatusCode && a.Completed == b.Completed && a.Duration == b.Duration && a.DurationKnown == b.DurationKnown && sameHeaders(a.RequestHeaders, b.RequestHeaders) && sameHeaders(a.ResponseHeaders, b.ResponseHeaders) && sameBody(a.RequestBody, b.RequestBody) && sameBody(a.ResponseBody, b.ResponseBody)
}

func sameHeaders(a, b map[string][]string) bool {
	return (a == nil) == (b == nil) && maps.EqualFunc(a, b, func(x, y []string) bool { return (x == nil) == (y == nil) && slices.Equal(x, y) })
}

func sameBody(a, b HTTPBody) bool {
	return a.State == b.State && a.ContentType == b.ContentType && a.Encoding == b.Encoding && (a.Bytes == nil) == (b.Bytes == nil) && bytes.Equal(a.Bytes, b.Bytes)
}

package pattern

import (
	"reflect"
	"slices"

	"github.com/lanej/charmed/collection"
)

// ApplyEntries applies changed observations without treating equal snapshots as
// a presentation refresh. Requests own their captured bytes; logs retain their
// immutable nested-field contract. Validation precedes either pane's mutation.
// Use SetEntries after directly changing native presentation configuration.
func (m *LogInspector) ApplyEntries(requests []collection.HTTPExchange, logs []collection.LogRecord) (bool, error) {
	knownLogs := len(logs) <= m.limit && (len(m.logs) == 0 && len(logs) == 0 || reflect.DeepEqual(m.logs, logs))
	if !knownLogs {
		if err := collection.ValidateLogRecords(logs); err != nil {
			return false, err
		}
	}

	nextLogs := logs[max(0, len(logs)-m.limit):]
	logsChanged := !knownLogs && !reflect.DeepEqual(m.logs, nextLogs)

	requestsChanged, err := m.HTTP.ApplyEntries(requests)
	if err != nil {
		return false, err
	}

	if !requestsChanged && !logsChanged {
		return false, nil
	}

	if logsChanged {
		m.logs = slices.Clone(nextLogs)
	}

	m.correlate()
	m.pruneJSONStates()

	return true, nil
}

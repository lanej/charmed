package collection

import (
	"fmt"
	"slices"
)

const (
	DefaultLogLimit         = 500
	DefaultFieldDepth       = 16
	DefaultFieldNodes       = 1024
	DefaultBodyPreviewBytes = 64 * 1024
	DefaultFieldValueBytes  = 4096
)

// LogRecord adds explicit correlation identities without changing LogEntry.
// Fields are caller-owned immutable observations, including nested values.
type LogRecord struct {
	LogEntry
	ID, RequestID, TraceID, SpanID, ParentSpanID string
}

func validateRecordIDs[T any](records []T, id func(T) string) error {
	seen := make(map[string]bool, len(records))
	for _, record := range records {
		key := id(record)
		if key == "" {
			return fmt.Errorf("record ID must not be empty")
		}

		if seen[key] {
			return fmt.Errorf("duplicate record ID %q", key)
		}

		seen[key] = true
	}

	return nil
}

func retainRecords[T any](records []T, limit int) []T {
	return slices.Clone(records[max(0, len(records)-max(1, limit)):])
}

func upsertRecords[T any](records, incoming []T, id func(T) string) []T {
	out := slices.Clone(records)

	indices := make(map[string]int, len(out))
	for i, record := range out {
		indices[id(record)] = i
	}

	for _, record := range incoming {
		if i, ok := indices[id(record)]; ok {
			out[i] = record
		} else {
			indices[id(record)] = len(out)
			out = append(out, record)
		}
	}

	return out
}

func ValidateLogRecords(records []LogRecord) error {
	return validateRecordIDs(records, func(r LogRecord) string { return r.ID })
}

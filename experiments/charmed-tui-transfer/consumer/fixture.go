package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type fixture struct {
	Clock       time.Time                  `json:"clock"`
	Service     string                     `json:"service"`
	Environment string                     `json:"environment"`
	States      []state                    `json:"states"`
	Diagnostics map[string]json.RawMessage `json:"diagnostics"`
}

type state struct {
	ID               string                     `json:"id"`
	Observation      *observation               `json:"observation"`
	ObservationRef   string                     `json:"observation_ref"`
	ObservationPatch map[string]json.RawMessage `json:"observation_patch"`
	Refresh          refresh                    `json:"refresh"`
}

type refresh struct {
	State       string     `json:"state"`
	RequestedAt *time.Time `json:"requested_at"`
	ConfirmedAt *time.Time `json:"confirmed_at"`
	Error       string     `json:"error"`
}

type observation struct {
	ObservedAt              time.Time  `json:"observed_at"`
	CompletedChecks         *uint64    `json:"completed_checks"`
	ChecksDue               *uint64    `json:"checks_due"`
	IntervalDurationMinutes *int       `json:"interval_duration_minutes"`
	Intervals               []interval `json:"intervals"`
	Regions                 []region   `json:"regions"`
}

type interval struct {
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	Coverage        string    `json:"coverage"`
	CompletedChecks *uint64   `json:"completed_checks"`
}

type region struct {
	Name         string     `json:"name"`
	Availability string     `json:"availability"`
	Health       string     `json:"health"`
	LastVerified *time.Time `json:"last_verified"`
}

// Resolve references recursively and merge only named patch fields. Each result
// is decoded independently so replacing regions cannot mutate the prior state.
func (f fixture) resolve(id string, seen map[string]bool) (*observation, error) {
	if seen[id] {
		return nil, fmt.Errorf("cyclic observation reference %q", id)
	}

	seen[id] = true
	for _, s := range f.States {
		if s.ID != id {
			continue
		}

		if s.ObservationRef == "" {
			return s.Observation, nil
		}

		base, err := f.resolve(s.ObservationRef, seen)
		if err != nil {
			return nil, err
		}

		if base == nil {
			return nil, fmt.Errorf("reference %q has no observation", s.ObservationRef)
		}

		data, err := json.Marshal(base)
		if err != nil {
			return nil, fmt.Errorf("encode reference: %w", err)
		}

		var fields map[string]json.RawMessage
		if decodeErr := json.Unmarshal(data, &fields); decodeErr != nil {
			return nil, fmt.Errorf("decode reference fields: %w", decodeErr)
		}

		for name, value := range s.ObservationPatch {
			fields[name] = value
		}

		data, err = json.Marshal(fields)
		if err != nil {
			return nil, fmt.Errorf("encode patch: %w", err)
		}

		var result observation
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("decode patch: %w", err)
		}

		return &result, nil
	}

	return nil, fmt.Errorf("unknown state %q", id)
}

func readFixture(path string) (fixture, error) {
	// The operator explicitly chooses a local fixture; this program has no elevated privileges.
	data, err := os.ReadFile(path) //nolint:gosec // Local fixture reading is the requested input capability.
	if err != nil {
		return fixture{}, fmt.Errorf("read fixture: %w", err)
	}

	var f fixture
	if err := json.Unmarshal(data, &f); err != nil {
		return f, fmt.Errorf("decode fixture: %w", err)
	}

	if len(f.States) == 0 {
		return f, fmt.Errorf("fixture has no states")
	}

	for i := range f.States {
		obs, err := f.resolve(f.States[i].ID, map[string]bool{})
		if err != nil {
			return f, err
		}

		f.States[i].Observation = obs
	}

	return f, nil
}

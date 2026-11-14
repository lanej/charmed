// Package periodic supplies deterministic observations for the periodic UI probe.
package periodic

import (
	"fmt"
	"slices"
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/collection"
)

type Scenario string

const (
	Unchanged Scenario = "unchanged"
	Metric    Scenario = "metric"
	Request   Scenario = "request"
	Append    Scenario = "append"
	Mixed     Scenario = "mixed"
	Hidden    Scenario = "hidden"
)

// Config bounds the synthetic dataset; it starts no ingestion or redraw loop.
type Config struct {
	Records, Points int
	Scenario        Scenario
	Legacy          bool
	History         bool
}

func (c Config) Validate() error {
	if c.Records != 100 && c.Records != 1000 {
		return fmt.Errorf("records must be 100 or 1000")
	}

	if c.Points != 60 && c.Points != 300 {
		return fmt.Errorf("points must be 60 or 300")
	}

	switch c.Scenario {
	case Unchanged, Metric, Request, Append, Mixed, Hidden:
		return nil
	default:
		return fmt.Errorf("unknown scenario %q", c.Scenario)
	}
}

type Snapshot struct {
	Charts   [4][]chart.TimeSeriesData
	Requests []collection.HTTPExchange
	Logs     []collection.LogRecord
	Health   []collection.GridEntry[int]
}

type Source struct {
	config   Config
	current  Snapshot
	sequence int
}

func NewSource(config Config) (*Source, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	source := &Source{config: config}
	for i := range source.current.Charts {
		points := make([]timeserieslinechart.TimePoint, config.Points)
		for j := range points {
			points[j] = timeserieslinechart.TimePoint{Time: time.Unix(int64(j*5), 0), Value: float64(10 + i + j%7)}
		}

		source.current.Charts[i] = []chart.TimeSeriesData{{ID: fmt.Sprintf("metric-%d", i), Label: fmt.Sprintf("Service %d", i+1), Points: points}}
	}

	for i := range config.Records {
		id := fmt.Sprintf("request-%04d", i)
		source.current.Requests = append(source.current.Requests, collection.HTTPExchange{ID: id, RequestID: id, Method: "GET", URL: "https://example.test/health", StatusCode: 200, Completed: true, Time: time.Unix(int64(i), 0), RequestBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(fmt.Sprintf(`{"request":"%s"}`, id))}, ResponseBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(`{"healthy":true,"values":[1,2,3]}`)}})

		logRequestID := id
		if i >= config.Records-10 {
			logRequestID = fmt.Sprintf("request-%04d", config.Records-1)
		}

		source.current.Logs = append(source.current.Logs, collection.LogRecord{ID: fmt.Sprintf("log-%04d", i), RequestID: logRequestID, LogEntry: collection.LogEntry{Message: "Health observation received"}})
	}

	for r := range 5 {
		for c := range 5 {
			source.current.Health = append(source.current.Health, collection.GridEntry[int]{RowID: fmt.Sprint(r), ColumnID: fmt.Sprint(c)})
		}
	}

	return source, nil
}

func (s *Source) Snapshot() Snapshot { return s.current }

// Next returns immutable snapshots; changed slices are replaced, never mutated.
func (s *Source) Next() Snapshot {
	s.sequence++
	next := s.current
	metric := s.config.Scenario == "metric" || s.config.Scenario == "mixed"
	request := s.config.Scenario == "request" || s.config.Scenario == "mixed" || s.config.Scenario == "hidden"

	appendData := s.config.Scenario == "append" || s.config.Scenario == "mixed"
	if metric || appendData {
		next.Charts[0] = slices.Clone(next.Charts[0])

		points := slices.Clone(next.Charts[0][0].Points)
		if appendData {
			points = append(points, timeserieslinechart.TimePoint{Time: points[len(points)-1].Time.Add(5 * time.Second), Value: float64(10 + s.sequence%7)})
			points = points[len(points)-s.config.Points:]
		} else {
			points[len(points)-1].Value += 1
		}

		next.Charts[0][0].Points = points
	}

	if request {
		next.Requests = slices.Clone(next.Requests)

		next.Requests[len(next.Requests)-1].StatusCode = 200 + s.sequence%2*300
		if s.config.Scenario != "hidden" {
			next.Health = slices.Clone(next.Health)
			next.Health[0].Value = s.sequence % 2
		}
	}

	if appendData {
		next.Logs = slices.Clone(next.Logs)
		next.Logs = append(next.Logs, collection.LogRecord{ID: fmt.Sprintf("appended-%d", s.sequence), RequestID: next.Requests[len(next.Requests)-1].ID, LogEntry: collection.LogEntry{Message: "A new observation arrived"}})
		next.Logs = next.Logs[len(next.Logs)-s.config.Records:]
	}

	s.current = next

	return next
}

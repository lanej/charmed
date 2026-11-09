package collection

import (
	"net"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

// HTTPGroup returns a grouping key without modifying the captured exchange.
// Callbacks must be pure and inexpensive. Empty values display as unavailable.
type HTTPGroup func(HTTPExchange) string

func GroupByService(record HTTPExchange) string { return record.Service }

func GroupByHost(record HTTPExchange) string {
	parsed, err := url.Parse(record.URL)
	if err != nil {
		return ""
	}

	return strings.ToLower(parsed.Hostname())
}

// GroupByPort resolves the conventional HTTP/HTTPS port only for absolute URLs.
func GroupByPort(record HTTPExchange) string {
	parsed, err := url.Parse(record.URL)
	if err != nil || parsed.Hostname() == "" {
		return ""
	}

	if port := parsed.Port(); port != "" {
		number, err := strconv.ParseUint(port, 10, 16)
		if err != nil {
			return ""
		}

		return strconv.FormatUint(number, 10)
	}

	switch strings.ToLower(parsed.Scheme) {
	case "http":
		return "80"
	case "https":
		return "443"
	default:
		return ""
	}
}

func GroupByEndpoint(record HTTPExchange) string {
	host := GroupByHost(record)
	if host == "" {
		return ""
	}

	port := GroupByPort(record)
	if port == "" {
		return host
	}

	return net.JoinHostPort(host, port)
}

// SetTimelineGrouping composes criteria in order, such as service then endpoint.
// No criteria restores the ungrouped view. Source order and selected IDs survive.
func (m *HTTPLog) SetTimelineGrouping(criteria ...HTTPGroup) {
	m.TimelineGroupBy = slices.Clone(criteria)
}

type timelineRow struct {
	record HTTPExchange
	group  string
	header bool
}

func (m *HTTPLog) timelineRows() []timelineRow {
	var groups [][]timelineRow

	var labels []string

	positions := map[string]int{}

	grouped := slices.ContainsFunc(m.TimelineGroupBy, func(group HTTPGroup) bool { return group != nil })
	for _, record := range m.records {
		var key string

		var values []string

		for _, criterion := range m.TimelineGroupBy {
			if criterion == nil {
				continue
			}

			value := criterion(record)
			key += strconv.Itoa(len(value)) + ":" + value

			if value == "" {
				value = "[unavailable]"
			}

			values = append(values, value)
		}

		name := strings.Join(values, " · ")

		index, ok := positions[key]
		if !ok {
			index = len(groups)
			positions[key] = index

			groups = append(groups, nil)
			labels = append(labels, name)
		}

		groups[index] = append(groups[index], timelineRow{record: record, group: name})
	}

	var rows []timelineRow

	for i, group := range groups {
		if grouped {
			rows = append(rows, timelineRow{group: labels[i], header: true})
		}

		rows = append(rows, group...)
	}

	return rows
}

func (m *HTTPLog) timelineRecords() []HTTPExchange {
	var records []HTTPExchange

	for _, row := range m.timelineRows() {
		if !row.header {
			records = append(records, row.record)
		}
	}

	return records
}

func timelineWindow(rows []timelineRow, selected, available int) []timelineRow {
	if available <= 0 {
		return nil
	}

	first := min(max(0, selected-available/2), max(0, len(rows)-available))
	if first > 0 && !rows[first].header && rows[first].group != "" && available > 1 {
		first = max(first, max(0, selected-available+2))
		if !rows[first].header {
			window := []timelineRow{{group: rows[first].group, header: true}}
			return append(window, rows[first:min(len(rows), first+available-1)]...)
		}
	}

	return rows[first:min(len(rows), first+available)]
}

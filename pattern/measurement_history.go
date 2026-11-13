package pattern

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// MeasurementState preserves missing members, missing metrics and unavailable ratios.
type MeasurementState uint8

const (
	MeasurementObserved MeasurementState = iota
	MeasurementMemberAbsent
	MeasurementMetricAbsent
	MeasurementRatioUnavailable
)

// HistoryMeasurement retains the exact supplied interval and raw measurement.
type HistoryMeasurement struct {
	Member      cohort.MemberKey
	Metric      cohort.Metric
	Start, End  time.Time
	Measurement cohort.Measurement
	State       MeasurementState
}

// MeasurementHistory inspects one member across metric rows and supplied windows.
// It copies validated data and never infers sampling cadence or health thresholds.
type MeasurementHistory struct {
	Inspector *collection.GridInspector[HistoryMeasurement]
	data      cohort.Data
	member    cohort.MemberKey
}

func NewMeasurementHistory(data cohort.Data, member cohort.MemberKey, size layout.Size, t theme.Theme) (*MeasurementHistory, error) {
	grid, err := collection.NewDataGrid[HistoryMeasurement](nil, nil, nil, size, t)
	if err != nil {
		return nil, err
	}

	grid.Corner = "Metric / Window"
	grid.RowHeaderWidth = 20
	grid.CellWidth = 14
	grid.RenderCell = func(c collection.GridCell[HistoryMeasurement]) string {
		label := feedback.CellLabel(historyLabel(c.Entry.Value))
		if c.Selected {
			return c.Theme.Focus().Render(layout.Fit(label, layout.Size{Width: c.Width, Height: 1}))
		}

		return layout.Fit(label, layout.Size{Width: c.Width, Height: 1})
	}
	m := &MeasurementHistory{member: member}

	m.Inspector = collection.NewGridInspector(grid, func(e collection.GridEntry[HistoryMeasurement], present bool, _ int) string {
		if !present {
			return "No measurement selected"
		}

		v := e.Value

		detail := fmt.Sprintf("%s / %s\n%s\n%s – %s\n%s", v.Member.CohortID, v.Member.MemberID, v.Metric.Label, v.Start.Format(time.RFC3339Nano), v.End.Format(time.RFC3339Nano), historyLabel(v))
		if v.State == MeasurementMemberAbsent || v.State == MeasurementMetricAbsent {
			return detail
		}

		if v.Metric.Kind == cohort.Scalar {
			return detail + fmt.Sprintf("\nRaw value: %g", v.Measurement.Value)
		}

		return detail + fmt.Sprintf("\nNumerator: %d\nDenominator: %d", v.Measurement.Numerator, v.Measurement.Denominator)
	}, size, t)
	if err := m.SetData(data); err != nil {
		return nil, err
	}

	return m, nil
}

func historyLabel(v HistoryMeasurement) string {
	switch v.State {
	case MeasurementObserved:
		value := v.Measurement.Value
		if v.Metric.Kind == cohort.Ratio {
			value = float64(v.Measurement.Numerator) / float64(v.Measurement.Denominator) * 100
		}

		return fmt.Sprintf("%g%s", value, v.Metric.DisplayUnit())
	case MeasurementMemberAbsent:
		return "Member absent"
	case MeasurementMetricAbsent:
		return "Metric absent"
	case MeasurementRatioUnavailable:
		return "Ratio unavailable"
	default:
		return "Invalid observation"
	}
}

// SetData retains a surviving metric and exact window start; a removed window selects latest.
func (m *MeasurementHistory) SetData(data cohort.Data) error {
	cloned, err := data.Clone()
	if err != nil {
		return err
	}

	m.data = cloned
	m.rebuild()

	return nil
}

func (m *MeasurementHistory) SetMember(member cohort.MemberKey) {
	m.member = member
	m.rebuild()
	m.Inspector.Details.Model.GotoTop()
}

func (m *MeasurementHistory) rebuild() {
	previous, _ := m.Inspector.Selected()
	rows := make([]collection.GridAxis, len(m.data.Metrics))
	columns := make([]collection.GridAxis, len(m.data.Windows))

	entries := make([]collection.GridEntry[HistoryMeasurement], 0, len(rows)*len(columns))
	for i, metric := range m.data.Metrics {
		rows[i] = collection.GridAxis{ID: metric.ID, Label: metric.Label}
	}

	retained := false

	for i, w := range m.data.Windows {
		id := w.Start.UTC().Format(time.RFC3339Nano)
		columns[i] = collection.GridAxis{ID: id, Label: w.Start.Format("15:04:05")}
		retained = retained || id == previous.ColumnID

		for _, metric := range m.data.Metrics {
			value := HistoryMeasurement{Member: m.member, Metric: metric, Start: w.Start, End: w.End, State: MeasurementMemberAbsent}
			for _, sample := range w.Samples {
				if sample.MemberKey != m.member {
					continue
				}

				measurement, present := cohort.MeasurementFor(sample, metric)
				value.Measurement = measurement

				value.State = MeasurementObserved
				if !present {
					value.State = MeasurementMetricAbsent
				} else if metric.Kind == cohort.Ratio && measurement.Denominator == 0 {
					value.State = MeasurementRatioUnavailable
				}

				break
			}

			entries = append(entries, collection.GridEntry[HistoryMeasurement]{RowID: metric.ID, ColumnID: id, Value: value})
		}
	}
	// Axes and cells are generated from validated data; their identities are unique.
	if err := m.Inspector.SetData(rows, columns, entries); err != nil {
		panic(err)
	}

	if !retained && len(columns) > 0 {
		selected, _ := m.Inspector.Selected()
		m.Inspector.Select(selected.RowID, columns[len(columns)-1].ID)
	}
}

func (m *MeasurementHistory) Selected() (HistoryMeasurement, bool) {
	e, ok := m.Inspector.Selected()
	return e.Value, ok
}
func (m *MeasurementHistory) Resize(size layout.Size)    { m.Inspector.Resize(size) }
func (m *MeasurementHistory) Update(msg tea.Msg) tea.Cmd { return m.Inspector.Update(msg) }
func (m *MeasurementHistory) View() string               { return m.Inspector.View() }

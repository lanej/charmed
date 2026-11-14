package periodic

import (
	"errors"
	"fmt"
	"reflect"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

type SnapshotMsg struct{ Snapshot Snapshot }

type Changes struct {
	Charts            [4]bool
	Inspector, Health bool
}

// Dashboard is a developer fixture, not an application ingestion service.
type Dashboard struct {
	config        Config
	charts        [4]chart.TimeSeries
	inspector     *pattern.LogInspector
	grid          *collection.DataGrid[int]
	screen        *tui.Screen
	declaration   *layout.Compiled
	size          layout.Size
	latest        Snapshot
	health        []collection.GridEntry[int]
	palette       theme.Theme
	revision      uint64
	paneRevisions [7]uint64
	hidden        bool
	err           error
	layoutErr     error
	renderErr     error
	LastChanges   Changes
	Views         uint64
}

func NewDashboard(config Config, initial Snapshot) (*Dashboard, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	palette := theme.New(theme.Dark)
	d := &Dashboard{config: config, palette: palette, latest: initial, hidden: config.Scenario == "hidden"}

	var err error

	d.screen, err = tui.NewScreen("header", "chart0", "chart1", "chart2", "chart3", "health", "inspector")
	if err != nil {
		return nil, err
	}

	d.inspector, err = pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: palette, Requests: initial.Requests, Logs: initial.Logs, Limit: config.Records, NoHelp: true, NoHeadline: true})
	if err != nil {
		return nil, err
	}

	if config.History && len(initial.Requests) > 0 {
		d.inspector.SelectRequest(initial.Requests[0].ID)
	}

	for i := range d.charts {
		d.charts[i] = chart.NewTimeSeries(initial.Charts[i], layout.Size{Width: 60, Height: 12}, palette)
	}

	rows, columns := axes("host"), axes("svc")

	d.grid, err = collection.NewDataGrid(rows, columns, initial.Health, layout.Size{Width: 60, Height: 12}, palette)
	if err != nil {
		return nil, err
	}

	d.grid.CellWidth, d.grid.RowHeaderWidth = 8, 10
	d.grid.RenderCell = func(cell collection.GridCell[int]) string {
		if !cell.Present {
			return palette.Text().Foreground(palette.Unknown).Render("— N/A")
		}

		if cell.Entry.Value == 0 {
			return palette.Text().Foreground(palette.Success).Render("● OK")
		}

		return palette.Text().Foreground(palette.Warning).Render("▲ WARN")
	}
	d.health = initial.Health
	d.Resize(layout.Size{Width: 120, Height: 40})

	return d, d.Err()
}

func axes(label string) []collection.GridAxis {
	result := make([]collection.GridAxis, 5)
	for i := range result {
		result[i] = collection.GridAxis{ID: fmt.Sprint(i), Label: fmt.Sprintf("%s-%d", label, i)}
	}

	return result
}

func (d *Dashboard) Init() tea.Cmd                { return nil }
func (d *Dashboard) PresentationRevision() uint64 { return d.revision }
func (d *Dashboard) Err() error                   { return errors.Join(d.err, d.layoutErr, d.renderErr) }

func (d *Dashboard) Apply(snapshot Snapshot) (Changes, error) {
	var changes Changes

	if d.hidden {
		if err := collection.ValidateHTTPExchanges(snapshot.Requests); err != nil {
			return changes, err
		}

		if err := collection.ValidateLogRecords(snapshot.Logs); err != nil {
			return changes, err
		}
	}

	if !d.hidden {
		var err error
		if d.config.Legacy {
			err = d.inspector.SetEntries(snapshot.Requests, snapshot.Logs)
			changes.Inspector = true
		} else {
			changes.Inspector, err = d.inspector.ApplyEntries(snapshot.Requests, snapshot.Logs)
		}

		if err != nil {
			return changes, err
		}

		if changes.Inspector {
			d.paneRevisions[6]++
		}
	}

	for i := range d.charts {
		if d.config.Legacy {
			d.charts[i].SetData(snapshot.Charts[i])
			changes.Charts[i] = true
		} else {
			changes.Charts[i] = d.charts[i].ApplyData(snapshot.Charts[i])
		}

		if changes.Charts[i] {
			d.paneRevisions[i+1]++
		}
	}

	changes.Health = d.config.Legacy || !reflect.DeepEqual(d.health, snapshot.Health)
	if changes.Health {
		if err := d.grid.SetData(axes("host"), axes("svc"), snapshot.Health); err != nil {
			return changes, err
		}

		d.health = snapshot.Health
		d.paneRevisions[5]++
	}

	d.latest = snapshot

	changed := changes.Health || changes.Inspector
	for _, value := range changes.Charts {
		changed = changed || value
	}

	if changed {
		d.revision++
	}

	d.LastChanges = changes

	return changes, nil
}

func (d *Dashboard) Update(message tea.Msg) tea.Cmd {
	switch msg := message.(type) {
	case SnapshotMsg:
		previousErr := d.Err()

		d.LastChanges, d.err = d.Apply(msg.Snapshot)
		if d.err != nil || previousErr != nil {
			d.revision++
		}
	case tea.KeyPressMsg:
		switch msg.String() {
		case "h":
			d.hidden = !d.hidden

			d.declaration = nil
			if !d.hidden {
				_, d.err = d.inspector.ApplyEntries(d.latest.Requests, d.latest.Logs)
				d.paneRevisions[6]++
			}

			d.Resize(d.size)

			d.revision++
			if !d.hidden {
				return d.inspector.Init()
			}
		default:
			if !d.hidden {
				d.paneRevisions[6]++
				d.revision++

				return d.inspector.Update(msg)
			}
		}
	default:
		if !d.hidden {
			d.paneRevisions[6]++
			d.revision++

			return d.inspector.Update(message)
		}
	}

	return nil
}

func (d *Dashboard) Resize(size layout.Size) {
	d.size = size
	if d.declaration == nil {
		cells := []layout.Cell{{Node: &layout.Region{ID: "header", Role: layout.ChromeRegion}, ColumnSpan: 2}, {Node: &layout.Region{ID: "chart0"}, Row: 1}, {Node: &layout.Region{ID: "chart1"}, Column: 1, Row: 1}, {Node: &layout.Region{ID: "chart2"}, Row: 2}, {Node: &layout.Region{ID: "chart3"}, Column: 1, Row: 2}, {Node: &layout.Region{ID: "health"}, Row: 3}}

		hidden := []layout.Omission{}
		if d.hidden {
			cells[5].ColumnSpan = 2

			hidden = append(hidden, layout.Omission{ID: "inspector", Reason: "Hidden inspector retains latest observations without preparing output."})
		} else {
			cells = append(cells, layout.Cell{Node: &layout.Region{ID: "inspector"}, Column: 1, Row: 3})
		}

		grid := &layout.Grid{Columns: []layout.Track{layout.Grow(20), layout.Grow(20)}, Rows: []layout.Track{layout.Fixed(2), layout.Grow(6), layout.Grow(6), layout.Grow(6)}, Cells: cells}

		d.declaration, d.layoutErr = layout.Compile(&layout.Responsive{Name: "periodic", Variants: []layout.Variant{{Name: "dashboard", Root: grid, Hidden: hidden}}})
		if d.layoutErr != nil {
			return
		}
	}

	d.layoutErr = d.screen.ArrangeCompiled(d.declaration, size)
	if d.layoutErr != nil {
		return
	}

	for i := range d.charts {
		if bounds, ok := d.screen.Region(fmt.Sprintf("chart%d", i)); ok {
			d.charts[i].Resize(layout.Size{Width: bounds.Dx(), Height: bounds.Dy()})
		}
	}

	if bounds, ok := d.screen.Region("health"); ok {
		d.grid.Resize(layout.Size{Width: bounds.Dx(), Height: bounds.Dy()})
	}

	if bounds, ok := d.screen.Region("inspector"); ok {
		d.inspector.ResizeViewport(layout.Size{Width: bounds.Dx(), Height: bounds.Dy()})
	}
}

func (d *Dashboard) View() string {
	d.Views++
	if failure := errors.Join(d.err, d.layoutErr); failure != nil {
		return layout.Fit("Periodic probe: "+failure.Error(), d.size)
	}

	views := map[string]tui.RegionView{"header": {Render: func(size layout.Size) string {
		return layout.Fit("Periodic observability · synthetic data\nH toggles inspector · native pane keys remain available", size)
	}}}
	for i := range d.charts {
		views[fmt.Sprintf("chart%d", i)] = tui.RegionView{Revision: d.paneRevisions[i+1], Render: func(layout.Size) string { return d.charts[i].View() }}
	}

	views["health"] = tui.RegionView{Revision: d.paneRevisions[5], Render: func(layout.Size) string { return d.grid.View() }}
	if !d.hidden {
		views["inspector"] = tui.RegionView{Revision: d.paneRevisions[6], Render: func(layout.Size) string { return d.inspector.View() }}
	}

	frame, err := d.screen.RenderRetained(views)
	if err != nil {
		d.renderErr = err
		return layout.Fit("Periodic probe: "+err.Error(), d.size)
	}

	d.renderErr = nil

	return frame.Content
}

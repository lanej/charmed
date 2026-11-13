package story

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func nestedLayoutDefinitions() []Definition {
	config := DefaultConfig()
	config.Columns, config.Rows = 100, 20

	return []Definition{
		{ID: "template/nested-columns", Title: "Templates/Layout/Nested columns", Level: "Templates", Family: "layout", Description: "Place a host list beside an inspection column with separately sized summary and evidence sections. Resize to stack the columns, or retain the inspection column when height is limited.", Source: "story/nested_layouts.go", Symbol: "nestedLayoutStory", Composes: []string{"tui/screen", "layout/flex", "layout/responsive", "feedback/status"}, Defaults: config, Factory: func(e Environment) (runtime.Scene, error) { return nestedLayoutStory(false, e) }},
		{ID: "template/nested-grid", Title: "Templates/Layout/Grid inside a column", Level: "Templates", Family: "layout", Description: "Nest a two-cell metrics grid between summary and evidence sections inside a column. Each leaf receives its own allocation while responsive recipes preserve the selected host context.", Source: "story/nested_layouts.go", Symbol: "nestedLayoutStory", Composes: []string{"tui/screen", "layout/flex", "layout/grid", "layout/responsive", "feedback/status"}, Defaults: config, Factory: func(e Environment) (runtime.Scene, error) { return nestedLayoutStory(true, e) }},
	}
}

func nestedLayout(withGrid bool) (layout.Node, []string) {
	region := func(id string, width, height int) *layout.Region {
		return &layout.Region{ID: id, Min: layout.Size{Width: width, Height: height}}
	}
	hosts, summary, details := region("hosts", 24, 6), region("summary", 31, 4), region("details", 31, 5)
	ids := []string{"hosts", "summary", "details"}
	inspection := &layout.Flex{
		Name: "host-inspection", Axis: layout.Vertical, Gap: 1,
		Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Summary and evidence are read top to bottom within the selected host's column."}},
		Children:   []layout.Slot{{Node: summary, Track: layout.Fixed(5)}, {Node: details, Track: layout.Grow(5)}},
	}

	if withGrid {
		cpu, memory := region("cpu", 14, 5), region("memory", 14, 5)
		metrics := &layout.Grid{Columns: []layout.Track{layout.Grow(14), layout.Grow(14)}, Rows: []layout.Track{layout.Grow(5)}, Gap: layout.Size{Width: 1}, Cells: []layout.Cell{{Node: cpu}, {Node: memory, Column: 1}}}

		ids = append(ids, "cpu", "memory")
		inspection.Children = []layout.Slot{{Node: summary, Track: layout.Fixed(4)}, {Node: metrics, Track: layout.Fixed(5)}, {Node: details, Track: layout.Grow(5)}}
	}

	columns := &layout.Flex{Gap: 1, Children: []layout.Slot{{Node: hosts, Track: layout.Fixed(26)}, {Node: inspection, Track: layout.Grow(31)}}}
	stacked := &layout.Flex{Axis: layout.Vertical, Gap: 1, Children: []layout.Slot{{Node: hosts, Track: layout.Fixed(6)}, {Node: inspection, Track: layout.Grow(0)}}}

	return &layout.Responsive{Name: "nested-workspace", Variants: []layout.Variant{
		{Name: "columns", Min: layout.Size{Width: 80, Height: 18}, Root: columns},
		{Name: "stacked", Min: layout.Size{Height: 24}, Root: stacked},
		{Name: "inspection", Root: inspection, Hidden: []layout.Omission{{ID: "hosts", Reason: "The short allocation keeps the selected host's inspection; the host list returns with more room."}}},
	}}, ids
}

func nestedLayoutStory(withGrid bool, e Environment) (runtime.Scene, error) {
	root, ids := nestedLayout(withGrid)

	screen, err := tui.NewScreen(ids...)
	if err != nil {
		return nil, err
	}

	focused := "hosts"
	showRegions := false

	var arrangeErr error

	return &scene{
		resize: func(size layout.Size) {
			arrangeErr = screen.Arrange(root, size)
			if arrangeErr != nil {
				return
			}

			if _, visible := screen.Region(focused); !visible {
				focused = screen.Plan().Regions()[0].ID
			}
		},
		update: func(msg tea.Msg) tea.Cmd {
			if k, ok := msg.(tea.KeyPressMsg); ok {
				switch k.String() {
				case "ctrl+l":
					showRegions = !showRegions
				case "tab", "shift+tab":
					regions := screen.Plan().Regions()
					for i, r := range regions {
						if r.ID == focused {
							delta := 1
							if k.String() == "shift+tab" {
								delta = -1
							}

							focused = regions[(i+delta+len(regions))%len(regions)].ID

							break
						}
					}
				}
			}

			return nil
		},
		view: func(size layout.Size) string {
			if arrangeErr != nil {
				return e.Config.Palette().Text().Foreground(e.Config.Palette().Danger).Render(arrangeErr.Error())
			}

			if showRegions {
				return screen.RegionMap(e.Config.Palette())
			}

			palette := e.Config.Palette()
			host := palette.Glyph(theme.GlyphHost)
			contents := map[string]string{
				"hosts":   strings.Join([]string{host + " host-01", palette.Heading().Render(host + " host-02"), host + " host-03"}, "\n"),
				"summary": host + " host-02 · v2.8.1\n" + feedback.OK.View(text.Normal, palette),
				"details": "Pool: public\nProbe: HTTP 200\nLatency: 12 ms",
				"cpu":     "43%\n4 cores",
				"memory":  "61%\n16 GiB",
			}
			titles := map[string]string{"hosts": "Hosts", "summary": "Selected host", "details": "Probe evidence", "cpu": "CPU", "memory": "Memory"}
			views := make(map[string]string, len(contents))

			for _, r := range screen.Plan().Regions() {
				allocation := layout.Size{Width: r.Bounds.Dx(), Height: r.Bounds.Dy()}
				views[r.ID] = layout.Panel(titles[r.ID], contents[r.ID], allocation, palette, r.ID == focused)
			}

			frame, renderErr := screen.Render(views)
			if renderErr != nil {
				return renderErr.Error()
			}

			return frame.Content
		},
	}, nil
}

package pattern

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type CorrelationScope string

const (
	RequestScope CorrelationScope = ""
	TraceScope   CorrelationScope = "trace"
)

const (
	requestPane          = "requests"
	requestDataPane      = "request"
	httpPane             = "http"
	logPane              = "logs"
	inspectorMinRequests = 34
	inspectorMinDetails  = 52
	inspectorMinHeight   = 20
)

type LogInspectorOptions struct {
	Layout                       InspectorLayout
	KeyLayout                    interaction.KeyLayout
	BackdropMode                 layout.BackdropMode
	TimelineGroups               []collection.HTTPGroup
	TimelineNow                  time.Time
	Symbols                      text.Presentation
	URLFormat                    text.URLFormat
	VerboseDetails, ColumnLabels bool
	Theme                        theme.Theme
	Requests                     []collection.HTTPExchange
	Logs                         []collection.LogRecord
	Limit                        int
	NoFolds                      bool
	NoHelp, NoHeadline           bool
	DetachedLogDetails           bool
}

// LogInspector owns coordinated UI state. Producers send observations through
// the application's Bubble Tea update loop; this model performs no capture I/O.
// Request and Response are independently scrollable native viewports. Apply
// payload styles/key maps there; HTTP.Details remains the standalone combined
// view and is not rendered by this inspector. Resizing and data updates own
// payload content, preserve same-selection offsets and reset both on selection changes.
type LogInspector struct {
	Symbols                         text.Presentation
	HTTP                            *collection.HTTPLog
	Palette                         *collection.CommandPalette
	Logs                            *collection.LogTree
	Request, Response               viewport.Model
	RequestJSON, ResponseJSON       *collection.JSONViewer
	focus                           interaction.Focus
	palette                         theme.Theme
	size, contentSize               layout.Size
	regions                         inspectorRegions
	preferred, active               InspectorLayout
	wide                            bool
	scope                           CorrelationScope
	logs                            []collection.LogRecord
	limit                           int
	correlationMessage              string
	collapsed                       map[string]bool
	folding, showHelp, showHeadline bool
	detailID                        string
	detailRegions                   inspectorDetails
	jsonStates                      map[string]inspectorJSONState
}

func NewLogInspector(o LogInspectorOptions) (*LogInspector, error) {
	limit := o.Limit
	if limit <= 0 {
		limit = collection.DefaultLogLimit
	}

	m := &LogInspector{palette: o.Theme, limit: limit, focus: interaction.NewFocus(requestPane, requestDataPane, httpPane, logPane), collapsed: map[string]bool{}, folding: !o.NoFolds, showHelp: !o.NoHelp, showHeadline: !o.NoHeadline}
	m.Request, m.Response = collection.NewViewport("", 40, 12), collection.NewViewport("", 60, 12)
	m.RequestJSON, m.ResponseJSON = collection.NewJSONViewer(40, 12, o.Theme), collection.NewJSONViewer(60, 12, o.Theme)
	m.RequestJSON.Model, m.ResponseJSON.Model = &m.Request, &m.Response
	m.jsonStates = map[string]inspectorJSONState{}
	m.HTTP = collection.NewHTTPLog(40, 12, o.Theme)
	m.Symbols = o.Symbols
	m.HTTP.Symbols = o.Symbols
	m.HTTP.URLFormat = o.URLFormat
	m.HTTP.VerboseDetails = o.VerboseDetails
	m.HTTP.ShowColumnLabels = o.ColumnLabels
	m.HTTP.Limit = limit
	m.HTTP.TimelineNow = o.TimelineNow
	m.HTTP.SetTimelineGrouping(o.TimelineGroups...)
	m.Logs = collection.NewLogTree(40, 12, o.Theme)

	m.Logs.Symbols = o.Symbols
	m.Logs.Limit = limit
	m.Logs.SetDetachedDetails(o.DetachedLogDetails)
	m.HTTP.SetFolding(!o.NoFolds)
	m.Logs.SetFolding(!o.NoFolds)

	if err := m.SetEntries(o.Requests, o.Logs); err != nil {
		return nil, err
	}

	if !m.SetLayout(o.Layout) {
		return nil, fmt.Errorf("unknown inspector layout %q", o.Layout)
	}

	palette, err := collection.NewCommandPalette(nil, 72, 24, o.Theme)
	if err != nil {
		return nil, err
	}

	m.Palette = palette
	m.Palette.SetHelpVisible(m.showHelp)

	if !m.Palette.SetBackdropMode(o.BackdropMode) {
		return nil, fmt.Errorf("unknown backdrop mode %q", o.BackdropMode)
	}

	if o.KeyLayout != interaction.DefaultKeys {
		if _, err := m.SetKeyLayout(o.KeyLayout); err != nil {
			return nil, err
		}
	}

	m.Resize(layout.Size{Width: 120, Height: 36})

	return m, nil
}

func (m *LogInspector) Init() tea.Cmd {
	if m.Palette != nil && m.Palette.IsOpen() && m.Palette.Mode() == collection.PaletteSearch {
		return m.Palette.Search.Focus()
	}

	return nil
}
func (m *LogInspector) FocusPane(id string) bool   { return m.focus.Set(inspectorPaneID(id)) }
func (m *LogInspector) FocusedPane() string        { return m.focus.Current() }
func (m *LogInspector) Scope() CorrelationScope    { return m.scope }
func (m *LogInspector) CorrelationMessage() string { return m.correlationMessage }
func (m *LogInspector) SetEntries(requests []collection.HTTPExchange, logs []collection.LogRecord) error {
	if err := collection.ValidateHTTPExchanges(requests); err != nil {
		return err
	}

	if err := collection.ValidateLogRecords(logs); err != nil {
		return err
	}

	if err := m.HTTP.SetEntries(requests); err != nil {
		return err
	}

	m.logs = slices.Clone(logs[max(0, len(logs)-m.limit):])
	m.correlate()
	m.pruneJSONStates()

	return nil
}

func (m *LogInspector) UpsertRequests(requests ...collection.HTTPExchange) error {
	if err := m.HTTP.Upsert(requests...); err != nil {
		return err
	}

	m.correlate()
	m.pruneJSONStates()

	return nil
}

func (m *LogInspector) UpsertLogs(logs ...collection.LogRecord) error {
	if err := collection.ValidateLogRecords(logs); err != nil {
		return err
	}

	positions := map[string]int{}

	next := slices.Clone(m.logs)
	for i, r := range next {
		positions[r.ID] = i
	}

	for _, r := range logs {
		if i, ok := positions[r.ID]; ok {
			next[i] = r
		} else {
			positions[r.ID] = len(next)
			next = append(next, r)
		}
	}

	m.logs = slices.Clone(next[max(0, len(next)-m.limit):])
	m.correlate()

	return nil
}

func (m *LogInspector) SelectRequest(id string) bool {
	if !m.HTTP.Select(id) {
		return false
	}

	m.HTTP.Follow = false
	m.correlate()

	return true
}

func (m *LogInspector) SetSymbols(mode text.Presentation) {
	m.Symbols = mode
	m.HTTP.SetSymbols(mode)
	m.Logs.SetSymbols(mode)
	m.Resize(m.size)
}

func (m *LogInspector) SetScope(scope CorrelationScope) bool {
	switch scope {
	case RequestScope:
	case TraceScope:
		record, ok := m.HTTP.Selected()
		if !ok || record.TraceID == "" {
			return false
		}
	default:
		return false
	}

	m.scope = scope
	m.correlate()

	return true
}

func (m *LogInspector) correlate() {
	record, ok := m.HTTP.Selected()
	matched := []collection.LogRecord{}
	m.correlationMessage = "Select an HTTP request"

	if ok {
		id := record.RequestID

		scope := "request"
		if m.scope == TraceScope {
			id = record.TraceID
			scope = "trace"
		}

		if id == "" {
			m.correlationMessage = "Correlation unavailable: no " + scope + " ID"
		} else {
			for _, entry := range m.logs {
				match := entry.RequestID == id
				if m.scope == TraceScope {
					match = entry.TraceID == id
				}

				if match {
					matched = append(matched, entry)
				}
			}

			m.correlationMessage = fmt.Sprintf("%s %s · %d matching logs", scope, text.Inline(id), len(matched))
			if len(matched) == 0 {
				m.correlationMessage = "No matching logs for " + scope + " " + text.Inline(id)
			}
		}
	}
	// WHY: The source dataset was validated before committing; a subset cannot add duplicate IDs.
	if err := m.Logs.SetEntries(matched); err != nil {
		m.correlationMessage = "Invalid correlated records: " + text.Inline(err.Error())
	}

	if m.Palette != nil && m.Palette.IsOpen() {
		if err := m.Palette.SetCommands(m.Commands()); err != nil {
			m.correlationMessage = "Invalid commands: " + text.Inline(err.Error())
		}
	}

	m.Resize(m.size)
}

func (m *LogInspector) Resize(size layout.Size) {
	m.size = size
	if m.Palette != nil {
		m.Palette.Resize(size)
	}

	bodySize := size
	if m.showHelp {
		bodySize.Height -= 2
	}

	if headline := m.headlineView(); headline != "" {
		bodySize.Height -= lipgloss.Height(headline)
	}

	m.contentSize = bodySize
	m.active, m.regions = planInspector(bodySize, m.preferred, m.HTTP.TimelineHeight())
	m.regions = foldInspector(bodySize, m.active, m.regions, map[string]bool{requestPane: m.PaneCollapsed(requestPane), httpPane: m.PaneCollapsed(requestDataPane) && m.PaneCollapsed(httpPane), logPane: m.PaneCollapsed(logPane)})
	m.wide = m.active != InspectorFocused
	m.HTTP.Timeline = m.active == InspectorTimeline || (!m.wide && m.preferred == InspectorTimeline)
	requests := layout.PanelContentSize(m.regions.requests)
	details := layout.PanelContentSize(m.regions.details)
	logs := layout.PanelContentSize(m.regions.logs)
	m.HTTP.ResizeList(requests.Width, requests.Height)

	if !m.PaneCollapsed(httpPane) {
		m.HTTP.ResizeDetails(details.Width, details.Height)
	}

	if !m.PaneCollapsed(logPane) {
		m.Logs.Resize(logs.Width, max(1, logs.Height-1))
	}

	m.resizeDetailPanes()
}

func (m *LogInspector) updateContent(msg tea.Msg) tea.Cmd {
	if key, ok := msg.(tea.KeyPressMsg); ok {
		if m.focus.Handle(msg) {
			return nil
		}

		switch key.String() {
		case "1", "2":
			if m.focus.Current() == requestDataPane || m.focus.Current() == httpPane {
				return m.updateHTTPControls(msg)
			}
		case "?":
			m.SetHelpVisible(!m.HelpVisible())
			return nil
		case "z":
			m.SetPaneCollapsed(m.focus.Current(), !m.PaneCollapsed(m.focus.Current()))
			return nil
		case "L":
			m.cycleLayout()
			return nil
		case "alt+1":
			m.focus.Set(requestPane)
			return nil
		case "alt+2":
			m.focus.Set(httpPane)
			return nil
		case "alt+3":
			m.focus.Set(logPane)
			return nil
		case "alt+4":
			m.focus.Set(requestDataPane)
			return nil
		case "t":
			if m.scope == TraceScope {
				m.SetScope(RequestScope)
			} else {
				m.SetScope(TraceScope)
			}

			return nil
		case "F":
			switch m.focus.Current() {
			case requestPane:
				m.HTTP.ResumeFollow()
				m.correlate()
			case logPane:
				m.Logs.ResumeFollow()
			}

			return nil
		}
	}

	if m.PaneCollapsed(m.focus.Current()) {
		return nil
	}

	switch m.focus.Current() {
	case requestPane:
		before := m.HTTP.SelectedID()

		cmd := m.HTTP.UpdateRequests(msg)
		if before != m.HTTP.SelectedID() {
			m.correlate()
		}

		return cmd
	case requestDataPane, httpPane:
		return m.updateDetailPane(msg)
	case logPane:
		return m.Logs.Update(msg)
	}

	return nil
}

func (m *LogInspector) contentView() string {
	focused := m.focus.Current()

	requests := func(size layout.Size) string {
		title := "HTTP requests"
		if m.HTTP.Timeline {
			title = "Request timeline"
		}

		return m.section(requestPane, title+" · "+feedback.Follow(m.HTTP.Follow, text.Normal, m.palette, m.Symbols), m.HTTP.ListView(), size)
	}
	details := func(size layout.Size) string {
		return m.detailPanesView(size)
	}
	logs := func(size layout.Size) string {
		return m.section(logPane, "Related logs · "+feedback.Follow(m.Logs.Follow, text.Normal, m.palette, m.Symbols), m.palette.Subtle().Render(m.correlationMessage)+"\n"+m.Logs.View(), size)
	}

	var content string

	switch m.active {
	case InspectorAuto, InspectorSidebar:
		content = layout.Row(requests(m.regions.requests), " ", layout.Column(details(m.regions.details), logs(m.regions.logs)))
		if m.regions.requestsAbove {
			content = layout.Column(requests(m.regions.requests), details(m.regions.details), logs(m.regions.logs))
		}
	case InspectorTimeline:
		content = layout.Column(requests(m.regions.requests), layout.Row(details(m.regions.details), " ", logs(m.regions.logs)))
		if m.regions.bottomStacked {
			content = layout.Column(requests(m.regions.requests), details(m.regions.details), logs(m.regions.logs))
		}
	case InspectorColumns:
		var headings, columns []string

		for _, pane := range []struct{ id, view string }{{requestPane, requests(m.regions.requests)}, {httpPane, details(m.regions.details)}, {logPane, logs(m.regions.logs)}} {
			if m.PaneCollapsed(pane.id) && (pane.id != httpPane || m.PaneCollapsed(requestDataPane)) {
				headings = append(headings, pane.view)
			} else {
				columns = append(columns, pane.view)
			}
		}

		var parts []string

		for i, column := range columns {
			if i > 0 {
				parts = append(parts, " ")
			}

			parts = append(parts, column)
		}

		if len(parts) > 0 {
			headings = append(headings, layout.Row(parts...))
		}

		content = layout.Column(headings...)
	case InspectorFocused:
		switch focused {
		case requestPane:
			content = requests(m.regions.requests)
		case requestDataPane, httpPane:
			content = details(m.regions.details)
		case logPane:
			content = logs(m.regions.logs)
		}
	default:
		content = ""
	}

	content = layout.Fit(content, m.contentSize)
	if headline := m.headlineView(); headline != "" {
		content = headline + "\n" + content
	}

	if !m.showHelp {
		return layout.Fit(content, m.size)
	}

	help := text.KeyHint("Tab", "panes", m.palette) + " · " + text.KeyHint("L", "layout", m.palette) + " · " + text.KeyHint("Alt+1/2/3/4", "views", m.palette) + "\n"

	if m.size.Width >= 70 {
		help = strings.TrimSuffix(help, "\n") + " · " + text.KeyHint(m.Palette.Keys.Open.Help().Key, "commands", m.palette) + " · " + text.KeyHint("?", "help", m.palette) + "\n"
	} else {
		help = text.KeyHint("Tab", "panes", m.palette) + " · " + text.KeyHint("?", "help", m.palette) + " · " + text.KeyHint(m.Palette.Keys.Open.Help().Key, "commands", m.palette) + "\n"
	}

	switch focused {
	case requestPane:
		help += text.KeyHint(m.HTTP.Table.KeyMap.LineUp.Help().Key+"/"+m.HTTP.Table.KeyMap.LineDown.Help().Key, "select", m.palette) + " · " + text.KeyHint("F", "follow", m.palette)
	case requestDataPane, httpPane:
		keys := m.Response.KeyMap
		if focused == requestDataPane {
			keys = m.Request.KeyMap
		}

		viewer := m.focusedJSON()
		if viewer != nil && viewer.Interactive() {
			help += text.KeyHint(keys.Up.Help().Key+"/"+keys.Down.Help().Key, "fields", m.palette) + " · " + text.KeyHint(viewer.Toggle.Help().Key, "fold JSON", m.palette)
		} else {
			help += text.KeyHint(keys.Up.Help().Key+"/"+keys.Down.Help().Key, "scroll", m.palette)
		}

		if m.HTTP.Folding && (viewer == nil || !viewer.Interactive()) {
			help += " · " + text.KeyHint("1/2/3", "sections", m.palette)
		}

		if m.HTTP.CanToggleRaw() {
			help += " · " + text.KeyHint("r", "raw", m.palette)
		}
	case logPane:
		parts := []string{text.KeyHint("g", "group", m.palette), text.KeyHint("F", "follow", m.palette)}
		if m.Logs.CanFocusDetails() {
			target := "details"
			if m.Logs.DetailsFocused {
				target = "tree"
			}

			parts = append([]string{text.KeyHint("D", target, m.palette)}, parts...)
		}

		if m.Logs.Folding && !m.Logs.CanFocusDetails() {
			parts = append([]string{text.KeyHint("Enter", "expand", m.palette)}, parts...)
		}

		help += strings.Join(parts, " · ")
	}

	if selected, ok := m.HTTP.Selected(); ok && selected.TraceID != "" && focused == requestPane {
		help += " · " + text.KeyHint("t", "scope", m.palette)
	}

	if m.folding {
		label := "fold pane"
		if m.PaneCollapsed(focused) {
			label = "expand pane"
			help = strings.SplitN(help, "\n", 2)[0] + "\nPane folded"
		}

		help += " · " + text.KeyHint("z", label, m.palette)
	}

	return layout.Fit(content+"\n"+m.palette.Subtle().Render(help), m.size)
}

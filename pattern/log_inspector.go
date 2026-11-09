package pattern

import (
	"fmt"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
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
	httpPane             = "http"
	logPane              = "logs"
	inspectorMinRequests = 34
	inspectorMinDetails  = 52
	inspectorMinHeight   = 20
)

type LogInspectorOptions struct {
	Symbols                      text.Presentation
	URLFormat                    text.URLFormat
	VerboseDetails, ColumnLabels bool
	Theme                        theme.Theme
	Requests                     []collection.HTTPExchange
	Logs                         []collection.LogRecord
	Limit                        int
	NoFolds                      bool
	DetachedLogDetails           bool
}

// LogInspector owns coordinated UI state. Producers send observations through
// the application's Bubble Tea update loop; this model performs no capture I/O.
type LogInspector struct {
	Symbols                              text.Presentation
	HTTP                                 *collection.HTTPLog
	Logs                                 *collection.LogTree
	focus                                interaction.Focus
	palette                              theme.Theme
	size                                 layout.Size
	left, right, detailHeight, logHeight int
	wide                                 bool
	scope                                CorrelationScope
	logs                                 []collection.LogRecord
	limit                                int
	correlationMessage                   string
}

func NewLogInspector(o LogInspectorOptions) (*LogInspector, error) {
	limit := o.Limit
	if limit <= 0 {
		limit = collection.DefaultLogLimit
	}

	m := &LogInspector{palette: o.Theme, limit: limit, focus: interaction.NewFocus(requestPane, httpPane, logPane)}
	m.HTTP = collection.NewHTTPLog(40, 12, o.Theme)
	m.Symbols = o.Symbols
	m.HTTP.Symbols = o.Symbols
	m.HTTP.URLFormat = o.URLFormat
	m.HTTP.VerboseDetails = o.VerboseDetails
	m.HTTP.ShowColumnLabels = o.ColumnLabels
	m.HTTP.Limit = limit
	m.Logs = collection.NewLogTree(40, 12, o.Theme)

	m.Logs.Symbols = o.Symbols
	m.Logs.Limit = limit
	m.Logs.SetDetachedDetails(o.DetachedLogDetails)
	m.HTTP.SetFolding(!o.NoFolds)
	m.Logs.SetFolding(!o.NoFolds)

	if err := m.SetEntries(o.Requests, o.Logs); err != nil {
		return nil, err
	}

	m.Resize(layout.Size{Width: 120, Height: 36})

	return m, nil
}
func (m *LogInspector) Init() tea.Cmd              { return nil }
func (m *LogInspector) FocusPane(id string) bool   { return m.focus.Set(id) }
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

	return nil
}

func (m *LogInspector) UpsertRequests(requests ...collection.HTTPExchange) error {
	if err := m.HTTP.Upsert(requests...); err != nil {
		return err
	}

	m.correlate()

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
}

func (m *LogInspector) Resize(size layout.Size) {
	m.size = size
	m.left, m.right, m.wide = layout.Split(size.Width, theme.SpaceSmall, inspectorMinRequests, inspectorMinDetails, .34)
	m.wide = m.wide && size.Height >= inspectorMinHeight

	contentHeight := max(1, size.Height-2)
	if m.wide {
		m.detailHeight = max(7, contentHeight*3/5)
		m.logHeight = max(3, contentHeight-m.detailHeight)
		requests := layout.PanelContentSize(layout.Size{Width: m.left, Height: contentHeight})
		details := layout.PanelContentSize(layout.Size{Width: m.right, Height: m.detailHeight})
		logs := layout.PanelContentSize(layout.Size{Width: m.right, Height: m.logHeight})
		m.HTTP.Resize(requests.Width, requests.Height)
		m.HTTP.ResizeList(requests.Width, requests.Height)
		m.HTTP.ResizeDetails(details.Width, details.Height)
		m.Logs.Resize(logs.Width, max(1, logs.Height-1))
	} else {
		content := layout.PanelContentSize(layout.Size{Width: size.Width, Height: contentHeight})
		m.HTTP.Resize(content.Width, content.Height)
		m.HTTP.ResizeList(content.Width, content.Height)
		m.HTTP.ResizeDetails(content.Width, content.Height)
		m.Logs.Resize(content.Width, max(1, content.Height-1))
	}
}

func (m *LogInspector) Update(msg tea.Msg) tea.Cmd {
	if key, ok := msg.(tea.KeyPressMsg); ok {
		if m.focus.Handle(msg) {
			return nil
		}

		switch key.String() {
		case "alt+1":
			m.focus.Set(requestPane)
			return nil
		case "alt+2":
			m.focus.Set(httpPane)
			return nil
		case "alt+3":
			m.focus.Set(logPane)
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

	switch m.focus.Current() {
	case requestPane:
		before := m.HTTP.SelectedID()

		cmd := m.HTTP.UpdateRequests(msg)
		if before != m.HTTP.SelectedID() {
			m.correlate()
		}

		return cmd
	case httpPane:
		return m.HTTP.UpdateDetails(msg)
	case logPane:
		return m.Logs.Update(msg)
	}

	return nil
}

func (m *LogInspector) View() string {
	focused := m.focus.Current()

	requests := func(size layout.Size) string {
		return layout.Panel("HTTP requests · "+feedback.Follow(m.HTTP.Follow, text.Normal, m.palette, m.Symbols), m.HTTP.ListView(), size, m.palette, focused == requestPane)
	}
	details := func(size layout.Size) string {
		return layout.Panel("HTTP details", m.HTTP.DetailView(), size, m.palette, focused == httpPane)
	}
	logs := func(size layout.Size) string {
		return layout.Panel("Related logs · "+feedback.Follow(m.Logs.Follow, text.Normal, m.palette, m.Symbols), m.palette.Subtle().Render(m.correlationMessage)+"\n"+m.Logs.View(), size, m.palette, focused == logPane)
	}
	height := max(1, m.size.Height-2)

	var content string
	if m.wide {
		content = layout.Row(requests(layout.Size{Width: m.left, Height: height}), " ", layout.Column(details(layout.Size{Width: m.right, Height: m.detailHeight}), logs(layout.Size{Width: m.right, Height: m.logHeight})))
	} else {
		size := layout.Size{Width: m.size.Width, Height: height}

		switch focused {
		case requestPane:
			content = requests(size)
		case httpPane:
			content = details(size)
		case logPane:
			content = logs(size)
		}
	}

	help := text.KeyHint("Tab", "panes", m.palette) + " · " + text.KeyHint("Alt+1/2/3", "views", m.palette) + "\n"

	switch focused {
	case requestPane:
		help += text.KeyHint(m.palette.Glyph(theme.GlyphUp)+"/"+m.palette.Glyph(theme.GlyphDown), "select", m.palette) + " · " + text.KeyHint("F", "follow", m.palette)
	case httpPane:
		help += text.KeyHint(m.palette.Glyph(theme.GlyphUp)+"/"+m.palette.Glyph(theme.GlyphDown), "scroll", m.palette)
		if m.HTTP.Folding {
			help += " · " + text.KeyHint("1/2", "sections", m.palette)
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

	return layout.Fit(content+"\n"+m.palette.Subtle().Render(help), m.size)
}

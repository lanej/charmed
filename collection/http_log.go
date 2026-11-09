package collection

import (
	"net/http"
	"slices"
	"strings"
	"time"

	binding "charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// HTTPExchange is an observation, not an active request. No reader or network
// handle is retained. DurationKnown distinguishes unavailable latency from zero.
type HTTPExchange struct {
	ID, RequestID, TraceID          string
	Service                         string
	Time                            time.Time
	Method, URL                     string
	StatusCode                      int
	Completed                       bool
	Duration                        time.Duration
	DurationKnown                   bool
	RequestHeaders, ResponseHeaders http.Header
	RequestBody, ResponseBody       HTTPBody
}

// HTTPLog exposes native table and viewport models. Resize keeps custom styles
// and key maps; data changes rebuild rows and the selected detail content.
type HTTPLog struct {
	TimelineGroupBy                  []HTTPGroup
	Timeline                         bool
	TimelineLegend                   bool
	TimelineNow                      time.Time
	listHeight                       int
	Symbols                          text.Presentation
	URLFormat                        text.URLFormat
	URLFormatter                     func(string) string
	VerboseDetails, ShowColumnLabels bool
	Table                            table.Model
	Details                          viewport.Model
	Theme                            theme.Theme
	Limit, BodyPreviewBytes          int
	Follow, Raw                      bool
	RequestOpen, ResponseOpen        bool
	ParametersOpen                   bool
	records                          []HTTPExchange
	selectedID                       string
	listWidth, height                int
	detailsWidth                     int
	DetailsFocused                   bool
	Folding, ShowScrollbar           bool
}

func NewHTTPLog(width, height int, t theme.Theme) *HTTPLog {
	m := &HTTPLog{Theme: t, Limit: DefaultLogLimit, BodyPreviewBytes: DefaultBodyPreviewBytes, Follow: true, URLFormat: text.PathAndQuery, Folding: true, ShowScrollbar: true, RequestOpen: true, ResponseOpen: true, ParametersOpen: true}
	m.Table = NewDenseTable(nil, nil, max(1, width), max(1, height), t)
	m.Details = NewViewport("", width, height)
	m.Resize(width, height)

	return m
}
func (m *HTTPLog) Init() tea.Cmd      { return nil }
func (m *HTTPLog) SelectedID() string { return m.selectedID }
func (m *HTTPLog) Selected() (HTTPExchange, bool) {
	for _, record := range m.records {
		if record.ID == m.selectedID {
			return cloneExchange(record), true
		}
	}

	return HTTPExchange{}, false
}

func (m *HTTPLog) Entries() []HTTPExchange {
	out := make([]HTTPExchange, len(m.records))
	for i, record := range m.records {
		out[i] = cloneExchange(record)
	}

	return out
}

func cloneExchange(record HTTPExchange) HTTPExchange {
	record.RequestHeaders = record.RequestHeaders.Clone()
	record.ResponseHeaders = record.ResponseHeaders.Clone()
	record.RequestBody.Bytes = slices.Clone(record.RequestBody.Bytes)
	record.ResponseBody.Bytes = slices.Clone(record.ResponseBody.Bytes)

	return record
}

func (m *HTTPLog) SetEntries(records []HTTPExchange) error {
	if err := validateRecordIDs(records, func(r HTTPExchange) string { return r.ID }); err != nil {
		return err
	}

	for _, record := range records {
		if err := record.RequestBody.Validate(); err != nil {
			return err
		}

		if err := record.ResponseBody.Validate(); err != nil {
			return err
		}
	}

	previous := m.selectedID

	m.records = retainRecords(records, m.Limit)
	for i, record := range m.records {
		m.records[i] = cloneExchange(record)
	}

	m.rebuildRows()

	if m.Follow && len(m.records) > 0 {
		previous = m.records[len(m.records)-1].ID
	}

	if !m.Select(previous) {
		m.selectedID = ""
		if len(m.records) > 0 {
			m.Select(m.records[0].ID)
		} else {
			m.refreshDetails(true)
		}
	}

	return nil
}

func (m *HTTPLog) Upsert(records ...HTTPExchange) error {
	if err := validateRecordIDs(records, func(r HTTPExchange) string { return r.ID }); err != nil {
		return err
	}

	return m.SetEntries(upsertRecords(m.records, records, func(r HTTPExchange) string { return r.ID }))
}

func (m *HTTPLog) Select(id string) bool {
	for i, record := range m.records {
		if record.ID == id {
			changed := id != m.selectedID

			m.selectedID = id
			if changed {
				m.rebuildRows()
			}

			m.Table.SetCursor(i)
			m.refreshDetails(changed)

			return true
		}
	}

	return false
}

func (m *HTTPLog) ResumeFollow() {
	m.Follow = true
	if len(m.records) > 0 {
		m.Select(m.records[len(m.records)-1].ID)
	}
}

func (m *HTTPLog) Resize(width, height int) {
	m.height = max(1, height)
	m.listWidth = max(1, width)
	m.detailsWidth = max(1, width)

	listHeight := max(1, height/3)
	if m.Timeline {
		listHeight = max(1, height/2)
	}

	m.ResizeList(width, listHeight)
	m.Details.SetWidth(max(1, width-m.scrollbarWidth()))
	m.Details.SetHeight(max(1, height-listHeight-2))
	m.rebuildRows()
	m.refreshDetails(false)
}

func (m *HTTPLog) ResizeList(width, height int) {
	m.listWidth = max(1, width)
	m.listHeight = max(1, height)
	m.Table.SetWidth(m.listWidth)
	m.rebuildRows()

	headerHeight := m.headerHeight()
	if !m.ShowColumnLabels {
		height += headerHeight
	}

	m.Table.SetHeight(max(headerHeight, height))
}

// WHY: Native header styling can add multiple rows, even with blank titles.
func (m *HTTPLog) headerHeight() int {
	probe := m.Table
	probe.SetRows(nil)

	return max(1, lipgloss.Height(probe.View())-max(1, probe.Height()))
}

func (m *HTTPLog) SetURLFormat(format text.URLFormat) {
	m.URLFormat = format
	m.rebuildRows()
	m.refreshDetails(false)
}

func (m *HTTPLog) SetURLFormatter(formatter func(string) string) {
	m.URLFormatter = formatter
	m.rebuildRows()
	m.refreshDetails(false)
}

func (m *HTTPLog) SetSymbols(mode text.Presentation) { m.Symbols = mode; m.refreshDetails(false) }

func (m *HTTPLog) target(value string) string {
	if m.URLFormatter != nil {
		return text.Inline(m.URLFormatter(value))
	}

	return text.FormatURL(value, m.URLFormat)
}

func (m *HTTPLog) ResizeDetails(width, height int) {
	m.detailsWidth = max(1, width)
	m.Details.SetWidth(max(1, width-m.scrollbarWidth()))
	m.Details.SetHeight(max(1, height))
	m.refreshDetails(false)
}

func (m *HTTPLog) View() string {
	return layout.Fit(m.ListView()+"\n"+m.DetailView(), layout.Size{Width: m.listWidth, Height: m.height})
}

func (m *HTTPLog) Update(msg tea.Msg) tea.Cmd {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "tab" {
		m.DetailsFocused = !m.DetailsFocused
		return nil
	}

	if m.DetailsFocused {
		return m.UpdateDetails(msg)
	}

	return m.UpdateRequests(msg)
}

func (m *HTTPLog) ListView() string {
	if m.Timeline {
		return layout.Fit(m.TimelineView(m.TimelineNow, m.listWidth, m.listHeight), layout.Size{Width: m.listWidth, Height: m.listHeight})
	}

	if len(m.records) == 0 {
		return m.Theme.Subtle().Render("No HTTP requests")
	}

	view := m.Table.View()
	if !m.ShowColumnLabels {
		lines := strings.Split(view, "\n")
		view = strings.Join(lines[min(m.headerHeight(), len(lines)):], "\n")
	}

	return layout.Fit(view, layout.Size{Width: m.listWidth, Height: m.listHeight})
}

func (m *HTTPLog) DetailView() string {
	if m.scrollbarWidth() == 0 {
		return m.Details.View()
	}

	bar := feedback.Scrollbar{Total: m.Details.TotalLineCount(), Visible: m.Details.Height(), Offset: m.Details.YOffset(), Height: m.Details.Height()}

	return lipgloss.JoinHorizontal(lipgloss.Top, m.Details.View(), bar.View(m.Theme))
}

func (m *HTTPLog) scrollbarWidth() int {
	if m.ShowScrollbar && m.detailsWidth > 1 {
		return 1
	}

	return 0
}

func (m *HTTPLog) SetFolding(enabled bool) {
	m.Folding = enabled
	if !enabled {
		m.RequestOpen, m.ResponseOpen, m.ParametersOpen = true, true, true
	}

	m.refreshDetails(false)
}

func (m *HTTPLog) UpdateRequests(msg tea.Msg) tea.Cmd {
	if pressed, ok := msg.(tea.KeyPressMsg); ok {
		keys := m.Table.KeyMap
		if binding.Matches(pressed, keys.GotoBottom) {
			m.ResumeFollow()
			return nil
		}

		if binding.Matches(pressed, keys.LineUp, keys.LineDown, keys.PageUp, keys.PageDown, keys.HalfPageUp, keys.HalfPageDown, keys.GotoTop) {
			m.Follow = false
		}
	}

	var cmd tea.Cmd

	records := m.records
	if m.Timeline && len(m.TimelineGroupBy) > 0 {
		records = m.timelineRecords()
		for i, record := range records {
			if record.ID == m.selectedID {
				m.Table.SetCursor(i)
				break
			}
		}
	}

	m.Table, cmd = m.Table.Update(msg)
	if len(records) > 0 {
		index := min(max(0, m.Table.Cursor()), len(m.records)-1)
		m.Select(records[index].ID)
	}

	return cmd
}

func (m *HTTPLog) UpdateDetails(msg tea.Msg) tea.Cmd {
	if key, ok := msg.(tea.KeyPressMsg); ok {
		switch key.String() {
		case "r":
			if !m.CanToggleRaw() {
				return nil
			}

			m.Raw = !m.Raw
			m.refreshDetails(false)

			return nil
		case "1":
			if !m.Folding {
				return nil
			}

			m.RequestOpen = !m.RequestOpen
			m.refreshDetails(false)

			return nil
		case "3":
			if !m.Folding {
				return nil
			}

			m.ParametersOpen = !m.ParametersOpen
			m.refreshDetails(false)

			return nil

		case "2":
			if !m.Folding {
				return nil
			}

			m.ResponseOpen = !m.ResponseOpen
			m.refreshDetails(false)

			return nil
		}
	}

	var cmd tea.Cmd

	m.Details, cmd = m.Details.Update(msg)

	return cmd
}

func (m *HTTPLog) rebuildRows() {
	const methodWidth, statusWidth, durationWidth = 9, 5, 10

	columns := []table.Column{{Title: "", Width: 2}, {Title: " Method", Width: methodWidth}, {Title: "Code", Width: statusWidth}}

	showDuration := m.listWidth >= 56
	if showDuration {
		columns = append(columns, table.Column{Title: "   Latency", Width: durationWidth})
	}

	used := 2 + methodWidth + statusWidth
	if showDuration {
		used += durationWidth
	}

	columns = append(columns, table.Column{Title: "Target", Width: max(1, m.listWidth-used-DenseTablePaddingWidth(len(columns)+1))})

	compact := m.listWidth < 30
	if compact {
		columns = []table.Column{{Title: "Code", Width: statusWidth}, {Title: "Target", Width: max(1, m.listWidth-statusWidth-DenseTablePaddingWidth(2))}}
	}

	if !m.ShowColumnLabels {
		for i := range columns {
			columns[i].Title = ""
		}
	}
	// WHY: Native SetColumns renders existing rows immediately; clear the old schema first.
	m.Table.SetRows(nil)
	m.Table.SetColumns(columns)

	rows := make([]table.Row, 0, len(m.records))
	for _, record := range m.records {
		status := text.HTTPStatusCode(record.StatusCode, m.Theme)
		if !record.Completed && record.StatusCode == 0 {
			status = text.Pending(m.Theme)
		}

		cursor := ""
		if record.ID == m.selectedID {
			cursor = m.Theme.Heading().Render(m.Theme.Glyph(theme.GlyphRight))
		}

		row := table.Row{cursor, text.HTTPMethod(record.Method, m.Theme), status}

		if showDuration {
			duration := "n/a"
			if record.DurationKnown {
				duration = text.ShortDuration(record.Duration)
			}

			row = append(row, lipgloss.NewStyle().Width(durationWidth).Align(lipgloss.Right).Render(ansi.Truncate(duration, durationWidth, "…")))
		}

		if compact {
			row = table.Row{status}
		}

		rows = append(rows, append(row, m.target(record.URL)))
	}

	m.Table.SetRows(rows)

	for i, record := range m.records {
		if record.ID == m.selectedID {
			m.Table.SetCursor(i)
			break
		}
	}
}

func (m *HTTPLog) refreshDetails(reset bool) {
	offset := m.Details.YOffset()

	record, ok := m.Selected()
	if !ok {
		m.Details.SetContent("Select an HTTP request")
		return
	}

	width := max(1, m.Details.Width())
	rows := []string{}

	status := text.HTTPStatusCode(record.StatusCode, m.Theme)
	if !record.Completed && record.StatusCode == 0 {
		status = text.PendingWithPresentation(m.Symbols, m.Theme)
	}

	latency := "elapsed unavailable"
	if record.DurationKnown {
		latency = record.Duration.String()
	}

	stamp := "unavailable"
	if !record.Time.IsZero() {
		stamp = record.Time.Format(time.RFC3339Nano)
	}

	if m.VerboseDetails {
		rows = append(rows, detailLine("Method", unpadHTTPAtom(text.HTTPMethod(record.Method, m.Theme)), width), detailLine("Target", m.target(record.URL), width), detailLine("Time", m.Theme.Subtle().Render(stamp), width), detailLine("Latency", latency, width), detailLine("Status", unpadHTTPAtom(status)+" "+statusDescription(record), width))
	} else {
		rows = append(rows, hangingWrap(unpadHTTPAtom(text.HTTPMethod(record.Method, m.Theme))+" ", m.target(record.URL), width))

		rows = append(rows, unpadHTTPAtom(status)+" "+statusDescription(record)+" · "+latency)
		if !record.Time.IsZero() {
			rows = append(rows, m.Theme.Subtle().Render(stamp))
		}
	}

	rows = append(rows, "")
	if m.VerboseDetails || record.RequestID != "" {
		rows = append(rows, detailLine("Request ID", emptyID(record.RequestID), width))
	}

	if m.VerboseDetails || record.TraceID != "" {
		rows = append(rows, detailLine("Trace ID", emptyID(record.TraceID), width))
	}

	rows = append(rows, "")

	if parameters := m.parameterView(record, width); parameters != "" {
		rows = append(rows, parameters, "")
	}

	for _, part := range []struct {
		name    string
		open    bool
		headers http.Header
		body    HTTPBody
	}{{"Request", m.RequestOpen, record.RequestHeaders, record.RequestBody}, {"Response", m.ResponseOpen, record.ResponseHeaders, record.ResponseBody}} {
		heading := m.Theme.Heading().Render(part.name)
		if m.Folding {
			marker := text.Disclosure("", part.open, m.Theme, m.Symbols)
			if marker != "" {
				heading = marker + " " + heading
			}
		}

		rows = append(rows, heading)
		if part.open {
			rows = append(rows, headerView(part.headers, m.Theme, width), bodyWithContentType(part.body, part.headers).View(m.Raw, m.BodyPreviewBytes, m.Theme), "")
		}
	}

	content := ansi.Wrap(strings.Join(rows, "\n"), max(1, m.Details.Width()), "")
	m.Details.SetContent(content)

	if reset {
		m.Details.GotoTop()
	} else {
		m.Details.SetYOffset(offset)
	}
}

func emptyID(value string) string {
	if value == "" {
		return "unavailable"
	}

	return text.Inline(value)
}

func headerView(headers http.Header, t theme.Theme, width int) string {
	keys := make([]string, 0, len(headers))
	for key := range headers {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	rows := []string{}

	for _, key := range keys {
		for _, value := range headers[key] {
			rows = append(rows, hangingWrap(logStyle(t, t.Syntax.Name).Render(text.Inline(key))+": ", text.Inline(value), width))
		}
	}

	if len(rows) == 0 {
		return t.Subtle().Render("No captured headers")
	}

	return strings.Join(rows, "\n")
}

func ValidateHTTPExchanges(records []HTTPExchange) error {
	if err := validateRecordIDs(records, func(r HTTPExchange) string { return r.ID }); err != nil {
		return err
	}

	for _, r := range records {
		if err := r.RequestBody.Validate(); err != nil {
			return err
		}

		if err := r.ResponseBody.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (m *HTTPLog) CanToggleRaw() bool {
	r, ok := m.Selected()
	if !ok {
		return false
	}

	return ((r.RequestBody.State == Captured || r.RequestBody.State == Truncated) && len(r.RequestBody.Bytes) > 0) || ((r.ResponseBody.State == Captured || r.ResponseBody.State == Truncated) && len(r.ResponseBody.Bytes) > 0)
}

func bodyWithContentType(body HTTPBody, headers http.Header) HTTPBody {
	if body.ContentType == "" {
		body.ContentType = headers.Get("Content-Type")
	}

	return body
}

func detailLine(label, value string, width int) string {
	const labelWidth = 11

	prefix := lipgloss.NewStyle().Width(labelWidth).Render(label)

	return hangingWrap(prefix, value, width)
}

func hangingWrap(prefix, value string, width int) string {
	width = max(1, width)

	indent := ansi.StringWidth(prefix)
	if indent >= width-1 {
		return ansi.Wrap(prefix, width, "") + "\n" + ansi.Wrap(value, width, "")
	}

	wrapped := ansi.Wrap(value, max(1, width-indent), "")

	return prefix + strings.ReplaceAll(wrapped, "\n", "\n"+strings.Repeat(" ", indent))
}

func unpadHTTPAtom(value string) string {
	return ansi.Cut(value, 1, max(1, ansi.StringWidth(value)-1))
}

func statusDescription(record HTTPExchange) string {
	if !record.Completed && record.StatusCode == 0 {
		return ""
	}

	return text.Inline(http.StatusText(record.StatusCode))
}

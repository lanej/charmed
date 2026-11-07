// Package interaction contains focus and keyboard conventions for composed widgets.
package interaction

import tea "charm.land/bubbletea/v2"

type Focus struct {
	ids   []string
	index int
}

func NewFocus(ids ...string) Focus { return Focus{ids: append([]string(nil), ids...)} }
func (f Focus) Current() string {
	if len(f.ids) == 0 {
		return ""
	}
	return f.ids[f.index]
}
func (f *Focus) Set(id string) bool {
	for i, candidate := range f.ids {
		if candidate == id {
			f.index = i
			return true
		}
	}
	return false
}
func (f *Focus) Next() {
	if len(f.ids) > 0 {
		f.index = (f.index + 1) % len(f.ids)
	}
}
func (f *Focus) Previous() {
	if len(f.ids) > 0 {
		f.index = (f.index + len(f.ids) - 1) % len(f.ids)
	}
}
func (f *Focus) Handle(msg tea.Msg) bool {
	k, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return false
	}
	switch k.String() {
	case "tab":
		f.Next()
		return true
	case "shift+tab":
		f.Previous()
		return true
	}
	return false
}

type Selected struct {
	Component string `json:"component"`
	ID        string `json:"id"`
}
type Submitted struct {
	Component string `json:"component"`
	Value     string `json:"value"`
}
type RangeChanged struct {
	Component  string `json:"component"`
	Start, End float64
}

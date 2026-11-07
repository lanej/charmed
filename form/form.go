// Package form integrates Huh forms as one nested focus scope.
package form

import (
	"charm.land/huh/v2"
	"github.com/lanej/charmed/theme"
)

func Theme(t theme.Theme) huh.Theme {
	return huh.ThemeFunc(func(dark bool) *huh.Styles {
		s := huh.ThemeBase(dark)
		s.Focused.Title = t.Heading()
		s.Blurred.Title = t.Subtle()
		s.Focused.Description = t.Subtle()
		s.Blurred.Description = t.Subtle()
		s.Focused.SelectSelector = t.Heading()
		s.Focused.SelectedOption = t.Heading()
		s.Focused.TextInput.Text = t.Text()
		s.Blurred.TextInput.Text = t.Text()
		s.Focused.ErrorMessage = t.Text().Foreground(t.Danger)
		return s
	})
}

func New(t theme.Theme, groups ...*huh.Group) *huh.Form {
	return huh.NewForm(groups...).WithTheme(Theme(t)).WithShowHelp(true)
}

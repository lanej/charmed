package fixture

import "github.com/lanej/charmed/text"

// StatusTones is HTTP presentation policy supplied by the example application.
// These category meanings are separate from version-series identity and do not
// decide whether a change in a category constitutes a regression.
func StatusTones() map[string]text.Tone {
	return map[string]text.Tone{"2xx": text.Success, "3xx": text.Info, "4xx": text.Warning, "5xx": text.Danger}
}

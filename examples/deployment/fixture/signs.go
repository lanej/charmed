package fixture

import "github.com/lanej/charmed/text"

// StatusSigns is application display policy, independent of alert boundaries.
// Redirect/client-error changes are direction-only; their benefit needs context.
func StatusSigns() map[string]text.SignTones {
	return map[string]text.SignTones{
		"2xx": {Positive: text.Success, Negative: text.Danger},
		"3xx": {Positive: text.Info, Negative: text.Info},
		"4xx": {Positive: text.Info, Negative: text.Info},
		"5xx": {Positive: text.Danger, Negative: text.Success},
	}
}

package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextSignedNumber() string {
	change := text.SignedNumber{
		Value: .18, Available: true, Precision: 3, Unit: "pp", Emphasis: text.Bright,
		Tones: text.SignTones{Positive: text.Danger, Negative: text.Success},
	}

	return change.View(theme.New(theme.Dark))
}

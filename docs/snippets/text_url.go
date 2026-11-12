package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
)

func ExampleTextUrl() string {
	url := "https://api.example.test/v1/labels?expand=rate&locale=en_US"

	return layout.Column(
		text.FormatURL(url, text.PathOnly),
		text.FormatURL(url, text.PathAndQuery),
		text.FormatURL(url, text.FullURL),
	)
}

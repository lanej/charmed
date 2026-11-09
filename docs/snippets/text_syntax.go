package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextSyntax() string {
	body := `{
	  "error": "upstream_timeout",
	  "retryable": true,
	  "attempt": 2
	}`

	return text.Syntax(body, "json", theme.New(theme.Dark))
}

package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextLink() (string, error) {
	link := text.Link{Label: "PR #42", URL: "https://github.com/lanej/charmed/pull/42"}

	compact, err := link.View(theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	plain, err := link.Plain()
	if err != nil {
		return "", err
	}

	return compact + "\n" + plain, nil
}

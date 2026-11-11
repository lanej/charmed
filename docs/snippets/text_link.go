package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextLink() (string, error) {
	link, err := text.NewWebLink("PR #42", "https://github.com/lanej/charmed/pull/42")
	if err != nil {
		return "", err
	}

	compact, err := link.View(theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	plain, err := link.Plain()
	if err != nil {
		return "", err
	}

	session := text.Link{Label: "Worker session", URL: "terminal://session/build/pane/tests"}

	sessionView, err := session.View(theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	return compact + "\n" + sessionView + "\n" + plain, nil
}

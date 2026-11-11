package story

import (
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func linkStory(e Environment) (runtime.Scene, error) {
	palette := e.Config.Palette()
	webTargets := []text.Link{
		{Label: "PR #42", URL: "https://github.com/lanej/charmed/pull/42"},
		{Label: "Logs", URL: "https://example.test/builds/42/logs?step=test&attempt=1#output"},
		{Label: "Artifact", URL: "https://example.test/builds/42/artifacts/charmed.tar.gz"},
	}

	var compact, plain []string

	for _, target := range webTargets {
		link, err := text.NewWebLink(target.Label, target.URL)
		if err != nil {
			return nil, err
		}

		view, err := link.View(palette)
		if err != nil {
			return nil, err
		}

		fallback, err := link.Plain()
		if err != nil {
			return nil, err
		}

		compact = append(compact, view)
		plain = append(plain, fallback)
	}

	session := text.Link{Label: "Worker session", URL: "terminal://session/build/pane/tests"}

	sessionView, err := session.View(palette)
	if err != nil {
		return nil, err
	}

	sessionPlain, err := session.Plain()
	if err != nil {
		return nil, err
	}

	return static(func(size layout.Size) string {
		row := strings.Join(compact, "  ")
		if size.Width < 40 {
			row = strings.Join(compact, "\n")
		}

		rows := []string{palette.Heading().Render("Compact links"), "", "Build passed", row, "", sessionView, palette.Subtle().Render("Custom schemes need an opening handler."), "", palette.Subtle().Render("Plain output includes the destination:")}
		for _, value := range []string{plain[0], sessionPlain} {
			rows = append(rows, lipgloss.Wrap(value, size.Width, ""))
		}

		return strings.Join(rows, "\n")
	}), nil
}

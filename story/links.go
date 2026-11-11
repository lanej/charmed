package story

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func sampleLinks() ([]text.Link, error) {
	var links []text.Link

	for _, target := range []struct{ label, url string }{
		{"Logs", "https://example.test/builds/42/logs"},
		{"Artifact", "https://example.test/builds/42/artifacts/charmed.tar.gz"},
	} {
		link, err := text.NewWebLink(target.label, target.url)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return append(links, text.Link{Label: "Worker session", URL: "terminal://session/build/pane/tests", Emphasis: text.Dim}), nil
}

func linkStory(e Environment) (runtime.Scene, error) {
	palette := e.Config.Palette()

	link, err := text.NewWebLink("PR #42", "https://github.com/lanej/charmed/pull/42")
	if err != nil {
		return nil, err
	}

	rows := []string{palette.Heading().Render("Link / prominence"), ""}

	for _, emphasis := range []text.Emphasis{text.Normal, text.Bright, text.Dim} {
		link.Emphasis = emphasis

		view, renderErr := link.View(palette)
		if renderErr != nil {
			return nil, renderErr
		}

		rows = append(rows, palette.Subtle().Render(fmt.Sprintf("%-8s", emphasis))+view)
	}

	rows = append(rows, "", palette.Heading().Render("Presentation"))

	for _, presentation := range []text.Presentation{text.GlyphAndLabel, text.LabelOnly, text.GlyphOnly} {
		link.Presentation, link.Emphasis = presentation, text.Normal

		view, renderErr := link.View(palette)
		if renderErr != nil {
			return nil, renderErr
		}

		rows = append(rows, palette.Subtle().Render(fmt.Sprintf("%-8s", presentation))+view)
	}

	plain, err := link.Plain()
	if err != nil {
		return nil, err
	}

	rows = append(rows, "", palette.Heading().Render("Explicit plain output"))

	return static(func(size layout.Size) string {
		return strings.Join(rows, "\n") + "\n" + lipgloss.Wrap(plain, size.Width, "")
	}), nil
}

func linkGroupStory(e Environment) (runtime.Scene, error) {
	links, err := sampleLinks()
	if err != nil {
		return nil, err
	}

	palette := e.Config.Palette()
	group := layout.LinkGroup{Links: links}
	dense := layout.LinkGroup{Links: append([]text.Link(nil), links...)}

	for i := range dense.Links {
		dense.Links[i].Presentation = text.LabelOnly
	}

	return static(func(size layout.Size) string {
		wide, err := group.View(size.Width, palette)
		if err != nil {
			return err.Error()
		}

		narrow, err := group.View(min(size.Width, 24), palette)
		if err != nil {
			return err.Error()
		}

		labels, err := dense.View(size.Width, palette)
		if err != nil {
			return err.Error()
		}

		header := layout.Header{Title: "Build #42", Actions: text.Pill("Passed", text.Success, palette, text.Normal)}

		return strings.Join([]string{
			header.View(size.Width, palette), wide, "",
			palette.Subtle().Render("Reflow / 24 cells"), narrow, "",
			palette.Subtle().Render("Dense / label only"), labels,
		}, "\n")
	}), nil
}

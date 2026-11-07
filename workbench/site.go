package workbench

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

//go:embed site/*
var siteFiles embed.FS

type guidePage struct {
	File, Title string
	Content     template.HTML
	Groups      []guideGroup
	Coverage    []Coverage
	Preview     *ManifestEntry
	Showcase    []showcaseItem
	Atoms       []string
}
type showcaseItem struct{ Title, Description, Link, Frame, NarrowFrame string }
type guideGroup struct {
	Name    string
	Entries []ManifestEntry
}

func storyPage(entry ManifestEntry) string {
	switch entry.Level {
	case "Foundations":
		return "foundations.html"
	case "Screens", "Templates":
		return "examples.html"
	default:
		return "components.html"
	}
}

func writeSite(out, root string, entries []ManifestEntry) error {
	directory := filepath.Join(out, "dist")
	// WHY: dist is generator-owned; old browser bundles must not survive a static rebuild.
	if err := os.RemoveAll(directory); err != nil {
		return err
	}
	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}
	pages := []guidePage{
		{File: "index.html", Title: "Charmed"},
		{File: "foundations.html", Title: "Foundations"},
		{File: "components.html", Title: "Components"},
		{File: "examples.html", Title: "Composition & examples"},
		{File: "deployment.html", Title: "Compare a deployment"},
		{File: "development.html", Title: "Develop & publish"},
		{File: "brand.html", Title: "Brand sheet"},
	}
	markdown := goldmark.New(goldmark.WithExtensions(extension.GFM))
	for i := range pages {
		page := &pages[i]
		source, err := os.ReadFile(filepath.Join(root, "docs", strings.TrimSuffix(page.File, ".html")+".md"))
		if err != nil {
			return err
		}
		var content bytes.Buffer
		if err := markdown.Convert(source, &content); err != nil {
			return err
		}
		// WHY: Only repository-authored Markdown passes through the HTML renderer.
		rendered := highlightMarkdown(content.String())
		page.Content = template.HTML(rendered)
		if page.File == "index.html" {
			t := theme.New(theme.Dark)
			var states []string
			for _, state := range []feedback.Status{feedback.OK, feedback.Warn, feedback.Error, feedback.Unknown, feedback.Off} {
				states = append(states, story.FormatFrame(state.View(text.Bright, t), story.Config{Theme: theme.Dark, Profile: "ansi256"}))
			}
			page.Atoms = states
			for _, item := range []struct{ ID, Title, Description string }{
				{"chart/timeseries", "Time series", "Fine-grained dot traces over shared time windows."},
				{"chart/bars", "Bars", "Compare categories with native terminal geometry."},
				{"chart/distributions", "Distributions", "Cohort means, spread and selected observations."},
				{"text/signed-number", "Signed values", "Direction, precision and semantic emphasis."},
			} {
				d, err := story.Find(item.ID)
				if err != nil {
					return err
				}
				c := d.Defaults
				c.Legend = item.ID == "chart/distributions" || item.ID == "chart/timeseries"
				c.Columns = 64
				c.Rows = 16
				c.Profile = "ansi256"
				frame, err := story.Frame(item.ID, c)
				if err != nil {
					return err
				}
				c.Columns = 40
				narrow, err := story.Frame(item.ID, c)
				if err != nil {
					return err
				}
				page.Showcase = append(page.Showcase, showcaseItem{item.Title, item.Description, "components.html#" + item.ID, frame, narrow})
			}
			for i := range entries {
				if entries[i].ID == "screen/services" {
					preview := entries[i]
					c := preview.Defaults
					c.Rows = 24
					c.Profile = "ansi256"
					frame, err := story.Frame(preview.ID, c)
					if err != nil {
						return err
					}
					preview.Snapshots = []Snapshot{{Name: "Default", Config: c, Frame: frame}}
					c.Columns = 40
					c.Rows = 28
					narrow, err := story.Frame(preview.ID, c)
					if err != nil {
						return err
					}
					preview.Snapshots = append(preview.Snapshots, Snapshot{Name: "Narrow", Config: c, Frame: narrow})
					page.Preview = &preview
					break
				}
			}
		}
		if page.File == "development.html" {
			page.Coverage = CoverageInventory()
		}
		families := []string{}
		for _, entry := range entries {
			if storyPage(entry) == page.File && !slices.Contains(families, entry.Family) {
				families = append(families, entry.Family)
			}
		}
		slices.Sort(families)
		for _, family := range families {
			group := guideGroup{Name: family}
			for _, entry := range entries {
				if storyPage(entry) == page.File && entry.Family == family {
					group.Entries = append(group.Entries, entry)
				}
			}
			page.Groups = append(page.Groups, group)
		}
	}
	funcs := template.FuncMap{
		"terminal": terminalHTML,
		"highlight": highlightCode,
		"name":     func(title string) string { parts := strings.Split(title, "/"); return parts[len(parts)-1] },
		"link": func(id string) (string, error) {
			for _, entry := range entries {
				if entry.ID == id {
					return storyPage(entry) + "#" + id, nil
				}
			}
			return "", fmt.Errorf("unknown component link %q", id)
		},
	}
	layout, err := template.New("page.html").Funcs(funcs).ParseFS(siteFiles, "site/page.html")
	if err != nil {
		return err
	}
	for _, page := range pages {
		var output bytes.Buffer
		if err := layout.Execute(&output, struct {
			Page  guidePage
			Pages []guidePage
		}{page, pages}); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(directory, page.File), output.Bytes(), 0644); err != nil {
			return err
		}
	}
	css, err := siteFiles.ReadFile("site/style.css")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(directory, "style.css"), append(css, []byte(terminalCSS())...), 0644); err != nil {
		return err
	}
	tabs, err := siteFiles.ReadFile("site/tabs.js")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(directory, "tabs.js"), tabs, 0644); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(directory, "assets"), 0755); err != nil {
		return err
	}
	assets, err := filepath.Glob(filepath.Join(root, "docs", "assets", "*"))
	if err != nil {
		return err
	}
	for _, path := range assets {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(directory, "assets", filepath.Base(path)), data, 0644); err != nil {
			return err
		}
	}
	return os.WriteFile(filepath.Join(directory, ".nojekyll"), nil, 0644)
}

package workbench

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"image/color"
	"os"
	"path/filepath"
	"strings"

	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

//go:embed site/*
var siteFiles embed.FS

type guidePage struct {
	File, Title      string
	Source, Category string
	Overview         *guideLevel
	Legacy           []guideLink
	Resource         bool
	Content          template.HTML
	Levels           []guideLevel
	Navigation       []guideLevel
	Coverage         []Coverage
	Preview          *ManifestEntry
	Showcase         []showcaseItem
	Atoms            []string
	Brand            brandPalette
}
type showcaseItem struct{ Title, Description, Link, Frame, NarrowFrame string }

func writeSite(out, root string, entries []ManifestEntry) error {
	directory := filepath.Join(out, "dist")
	// WHY: dist is generator-owned; old browser bundles must not survive a static rebuild.
	if err := os.RemoveAll(directory); err != nil {
		return err
	}

	if err := os.MkdirAll(directory, 0o755); err != nil { //nolint:gosec // WHY: Static servers need traversal access to public guide directories.
		return err
	}

	pages := []guidePage{
		{File: "index.html", Title: "Charmed", Source: "index.md"},
		{File: "getting-started.html", Title: "Getting started", Source: "getting-started.md"},
		{File: "recipes.html", Title: "Recipes", Source: "recipes.md"},
		{File: "documentation-standard.html", Title: "Documentation standard", Source: "documentation-standard.md", Resource: true},
		{File: "log-inspector.html", Title: "Log inspector", Source: "log-inspector.md", Resource: true},
		{File: "command-palette.html", Title: "Command palette", Source: "command-palette.md", Resource: true},
		{File: "components.html", Title: "Component reference", Source: "components.md", Resource: true},
		{File: "examples.html", Title: "Composition guide", Source: "examples.md", Resource: true},
		{File: "deployment.html", Title: "Compare a deployment", Source: "deployment.md", Resource: true},
		{File: "json-viewer.html", Title: "JSON viewer", Source: "json-viewer.md", Resource: true},
		{File: "development.html", Title: "Develop & publish", Source: "development.md", Resource: true},
		{File: "brand.html", Title: "Brand sheet", Source: "brand.md", Resource: true},
	}
	navigation := atomicNavigation(entries)

	if err := addSpecimens(navigation); err != nil {
		return err
	}

	pages = append(pages, catalogPages(navigation)...)

	brand := newBrandPalette(theme.New(theme.Dark))
	markdown := goldmark.New(goldmark.WithExtensions(extension.GFM), goldmark.WithParserOptions(parser.WithAutoHeadingID()))

	for i := range pages {
		page := &pages[i]
		page.Brand = brand

		if page.Source != "" {
			source, err := os.ReadFile(filepath.Join(root, "docs", page.Source)) //nolint:gosec // WHY: Source names come only from the compiled guide catalog, never request input.
			if err != nil {
				return err
			}

			var content bytes.Buffer
			if err := markdown.Convert(source, &content); err != nil {
				return err
			}

			page.Content = template.HTML(highlightMarkdown(content.String())) //nolint:gosec // WHY: Only repository-authored Markdown passes through the escaping Markdown and code renderers.
		}

		if page.File == "index.html" {
			t := theme.New(theme.Dark)

			var states []string
			for _, state := range feedback.Statuses() {
				states = append(states, story.FormatFrame(state.View(text.Bright, t), story.Config{Theme: theme.Dark, Profile: "truecolor"}))
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
				c.Profile = "truecolor"

				frame, err := story.Frame(item.ID, c)
				if err != nil {
					return err
				}

				c.Columns = 40

				narrow, err := story.Frame(item.ID, c)
				if err != nil {
					return err
				}

				page.Showcase = append(page.Showcase, showcaseItem{item.Title, item.Description, storyPage(ManifestEntry{Definition: d}) + "#" + item.ID, frame, narrow})
			}

			for i := range entries {
				if entries[i].ID == "screen/services" {
					preview := entries[i]
					c := preview.Defaults
					c.Rows = 24
					c.Profile = "truecolor"

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

		page.Navigation = navigation
		page.Legacy = legacyLinks(page.File, navigation)
	}

	funcs := template.FuncMap{
		"lower":              strings.ToLower,
		"familyName":         familyName,
		"terminal":           terminalHTML,
		"snapshotForeground": snapshotForeground,
		"snapshotBackground": snapshotBackground,
		"highlight":          highlightCode,
		"name":               func(title string) string { parts := strings.Split(title, "/"); return parts[len(parts)-1] },
		"link": func(id string) (string, error) {
			for _, entry := range entries {
				if entry.ID == id {
					return storyPage(entry) + "#" + id, nil
				}
			}

			return "", fmt.Errorf("unknown component link %q", id)
		},
	}

	layout, err := template.New("page.html").Funcs(funcs).ParseFS(siteFiles, "site/page.html", "site/mark.svg")
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

		rendered, sourceErr := documentationSource(output.String(), os.Getenv("GITHUB_SHA"))
		if sourceErr != nil {
			return sourceErr
		}

		if writeErr := writeGuideFile(directory, page.File, []byte(rendered)); writeErr != nil {
			return writeErr
		}
	}

	css, err := siteFiles.ReadFile("site/style.css")
	if err != nil {
		return err
	}

	if writeErr := writeGuideFile(directory, "style.css", []byte(string(css)+brand.style())); writeErr != nil {
		return writeErr
	}

	tabs, err := siteFiles.ReadFile("site/tabs.js")
	if err != nil {
		return err
	}

	if writeErr := writeGuideFile(directory, "tabs.js", tabs); writeErr != nil {
		return writeErr
	}

	if mkdirErr := os.MkdirAll(filepath.Join(directory, "assets"), 0o755); mkdirErr != nil { //nolint:gosec // WHY: Static servers need traversal access to public guide directories.
		return mkdirErr
	}
	// The downloadable asset and inline lockups share one vector source.
	mark, err := siteFiles.ReadFile("site/mark.svg")
	if err != nil {
		return err
	}

	if writeErr := writeGuideFile(directory, "assets/charmed-mark.svg", mark); writeErr != nil {
		return writeErr
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

		if writeErr := writeGuideFile(directory, "assets/"+filepath.Base(path), data); writeErr != nil {
			return writeErr
		}
	}

	return writeGuideFile(directory, ".nojekyll", nil)
}

// brandPalette exposes the native dark-mode color roles to the static site.
// The browser has no independent brand or data-series hex palette.
type brandPalette struct {
	Background, Surface, Text, Muted, Border, Accent string
	Success, Info, Warning, Danger                   string
	Series                                           []string
}

func colorHex(c color.Color) string {
	if c == nil {
		return "transparent"
	}

	r, g, b, _ := c.RGBA()

	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}

func newBrandPalette(t theme.Theme) brandPalette {
	b := brandPalette{
		Background: colorHex(t.Background), Surface: colorHex(t.Surface),
		Text: colorHex(t.Foreground), Muted: colorHex(t.Muted),
		Border: colorHex(t.Border), Accent: colorHex(t.Accent),
		Success: colorHex(t.Success), Info: colorHex(t.Info),
		Warning: colorHex(t.Warning), Danger: colorHex(t.Danger),
	}
	for _, c := range t.Series {
		b.Series = append(b.Series, colorHex(c))
	}

	return b
}

func (b brandPalette) style() string {
	styles := "\n:root{--bg:" + b.Background +
		";--surface:" + b.Surface +
		";--text:" + b.Text +
		";--muted:" + b.Muted +
		";--line:" + b.Border +
		";--accent:" + b.Accent +
		";--success:" + b.Success +
		";--info:" + b.Info +
		";--warning:" + b.Warning +
		";--danger:" + b.Danger + ";"

	for i, c := range b.Series {
		styles += fmt.Sprintf("--series-%d:%s;", i, c)
	}

	return styles + "}\n"
}

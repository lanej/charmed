package workbench

import (
	"fmt"
	"slices"
	"strings"

	"github.com/lanej/charmed/story"
)

type guideLevel struct {
	Name, File, Anchor, Description string
	Groups                          []guideGroup
	Count                           int
}

type guideGroup struct {
	Name, Family, File, Anchor string
	Entries                    []ManifestEntry
	Specimen                   guideSpecimen
}

type guideLink struct {
	Name, Href, Anchor string
	Hidden             bool
}

func storyPage(entry ManifestEntry) string {
	return strings.ToLower(entry.Level) + "-" + entry.Family + ".html"
}

func atomicNavigation(entries []ManifestEntry) []guideLevel {
	var levels []guideLevel

	for _, name := range story.Levels() {
		level := guideLevel{Name: name, File: strings.ToLower(name) + ".html", Anchor: "level-" + strings.ToLower(name), Description: levelDescription(name)}
		families := make(map[string][]ManifestEntry)

		for _, entry := range entries {
			if entry.Level == name {
				families[entry.Family] = append(families[entry.Family], entry)
				level.Count++
			}
		}

		var names []string
		for family := range families {
			names = append(names, family)
		}

		slices.Sort(names)

		for _, family := range names {
			level.Groups = append(level.Groups, guideGroup{
				Name: familyName(family), Family: family, File: storyPage(families[family][0]),
				Anchor: level.Anchor + "-" + family, Entries: families[family],
			})
		}

		levels = append(levels, level)
	}

	return levels
}

func levelDescription(level string) string {
	switch level {
	case "Foundations":
		return "Shared rules for color, type, spacing, geometry and interaction. Start here to keep every component consistent."
	case "Atoms":
		return "Controls and visual primitives, ready to compose."
	case "Molecules":
		return "Small compositions with one purpose: a labeled field, a metric with history, or a searchable command palette."
	case "Organisms":
		return "Coordinated regions with their own behavior, from forms and log inspectors to comparison panels."
	case "Templates":
		return "Reusable screen structure. Define the layout and ownership of regions before supplying application content."
	case "Screens":
		return "Complete working compositions with sample data, linked selection and keyboard interaction."
	default:
		return ""
	}
}

func familyName(family string) string {
	switch family {
	case "chart":
		return "Charts"
	case "collection":
		return "Collections"
	case "commands":
		return "Commands"
	case "data":
		return "Data"
	case "diagram":
		return "Diagrams"
	case "feedback":
		return "Feedback"
	case "form":
		return "Forms"
	case "graphics":
		return "Graphics"
	case "http":
		return "HTTP"
	case "input":
		return "Input"
	case "layout":
		return "Layout"
	case "logs":
		return "Logs"
	case "motion":
		return "Motion"
	case "navigation":
		return "Navigation"
	case "pattern":
		return "Patterns"
	case "text":
		return "Text"
	case "theme":
		return "Theme"
	default:
		return family
	}
}

func catalogPages(levels []guideLevel) []guidePage {
	var pages []guidePage

	for _, level := range levels {
		pages = append(pages, guidePage{
			File: level.File, Title: level.Name, Category: level.Name,
			Source: strings.TrimSuffix(level.File, ".html") + ".md", Overview: &level,
		})

		for _, group := range level.Groups {
			pages = append(pages, guidePage{
				File: group.File, Title: level.Name + " / " + group.Name, Category: level.Name,
				Levels: []guideLevel{{Name: level.Name, File: level.File, Anchor: level.Anchor, Groups: []guideGroup{group}}},
			})
		}
	}

	return pages
}

// legacyLinks keeps published bookmarks discoverable without duplicating previews.
// The browser forwards exact fragments; without JavaScript these are ordinary links.
func legacyLinks(page string, levels []guideLevel) []guideLink {
	var links []guideLink

	seen := make(map[string]bool)

	for _, level := range levels {
		legacy := "components.html"

		switch level.Name {
		case "Foundations":
			legacy = "foundations.html"
		case "Screens", "Templates":
			legacy = "examples.html"
		}

		if legacy != page {
			continue
		}

		if level.File != page {
			links = append(links, guideLink{Name: level.Name, Href: level.File + "#" + level.Anchor, Anchor: level.Anchor})
		}

		for _, group := range level.Groups {
			target := group.File + "#" + group.Anchor
			links = append(links, guideLink{Name: level.Name + " / " + group.Name, Href: target, Anchor: group.Anchor})

			if !seen[group.Family] {
				links = append(links, guideLink{Name: group.Name, Href: target, Anchor: "family-" + group.Family})
				seen[group.Family] = true
			}

			for _, entry := range group.Entries {
				links = append(links, guideLink{Name: entry.Title, Href: group.File + "#" + entry.ID, Anchor: entry.ID})
				for _, suffix := range []string{"-preview", "-code"} {
					links = append(links, guideLink{Name: entry.Title + suffix, Href: group.File + "#" + entry.ID + suffix, Anchor: entry.ID + suffix, Hidden: true})
				}

				for i := range entry.Snapshots {
					for _, prefix := range []string{"tab-", "panel-"} {
						anchor := fmt.Sprintf("%s%s-%d", prefix, entry.ID, i)
						links = append(links, guideLink{Name: entry.Title, Href: group.File + "#" + anchor, Anchor: anchor, Hidden: true})
					}
				}
			}
		}
	}

	return links
}

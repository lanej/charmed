package workbench

import (
	"slices"
	"strings"

	"github.com/lanej/charmed/story"
)

type guideLevel struct {
	Name, Anchor string
	Groups       []guideGroup
}

type guideLink struct {
	Name, Href string
}

func atomicNavigation() []guideLink {
	var links []guideLink

	for _, level := range story.Levels() {
		page := storyPage(ManifestEntry{Definition: story.Definition{Level: level}})
		links = append(links, guideLink{Name: level, Href: page + "#level-" + strings.ToLower(level)})
	}

	return links
}

func groupLevels(page string, entries []ManifestEntry) []guideLevel {
	var levels []guideLevel

	seenFamilies := make(map[string]bool)

	for _, link := range atomicNavigation() {
		file, anchor, _ := strings.Cut(link.Href, "#")
		if file != page {
			continue
		}

		level := guideLevel{Name: link.Name, Anchor: anchor}
		families := make(map[string][]ManifestEntry)

		for _, entry := range entries {
			if entry.Level == level.Name {
				families[entry.Family] = append(families[entry.Family], entry)
			}
		}

		var names []string

		for family := range families {
			names = append(names, family)
		}

		slices.Sort(names)

		for _, family := range names {
			group := guideGroup{Name: family, Anchor: anchor + "-" + family, Entries: families[family]}
			if !seenFamilies[family] {
				group.LegacyAnchor = "family-" + family
				seenFamilies[family] = true
			}

			level.Groups = append(level.Groups, group)
		}

		levels = append(levels, level)
	}

	return levels
}

func navHref(href, page string) string {
	file, anchor, found := strings.Cut(href, "#")
	if found && file == page {
		return "#" + anchor
	}

	return href
}

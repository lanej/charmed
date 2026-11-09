package workbench

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/charmbracelet/x/ansi"
)

type guideSpecimen struct {
	ID, Name string
	Content  template.HTML
}

type specimenSelection struct {
	id          string
	first, rows int
}

// Select a useful excerpt of a saved native fixture, without its runner chrome.
// These are presentation choices for the guide, not alternate component renderers.
func specimenFor(level, family string) specimenSelection {
	selections := map[string]specimenSelection{
		"Foundations/layout":   {"foundations/layout", 2, 12},
		"Foundations/motion":   {"feedback/motion", 2, 3},
		"Foundations/text":     {"foundations/permissions", 4, 8},
		"Foundations/theme":    {"foundations/theme", 6, 7},
		"Atoms/chart":          {"chart/timeseries", 2, 10},
		"Atoms/collection":     {"collection/tree", 0, 9},
		"Atoms/feedback":       {"feedback/status", 2, 11},
		"Atoms/graphics":       {"chart/picture", 2, 11},
		"Atoms/http":           {"text/http-method", 2, 8},
		"Atoms/input":          {"input/text", 2, 3},
		"Atoms/logs":           {"text/observation", 2, 8},
		"Atoms/text":           {"text/pill", 2, 9},
		"Molecules/chart":      {"chart/distributions", 0, 5},
		"Molecules/commands":   {"collection/command-palette", 1, 12},
		"Molecules/diagram":    {"diagram/state-transitions", 2, 12},
		"Molecules/feedback":   {"feedback/percent", 2, 3},
		"Molecules/input":      {"input/field", 0, 3},
		"Molecules/layout":     {"layout/header", 0, 2},
		"Molecules/logs":       {"collection/http-log", 0, 10},
		"Molecules/navigation": {"navigation/tabs", 2, 1},
		"Organisms/chart":      {"chart/comparison-panel", 0, 12},
		"Organisms/collection": {"collection/list", 0, 12},
		"Organisms/form":       {"form/settings", 0, 12},
		"Organisms/logs":       {"pattern/log-inspector", 0, 8},
		"Organisms/text":       {"text/markdown", 0, 12},
		"Templates/pattern":    {"template/application", 0, 12},
		"Screens/pattern":      {"screen/services", 0, 12},
	}

	return selections[level+"/"+family]
}

func addSpecimens(levels []guideLevel) error {
	for i := range levels {
		for j := range levels[i].Groups {
			group := &levels[i].Groups[j]
			selection := specimenFor(levels[i].Name, group.Family)

			if selection.id == "" {
				selection = specimenSelection{id: group.Entries[0].ID, rows: 12}
			}

			found := false

			for _, entry := range group.Entries {
				if entry.ID != selection.id {
					continue
				}

				for _, snapshot := range entry.Snapshots {
					if snapshot.Name != "Narrow" {
						continue
					}

					content, err := specimenContent(snapshot.Frame, selection.first, selection.rows)
					if err != nil {
						return fmt.Errorf("catalog specimen %s: %w", entry.ID, err)
					}

					parts := strings.Split(entry.Title, "/")
					group.Specimen = guideSpecimen{ID: entry.ID, Name: parts[len(parts)-1], Content: content}
					found = true
				}
			}

			if !found {
				return fmt.Errorf("catalog group %s is missing a narrow snapshot for %s", group.File, selection.id)
			}
		}
	}

	return nil
}

func specimenContent(frame string, first, rows int) (template.HTML, error) {
	plain := strings.Split(ansi.Strip(frame), "\n")
	if first >= len(plain) || strings.TrimSpace(strings.Join(plain[first:min(first+rows, len(plain))], "\n")) == "" {
		return "", fmt.Errorf("selected native excerpt is empty")
	}

	// terminalHTML closes every span before each newline, so excerpts retain
	// complete markup and exact RGB even when ANSI styling spans source lines.
	lines := strings.Split(string(terminalHTML(frame)), "\n")
	last := min(first+rows, len(lines))

	return template.HTML(strings.Join(lines[first:last], "\n")), nil //nolint:gosec // WHY: Each line is complete escaped markup produced by terminalHTML, never caller-supplied HTML.
}

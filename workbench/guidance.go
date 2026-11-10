package workbench

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type componentGuidance struct {
	Use         map[string]string `json:"use"`
	Integration map[string]string `json:"integration"`
}

func loadGuidance(root string, entries []ManifestEntry) error {
	data, err := os.ReadFile(filepath.Join(root, "docs", "component-guidance.json")) //nolint:gosec // WHY: The filename is a fixed repository-owned documentation source.
	if err != nil {
		return err
	}

	var guidance componentGuidance
	if err := json.Unmarshal(data, &guidance); err != nil {
		return err
	}

	kinds := map[string]string{
		"chart/timeseries": "timeseries", "input/text": "input", "input/field": "input",
		"input/textarea": "textarea", "collection/files": "files", "pattern/log-inspector": "inspector",
		"collection/json-viewer": "json", "collection/command-palette": "palette", "navigation/tabs": "navigation",
		"text/markdown": "markdown", "feedback/spinner": "spinner", "feedback/time": "timers", "feedback/progress": "progress", "feedback/paginator": "interactive", "foundations/focus": "layout",
	}

	for i := range entries {
		entry := &entries[i]
		entry.Use = entry.Description

		if use := guidance.Use[entry.ID]; use != "" {
			entry.Use = use
		}

		kind := "static"

		switch {
		case entry.Family == "chart":
			kind = "chart"
		case entry.Family == "layout":
			kind = "layout"
		case entry.Family == "logs" && strings.HasPrefix(entry.ID, "collection/"):
			kind = "logs"
		case entry.Level == "Organisms" || entry.Level == "Templates" || entry.Level == "Screens":
			kind = "composed"
		case entry.Family == "collection":
			kind = "collection"
		}

		if override := kinds[entry.ID]; override != "" {
			kind = override
		}

		entry.Integration = guidance.Integration[kind]
		if entry.Integration == "" {
			return fmt.Errorf("%s has no integration guidance for %s", entry.ID, kind)
		}
	}

	return nil
}

// documentationSource pins source links to the revision under review in CI.
func documentationSource(rendered, revision string) (string, error) {
	if revision == "" {
		return rendered, nil
	}

	if len(revision) != 40 && len(revision) != 64 {
		return "", fmt.Errorf("invalid documentation source revision length")
	}

	if _, err := hex.DecodeString(revision); err != nil {
		return "", fmt.Errorf("invalid documentation source revision: %w", err)
	}

	return strings.ReplaceAll(rendered, "https://github.com/lanej/charmed/blob/main/", "https://github.com/lanej/charmed/blob/"+revision+"/"), nil
}

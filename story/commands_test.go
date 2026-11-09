package story_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestCommandConfigurationRejectsUnknownPresetsAndBackdrops(t *testing.T) {
	config := story.DefaultConfig()

	config.KeyLayout = "unknown"
	if config.Validate() == nil {
		t.Fatal("unknown key layout accepted")
	}

	config.KeyLayout = interaction.VimKeys

	config.PaletteBackdrop = "unknown"
	if config.Validate() == nil {
		t.Fatal("unknown backdrop accepted")
	}

	config.PaletteBackdrop = layout.BackdropMonochrome
	if err := config.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestCommandPaletteFixtureInitialFocusAndDeterministicStates(t *testing.T) {
	definition, err := story.Find("collection/command-palette")
	if err != nil {
		t.Fatal(err)
	}

	config := definition.Defaults

	root, err := story.New(definition.ID, config, nil)
	if err != nil {
		t.Fatal(err)
	}

	if root.Init() == nil {
		t.Fatal("initially open live palette lost focus command")
	}

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		config.Theme = mode
		for _, state := range []string{"ready", "disabled", "empty"} {
			config.State = state

			first, err := story.Frame(definition.ID, config)
			if err != nil {
				t.Fatal(err)
			}

			second, err := story.Frame(definition.ID, config)
			if err != nil {
				t.Fatal(err)
			}

			if first != second || !strings.Contains(ansi.Strip(first), "Commands") {
				t.Fatal("palette fixture lost deterministic state", mode, state)
			}
		}
	}
}

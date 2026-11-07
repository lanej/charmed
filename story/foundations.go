package story

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func foundationsDefinitions() []Definition {
	d := []Definition{
		widget("foundations/theme", "Foundations/Theme", "Foundations", "theme", "Semantic roles, terminal palettes, contrast and no-color output.", "themeStory", nil, themeStory),
		widget("foundations/layout", "Foundations/Layout", "Foundations", "layout", "Cell-based panels and responsive allocation.", "layoutStory", nil, layoutStory),
		widget("feedback/motion", "Foundations/Motion", "Foundations", "motion", "A Harmonica spring stepped explicitly by the owner; reduced motion resolves immediately.", "motionStory", []string{"amount"}, motionStory),
	}

	for i := range d {
		if d[i].ID == "feedback/motion" {
			d[i].Upstream = "https://github.com/charmbracelet/harmonica"
			d[i].Step = true
		}
	}
	return d
}

func themeStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()
		return title("Semantic tokens", e) + strings.Join([]string{t.Text().Render("Foreground / primary content"), t.Subtle().Render("Muted / supporting information"), t.Heading().Render("Accent / interactive emphasis"), t.Focus().Render("Focus / inverse + emphasis"), t.Text().Foreground(t.Success).Render("[OK] Success"), t.Text().Foreground(t.Warning).Render("[!] Warning"), t.Text().Foreground(t.Danger).Render("[x] Danger")}, "\n\n") + footer("Change theme, terminal size, or ASCII in Controls.", e)
	}), nil
}

func layoutStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		t := e.Config.Palette()
		left, right, split := layout.Split(size.Width, 1, 20, 30, .4)
		height := max(3, size.Height-4)
		a := layout.Panel("Navigation", "Cell-based dimensions\nBorders included in size", layout.Size{Width: left, Height: height}, t, true)
		b := layout.Panel("Content", "The right pane adapts to available space.\nNarrow terminals stack the same panels.", layout.Size{Width: right, Height: height}, t, false)
		if split {
			return title("Responsive layout", e) + layout.Row(a, " ", b)
		}
		height = max(3, height/2)
		return title("Stacked layout", e) + layout.Column(layout.Panel("Navigation", "Same content", layout.Size{Width: left, Height: height}, t, true), layout.Panel("Content", "Same state", layout.Size{Width: right, Height: height}, t, false))
	}), nil
}

func motionStory(e Environment) (runtime.Scene, error) {
	m := feedback.NewMotion(e.Config.Palette())
	return &scene{update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			m.Step(e.Config.Amount)
			emit(e, map[string]any{"type": "motion", "position": m.Position})
		}
		return nil
	}, view: func(layout.Size) string {
		return title("Spring motion", e) + fmt.Sprintf("Position %.2f  /  Target %.0f\n\nVelocity %.2f", m.Position, e.Config.Amount, m.Velocity) + footer("Ctrl+N steps the spring. Reduced motion resolves immediately.", e)
	}}, nil
}

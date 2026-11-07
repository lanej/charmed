// Run with: go run ./examples/deployment
package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/theme"
)

func main() {
	data, err := fixture.Data("isolated-hosts")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	palette := theme.New(theme.Dark)
	screen, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{
		Title: "Deployment comparison", Scope: "method=blue/green / Synthetic HTTP responses / all routes",
		Data: data, Theme: palette, MinRequests: 1000, Sigma: 3, MetricTones: fixture.StatusTones(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	size := layout.Size{Width: 132, Height: 40}
	screen.Resize(size)
	root := &runtime.Root{Scene: screen, Size: size, Background: palette.Background, AltScreen: true}
	if _, err := tea.NewProgram(root).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

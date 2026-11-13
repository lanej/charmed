package snippets

import (
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/examples/cohorts/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func ExamplePatternMeasurementHistory() (string, error) {
	history, err := pattern.NewMeasurementHistory(fixture.Data(), cohort.MemberKey{CohortID: "west", MemberID: "sensor-4"}, layout.Size{Width: 80, Height: 16}, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}
	// Temperature is missing for this sensor; other measurements remain observable.
	// SetData copies and validates. SetMember resets detail scroll, preserving the coordinate.
	// Parent owns Tab through history.Inspector.FocusDetail and forwards Update commands.
	return history.View(), nil
}

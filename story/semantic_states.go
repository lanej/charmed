package story

import (
	"strings"

	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type stateAtom interface {
	View(text.Emphasis, theme.Theme, ...text.Presentation) string
}

func stateAtomsStory(states ...stateAtom) func(Environment) (runtime.Scene, error) {
	return func(e Environment) (runtime.Scene, error) {
		return static(func(layout.Size) string {
			t := e.Config.Palette()
			emphasis := text.Emphasis(e.Config.Emphasis)

			rows := []string{"Pill                Compact", ""}
			for _, state := range states {
				rows = append(rows, t.Text().Width(20).Render(state.View(emphasis, t))+state.View(emphasis, t, text.Compact))
			}

			return strings.Join(rows, "\n")
		}), nil
	}
}

func semanticStateDefinitions() []Definition {
	return []Definition{
		{ID: "feedback/availability", Title: "Atoms/Feedback/Availability", Level: "Atoms", Family: "feedback", Description: "Up and Down describe supplied reachability without asserting readiness, health or traffic.", Source: "feedback/availability.go", Symbol: "Availability", Controls: []string{"emphasis"}, Composes: []string{"text/indicator"}, Defaults: atomConfig(), Factory: stateAtomsStory(feedback.Up, feedback.Down, feedback.AvailabilityUnknown)},
		{ID: "feedback/health", Title: "Atoms/Feedback/Health", Level: "Atoms", Family: "feedback", Description: "Healthy and Unhealthy describe health-check results independently of availability and rollout.", Source: "feedback/health.go", Symbol: "Health", Controls: []string{"emphasis"}, Composes: []string{"text/indicator"}, Defaults: atomConfig(), Factory: stateAtomsStory(feedback.Healthy, feedback.Unhealthy, feedback.HealthUnknown)},
		{ID: "feedback/readiness", Title: "Atoms/Feedback/Readiness", Level: "Atoms", Family: "feedback", Description: "Ready means the target may accept work; Not ready does not imply unhealthy or down.", Source: "feedback/readiness.go", Symbol: "Readiness", Controls: []string{"emphasis"}, Composes: []string{"text/indicator"}, Defaults: atomConfig(), Factory: stateAtomsStory(feedback.Ready, feedback.NotReady, feedback.ReadinessUnknown)},
		{ID: "feedback/rollout", Title: "Atoms/Feedback/Rollout", Level: "Atoms", Family: "feedback", Description: "Canary identifies the limited-rollout cohort; its pool membership is shown separately by CanaryOn or CanaryOff. Live denotes production traffic. Neither asserts health or means log-follow mode.", Source: "feedback/rollout.go", Symbol: "Rollout", Controls: []string{"emphasis"}, Composes: []string{"text/indicator"}, Defaults: atomConfig(), Factory: stateAtomsStory(feedback.Canary, feedback.Live, feedback.RolloutUnknown)},
		{ID: "feedback/canary-membership", Title: "Atoms/Feedback/Canary membership", Level: "Atoms", Family: "feedback", Description: "Canary on means the new-version target is in the serving pool; Canary off means it is out of the serving pool. Pool membership remains independent of health, readiness and reachability.", Source: "feedback/canary_membership.go", Symbol: "CanaryMembership", Controls: []string{"emphasis"}, Composes: []string{"text/indicator"}, Defaults: atomConfig(), Factory: stateAtomsStory(feedback.CanaryOn, feedback.CanaryOff, feedback.CanaryMembershipUnknown)},
		{ID: "text/indicator", Title: "Atoms/Text/Indicator", Level: "Atoms", Family: "text", Description: "A shared semantic glyph and label renderer: labeled pill, glyph-only pill, label-only pill or bare compact glyph. Content, emphasis and meaning remain independent.", Source: "text/text.go", Symbol: "Indicator", Controls: []string{"emphasis"}, Composes: []string{"text/pill"}, Defaults: atomConfig(), Factory: indicatorStory},
	}
}

func indicatorStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()

		rows := []string{}
		for _, mode := range []text.Presentation{text.DefaultPresentation, text.GlyphOnly, text.LabelOnly, text.Compact} {
			rows = append(rows, text.Indicator(t.Glyph(theme.GlyphUp), "Up", text.Success, t, text.Emphasis(e.Config.Emphasis), mode))
		}

		return strings.Join(rows, "\n")
	}), nil
}

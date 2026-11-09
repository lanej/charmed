package snippets

import (
	"fmt"

	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackMotion() string {
	motion := feedback.NewMotion(theme.New(theme.Dark))

	target := 68.0
	for range 30 {
		motion.Step(target) // In an application, step once per animation tick.
	}

	return fmt.Sprintf("Position %.2f / target %.2f", motion.Position, target)
}

package snippets

import (
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackAttention() string {
	palette := theme.New(theme.Dark)
	rule := attention.Rule{
		Warn:  &attention.Bounds{Upper: attention.Limit(60)},
		Error: &attention.Bounds{Upper: attention.Limit(80)},
	}
	result := rule.Evaluate(68, true) // Value and availability are separate.

	return feedback.Attention(result.Level, palette)
	// With no boundaries the rule is Off; unavailable data is Unknown.
}

package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackEvidenceSummary() string {
	counts := []feedback.EvidenceCount{
		{State: feedback.EvidenceHealthy, N: 2},
		{State: feedback.EvidenceUnknown, N: 1},
		{State: feedback.EvidenceStale, N: 1},
	}

	return feedback.SummaryLine(counts, 80, theme.New(theme.Dark))
}

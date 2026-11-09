package snippets

import (
	"time"

	"github.com/lanej/charmed/text"
)

func ExampleTextTime() string {
	at := time.Date(2026, 11, 8, 12, 4, 0, 123000000, time.UTC)
	return text.TimeOfDay(at) + "  " + text.ShortDuration(842*time.Millisecond)
}

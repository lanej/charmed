package text_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dlclark/regexp2"
)

func init() {
	// Exercise idle restarts during short matches; configure before any test uses a lexer.
	regexp2.SetTimeoutCheckPeriod(time.Millisecond)
}

func TestLexerTimeoutDeadlineSurvivesIdle(t *testing.T) {
	pattern := regexp2.MustCompile("needle", 0)
	pattern.MatchTimeout = 250 * time.Millisecond
	// A linear prefix search lets the clock tick between deadline creation and checking.
	input := strings.Repeat("a", 1<<20) + "needle"
	match := func(stage string) {
		t.Helper()

		found, err := pattern.MatchString(input)
		if err != nil {
			message, _, _ := strings.Cut(err.Error(), " on input ")
			t.Fatalf("%s: %s", stage, message)
		}

		if !found {
			t.Fatalf("%s: literal match disappeared", stage)
		}
	}

	match("warm")
	time.Sleep(3 * time.Second)
	match("after idle")
}

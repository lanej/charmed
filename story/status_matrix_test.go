package story_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
)

func TestStatusMatrixLabelsUseEntityGlyphsWithoutChangingIDs(t *testing.T) {
	for _, ascii := range []bool{false, true} {
		config := story.DefaultConfig()
		config.ASCII = ascii

		frame, err := story.Frame("collection/status-matrix", config)
		if err != nil {
			t.Fatal(err)
		}

		balancer, host := "⇄ lb-public", "▤ host-01"
		if ascii {
			balancer, host = "L lb-public", "H host-01"
		}

		plain := ansi.Strip(frame)
		if !strings.Contains(plain, balancer) || !strings.Contains(plain, host) || !strings.Contains(plain, "public / host-01: OK") {
			t.Fatalf("entity labels or stable IDs lost with ASCII=%v:\n%s", ascii, plain)
		}
	}
}

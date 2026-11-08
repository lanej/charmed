package workbench

import (
	"strings"
	"testing"
)

func TestHeadingPermalinks(t *testing.T) {
	result := headingPermalinks(`<h2 id="use-charmed">Use Charmed</h2>`)

	if !strings.Contains(result, `href="#use-charmed"`) {
		t.Fatal("missing heading permalink")
	}

	if !strings.Contains(result, `id="use-charmed"`) {
		t.Fatal("missing stable heading target")
	}
}

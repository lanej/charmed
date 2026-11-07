package workbench

import (
	"strings"
	"testing"
)

func TestTerminalHTMLPreservesColorAndRejectsActiveContent(t *testing.T) {
	frame := "\x1b[38;5;114mOK\x1b[0m <script>alert(1)</script>\n\x1b[38;5;234;48;5;221m WARN \x1b[0m"
	rendered := string(terminalHTML(frame))
	for _, wanted := range []string{`term-fgx114`, `term-bgx221`, `OK`, `WARN`, `&lt;script&gt;`} {
		if !strings.Contains(rendered, wanted) {
			t.Errorf("missing %s in %s", wanted, rendered)
		}
	}
	if strings.Contains(rendered, "<script") {
		t.Fatal("terminal text became executable markup")
	}
	malicious := "\x1b]1338;url=https://example.com/tracking.png;alt=tracking\a"
	if strings.Contains(string(terminalHTML(malicious)), "<img") {
		t.Fatal("terminal escape emitted an external image")
	}
}

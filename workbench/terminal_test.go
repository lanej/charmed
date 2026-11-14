package workbench

import (
	"strings"
	"testing"
)

func TestTerminalHTMLPreservesColorAndRejectsActiveContent(t *testing.T) {
	frame := "\x1b[38;5;114mOK\x1b[0m <script>alert(1)</script>\n\x1b[38;5;234;48;5;221m WARN \x1b[0m"

	rendered := string(terminalHTML(frame))
	for _, wanted := range []string{`color:#87d787`, `background-color:#ffd75f`, `OK`, `WARN`, `&lt;script&gt;`} {
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

func TestTerminalHTMLKeepsExactRGBAndCellPadding(t *testing.T) {
	frame := "\x1b[38;2;98;156;115;48;2;16;23;34m dim \x1b[0m\n界é"

	rendered := string(terminalHTML(frame))
	for _, wanted := range []string{"color:#629c73;", "background-color:#101722;", "> dim </span>", "width:2ch;text-align:center;", "界</span>é"} {
		if !strings.Contains(rendered, wanted) {
			t.Fatalf("missing %q in %s", wanted, rendered)
		}
	}

	for _, malicious := range []string{"\x1b]8;;javascript:alert(1)\a<script>\x1b]8;;\a", "<img src=x onerror=alert(1)>"} {
		rendered := string(terminalHTML(malicious))
		if strings.Contains(rendered, "<img") || strings.Contains(rendered, "<script") || strings.Contains(rendered, "<a ") {
			t.Fatal("terminal content gained HTML capabilities")
		}
	}
}

func TestTerminalHTMLReservesWideEmojiCells(t *testing.T) {
	rendered := string(terminalHTML("🐤🟢 Canary on\n🐤⚪️ Canary off\n? Unknown"))
	if strings.Count(rendered, "display:inline-block;width:2ch;") != 4 {
		t.Fatal("emoji fallback fonts lost native cell allocations", rendered)
	}

	for _, label := range []string{"Canary on", "Canary off", "? Unknown"} {
		if !strings.Contains(rendered, label) {
			t.Fatal("cell wrappers lost labels", label)
		}
	}
}

func TestTerminalHTMLWideCellsKeepTextDecorations(t *testing.T) {
	rendered := string(terminalHTML("\x1b[4;9m界\x1b[0m"))
	if !strings.Contains(rendered, `text-decoration:underline line-through;`) || !strings.Contains(rendered, `text-decoration:inherit;">界`) {
		t.Fatal("wide cell lost underline or strike-through across its inline-block", rendered)
	}
}

func TestTerminalHTMLPreservesMonochromeInverse(t *testing.T) {
	rendered := string(terminalHTML("\x1b[7m focus \x1b[0m"))
	for _, want := range []string{"color:#101722;", "background-color:#eeeeee;", "> focus </span>"} {
		if !strings.Contains(rendered, want) {
			t.Fatalf("monochrome inverse lost %q: %s", want, rendered)
		}
	}
}

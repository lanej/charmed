package workbench

import (
	"html/template"
	"strings"
	"testing"
)

func TestHighlightGoPreservesSourceAndEscapesHTML(t *testing.T) {
	source := "func value() string { // note\n\treturn \"<unsafe>\"\n}"
	got := string(highlightGo(source))
	for _, part := range []string{
		`<span class="syn-keyword">func</span>`,
		`<span class="syn-keyword">return</span>`,
		`<span class="syn-comment">// note</span>`,
		`&lt;unsafe&gt;`,
	} {
		if !strings.Contains(got, part) {
			t.Errorf("missing %q in %s", part, got)
		}
	}
	if strings.Contains(got, "<unsafe>") {
		t.Fatal("source injected raw markup")
	}
}

func TestHighlightMarkdownIncludesInlineAndFencedStyles(t *testing.T) {
	markdown := "<p>Use <code>theme.New</code> and <code>cmd/workbench</code>.</p>" +
		"<pre><code class=\"language-go\">func example() { return }</code></pre>" +
		"<pre><code class=\"language-sh\">go run ./cmd/workbench list</code></pre>"
	got := highlightMarkdown(markdown)
	for _, part := range []string{
		`class="inline-code"`,
		`class="code-block"`,
		`class="language-go"`,
		`class="language-sh"`,
		`class="syn-keyword"`,
		`class="syn-command"`,
	} {
		if !strings.Contains(got, part) {
			t.Errorf("missing %q in %s", part, got)
		}
	}
}

func TestUnknownCodeLanguagesRemainEscaped(t *testing.T) {
	got := highlightCode("<svg onload=alert(1)>", "unsupported")
	if got != template.HTML("&lt;svg onload=alert(1)&gt;") {
		t.Fatalf("unexpected output: %s", got)
	}
}

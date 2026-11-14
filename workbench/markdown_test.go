package workbench

import (
	"strings"
	"testing"
)

func TestGuideMarkdownPreservesConsumerContent(t *testing.T) {
	source := `<html><body><nav>Global navigation</nav><main>
<h1><a href="#widget">Widget<span aria-hidden="true">#</span></a></h1>
<p>Use <code>Widget</code> with <a href="rendering.html#capture">snapshots</a>.</p>
<div class="variant-tablist">Duplicate tab labels</div>
<figure><figcaption>Narrow output</figcaption><pre class="terminal">┌─界─┐
│ a │
└───┘</pre></figure>
<pre><code class="language-go">if ready {
    fmt.Println(&quot;&lt;ok&gt;&quot;)
}</code></pre>
<table><thead><tr><th>API</th><th>Owner</th></tr></thead>
<tbody><tr><td>View</td><td>Application</td></tr></tbody></table>
<ul><li>Return commands<ul><li>Handle errors</li></ul></li></ul>
<img src="assets/example.png" alt="Native example">
<details><summary>More examples</summary><p>Retained detail</p></details>
<footer>Site footer</footer></main></body></html>`

	data, err := guideMarkdown(source, "components.html")
	if err != nil {
		t.Fatal(err)
	}

	text := string(data)
	for _, want := range []string{
		"Widget", "`Widget`", "https://lanej.io/charmed/rendering.html#capture",
		"┌─界─┐\n│ a │\n└───┘", "```go\nif ready {\n    fmt.Println(\"<ok>\")\n}",
		"| API", "Application", "Return commands", "Handle errors",
		"![Native example](https://lanej.io/charmed/assets/example.png)", "Retained detail",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %q in Markdown:\n%s", want, text)
		}
	}

	for _, unwanted := range []string{"Global navigation", "Duplicate tab labels", "Site footer", "<span", "<pre"} {
		if strings.Contains(text, unwanted) {
			t.Errorf("export contains %q", unwanted)
		}
	}
}

func TestGuideMarkdownRequiresMainContent(t *testing.T) {
	if _, err := guideMarkdown("<p>No main content</p>", "broken.html"); err == nil {
		t.Fatal("accepted a page without main content")
	}
}

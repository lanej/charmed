package workbench_test

import (
	"bytes"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/workbench"
	"golang.org/x/net/html"
)

func TestStaticGuide(t *testing.T) {
	out := t.TempDir()
	if err := workbench.Generate(out); err != nil {
		t.Fatal(err)
	}

	directory := filepath.Join(out, "dist")

	files, err := filepath.Glob(filepath.Join(directory, "*.html"))
	if err != nil {
		t.Fatal(err)
	}

	groups := make(map[string]bool)
	for _, d := range story.Definitions() {
		groups[strings.ToLower(d.Level)+"-"+d.Family+".html"] = true
	}

	if want := 7 + len(story.Levels()) + len(groups); len(files) != want {
		t.Fatalf("got %d pages, want %d category, group and guide pages", len(files), want)
	}

	pages := map[string][]byte{}
	ids := map[string]map[string]bool{}
	links := map[string][]string{}
	navComponents := map[string]map[string]bool{}
	headingLinks := map[string]map[string]bool{}

	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		name := filepath.Base(path)

		for _, internal := range []string{"Fixture implementation", "func Example", "package snippets", "e.Config", "return &amp;scene"} {
			if bytes.Contains(data, []byte(internal)) {
				t.Errorf("%s exposes example scaffolding or fixture implementation: %s", name, internal)
			}
		}

		pages[name] = data
		ids[name] = map[string]bool{}
		navComponents[name] = map[string]bool{}

		headingLinks[name] = map[string]bool{}
		if bytes.Contains(data, []byte("\x1b")) {
			t.Errorf("%s contains terminal escape sequences", name)
		}

		node, err := html.Parse(bytes.NewReader(data))
		if err != nil {
			t.Fatal(err)
		}

		var walk func(*html.Node)

		walk = func(n *html.Node) {
			if n.Type == html.ElementNode {
				switch n.Data {
				case "script":
					allowed := false

					for _, a := range n.Attr {
						if a.Key == "src" && a.Val == "tabs.js" {
							allowed = true
						}
					}

					if !allowed {
						t.Errorf("%s contains unexpected script", name)
					}
				case "input", "button", "select", "iframe":
					t.Errorf("%s contains unsupported interactive element %s", name, n.Data)
				}

				for _, a := range n.Attr {
					if a.Key == "id" {
						if ids[name][a.Val] {
							t.Errorf("duplicate id %s in %s", a.Val, name)
						}

						ids[name][a.Val] = true
					}

					if a.Key == "src" || a.Key == "href" {
						links[name] = append(links[name], a.Val)
					}
				}
			}

			if n.Type == html.ElementNode && n.Data == "a" {
				componentLink := false
				permalink := false
				href := ""

				for _, a := range n.Attr {
					if a.Key == "class" && strings.Contains(a.Val, "nav-component") {
						componentLink = true
					}

					if a.Key == "class" && strings.Contains(a.Val, "heading-permalink") {
						permalink = true
					}

					if a.Key == "href" {
						href = a.Val
					}
				}

				insideNav := false

				for p := n.Parent; p != nil; p = p.Parent {
					if p.Type == html.ElementNode && p.Data == "nav" {
						insideNav = true
						break
					}
				}

				if componentLink && insideNav {
					if u, err := url.Parse(href); err == nil {
						navComponents[name][u.Fragment] = true
					}
				}

				if permalink && n.Parent != nil && (n.Parent.Data == "h1" || n.Parent.Data == "h2" || n.Parent.Data == "h3" || n.Parent.Data == "h4" || n.Parent.Data == "h5") {
					if u, err := url.Parse(href); err == nil {
						headingLinks[name][u.Fragment] = true
					}
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				walk(c)
			}
		}
		walk(node)
	}

	for name, targets := range links {
		for _, target := range targets {
			u, err := url.Parse(target)
			if err != nil {
				t.Fatal(err)
			}

			if u.IsAbs() {
				continue
			}

			if strings.HasPrefix(u.Path, "/") {
				t.Errorf("%s has root-relative link %s", name, target)
			}

			file := u.Path
			if file == "" {
				file = name
			}

			if _, err := os.Stat(filepath.Join(directory, file)); err != nil {
				t.Errorf("%s broken link %s: %v", name, target, err)
			}

			if u.Fragment != "" && !ids[file][u.Fragment] {
				t.Errorf("%s missing anchor %s", name, target)
			}
		}
	}

	for _, d := range story.Definitions() {
		found := 0

		for _, pageIDs := range ids {
			if pageIDs[d.ID] {
				found++
			}
		}

		if found != 1 {
			t.Errorf("fixture %s appears %d times, want once", d.ID, found)
		}
	}

	if bytes.Contains(pages["index.html"], []byte("<table")) {
		t.Error("coverage inventory must not replace the home showcase")
	}

	if !bytes.Contains(pages["index.html"], []byte("background-color:#86cf9a")) {
		t.Error("showcase must preserve native status pill fills")
	}

	for _, state := range []string{"Empty", "Disabled"} {
		if !bytes.Contains(pages["atoms-input.html"], []byte("<figcaption>"+state+" · truecolor")) {
			t.Errorf("input guide omitted %s state", state)
		}
	}

	if !bytes.Contains(pages["foundations-text.html"], []byte(`style="color:#172033;background-color:#f8fafc"`)) {
		t.Error("light permission preview must retain its theme foreground and background")
	}

	if !bytes.Contains(pages["molecules-input.html"], []byte("&lt;")) {
		t.Error("the field validation example must escape its less-than operator")
	}

	if !bytes.Contains(pages["atoms-input.html"], []byte(`class="variant-tablist" role="tablist"`)) {
		t.Error("component previews must provide variant tabs")
	}

	if !bytes.Contains(pages["atoms-input.html"], []byte(`class="language-go"`)) ||
		!bytes.Contains(pages["atoms-input.html"], []byte("syn-keyword")) {
		t.Error("source listings must be highlighted")
	}

	if bytes.Contains(pages["atoms-input.html"], []byte("Saved Go output &amp; variants")) ||
		bytes.Contains(pages["atoms-input.html"], []byte("<summary>Go example</summary>")) {
		t.Error("visual previews and source must be immediately visible")
	}

	for _, page := range []string{"atoms-input.html", "foundations-theme.html", "screens-pattern.html"} {
		if bytes.Contains(pages[page], []byte("Run natively")) ||
			bytes.Contains(pages[page], []byte("component-run")) ||
			bytes.Contains(pages[page], []byte("go run ./cmd/workbench run ")) {
			t.Errorf("%s should not render per-component runner commands", page)
		}
	}

	if !bytes.Contains(pages["brand.html"], []byte("brand-swatches")) {
		t.Error("brand sheet must expose the visual token palette")
	}

	if _, err := os.Stat(filepath.Join(directory, "tabs.js")); err != nil {
		t.Fatalf("missing variant-tabs enhancement: %v", err)
	}

	for _, d := range story.Definitions() {
		target := strings.ToLower(d.Level) + "-" + d.Family + ".html"

		if !navComponents[target][d.ID] {
			t.Errorf("missing sidebar link for %s", d.ID)
		}

		for _, id := range []string{d.ID, d.ID + "-preview", d.ID + "-code", "level-" + strings.ToLower(d.Level) + "-" + d.Family} {
			if !headingLinks[target][id] {
				t.Errorf("%s: missing self-link on heading %q", target, id)
			}
		}
	}

	for _, level := range story.Levels() {
		page := strings.ToLower(level) + ".html"
		if !headingLinks[page]["level-"+strings.ToLower(level)] {
			t.Errorf("%s is missing its category heading", page)
		}

		if bytes.Contains(pages[page], []byte(`class="component"`)) {
			t.Errorf("%s must list groups without loading every preview", page)
		}
	}

	stale := filepath.Join(directory, "obsolete.js")
	if err := os.WriteFile(stale, []byte("old browser bundle"), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := workbench.Generate(out); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("stale output survived regeneration: %v", err)
	}

	for name, before := range pages {
		after, err := os.ReadFile(filepath.Join(directory, name))
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(before, after) {
			t.Errorf("%s changed on regeneration", name)
		}
	}
}

func TestEveryComponentHasAConsumerExample(t *testing.T) {
	for _, definition := range story.Definitions() {
		if definition.ExampleSource == "" || !strings.HasSuffix(definition.ExampleSource, ".go") {
			t.Errorf("%s needs a compiled consumer example", definition.ID)
			continue
		}

		source, err := os.ReadFile(filepath.Join("..", definition.ExampleSource))
		if err != nil {
			t.Fatal(err)
		}

		for _, internal := range []string{"github.com/lanej/charmed/story", "/deployment/fixture", "e.Config", "sampleSeries(", "sampleValues("} {
			if bytes.Contains(source, []byte(internal)) {
				t.Errorf("%s: consumer example depends on %s", definition.ID, internal)
			}
		}
	}
}

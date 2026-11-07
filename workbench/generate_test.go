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
	if len(files) != 6 {
		t.Fatalf("got %d pages, want six grouped guides", len(files))
	}
	pages := map[string][]byte{}
	ids := map[string]map[string]bool{}
	links := map[string][]string{}
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		name := filepath.Base(path)
		pages[name] = data
		ids[name] = map[string]bool{}
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
				case "script", "input", "button", "select", "iframe":
					t.Errorf("%s contains interactive element %s", name, n.Data)
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
	if !bytes.Contains(pages["index.html"], []byte("term-bgx")) {
		t.Error("showcase must preserve native status pill fills")
	}

	if !bytes.Contains(pages["components.html"], []byte("&lt;")) {
		t.Error("expected escaped source or snapshot content")
	}
	stale := filepath.Join(directory, "obsolete.js")
	if err := os.WriteFile(stale, []byte("old browser bundle"), 0644); err != nil {
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

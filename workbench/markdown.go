package workbench

import (
	"fmt"
	"strings"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/table"
	"golang.org/x/net/html"
)

func (p guidePage) MarkdownFile() string {
	return strings.TrimSuffix(p.File, ".html") + ".md"
}

func guideMarkdown(source, file string) ([]byte, error) {
	doc, err := html.Parse(strings.NewReader(source))
	if err != nil {
		return nil, err
	}

	main := markdownMain(doc)
	if main == nil {
		return nil, fmt.Errorf("%s has no main content", file)
	}

	cleanMarkdown(main)

	conv := converter.NewConverter(converter.WithPlugins(
		base.NewBasePlugin(), commonmark.NewCommonmarkPlugin(), table.NewTablePlugin(),
	))

	data, err := conv.ConvertNode(main, converter.WithDomain("https://lanej.io/charmed/"+file))
	if err != nil {
		return nil, fmt.Errorf("convert %s to Markdown: %w", file, err)
	}

	return append(data, '\n'), nil
}

func markdownMain(node *html.Node) *html.Node {
	if node.Type == html.ElementNode && node.Data == "main" {
		return node
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if main := markdownMain(child); main != nil {
			return main
		}
	}

	return nil
}

func cleanMarkdown(node *html.Node) {
	for child := node.FirstChild; child != nil; {
		next := child.NextSibling
		if markdownChrome(child) {
			node.RemoveChild(child)
		} else {
			cleanMarkdown(child)
		}

		child = next
	}
}

func markdownChrome(node *html.Node) bool {
	if node.Type != html.ElementNode {
		return false
	}

	switch node.Data {
	case "footer", "script", "style", "svg", "select", "label":
		return true
	}

	for _, attr := range node.Attr {
		if attr.Key == "aria-hidden" && attr.Val == "true" {
			return true
		}

		if attr.Key == "class" {
			for _, class := range strings.Fields(attr.Val) {
				switch class {
				case "variant-tablist", "graph-picker", "legacy-links", "markdown-download":
					return true
				}
			}
		}
	}

	return false
}

func writeMarkdownIndex(directory string, pages []guidePage) error {
	lines := []string{"# Charmed\n\nGo terminal components, composition and contributor reference. Each link contains the complete page as Markdown, including consumer examples and saved native output.\n\n## Pages\n"}

	for _, page := range pages {
		lines = append(lines, fmt.Sprintf("- [%s](https://lanej.io/charmed/%s)", page.Title, page.MarkdownFile()))
	}

	lines = append(lines, "\n## Skills\n")
	for _, name := range guideSkills {
		lines = append(lines, fmt.Sprintf("- [%s](https://lanej.io/charmed/skills/%s/SKILL.md)", name, name))
	}

	return writeGuideFile(directory, "llms.txt", []byte(strings.Join(lines, "\n")+"\n"))
}

package workbench

import (
	"fmt"
	"regexp"
)

var markdownHeading = regexp.MustCompile(`<h([2-4]) id="([a-zA-Z0-9_-]+)">([^<]+)</h[2-4]>`)

// headingPermalinks makes plain Markdown headings link to their own generated IDs.
// The source is repository-authored Markdown already escaped by Goldmark.
func headingPermalinks(rendered string) string {
	return markdownHeading.ReplaceAllStringFunc(rendered, func(match string) string {
		parts := markdownHeading.FindStringSubmatch(match)

		return fmt.Sprintf(
			`<h%s id="%s"><a class="heading-permalink" href="#%s"><span>%s</span><span class="heading-link-icon" aria-hidden="true">#</span></a></h%s>`,
			parts[1], parts[2], parts[2], parts[3], parts[1],
		)
	})
}

package workbench

import (
	"strings"
	"testing"
)

func TestDocumentationSourcePinsOnlyRepositorySourceLinks(t *testing.T) {
	revision := strings.Repeat("a", 40)
	source := `<a href="https://github.com/lanej/charmed/blob/main/examples/quickstart/main.go">Run</a><a href="getting-started.html">Guide</a>`

	got, err := documentationSource(source, revision)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(got, "/blob/"+revision+"/examples/quickstart/main.go") || !strings.Contains(got, `href="getting-started.html"`) {
		t.Fatal("preview source must use the reviewed revision while local navigation stays relative")
	}

	got, err = documentationSource(source, "")
	if err != nil || got != source {
		t.Fatal("local builds must retain their ordinary source links")
	}

	if _, err := documentationSource(source, strings.Repeat("z", 40)); err == nil {
		t.Fatal("non-hex revisions must fail rather than inject source-link markup")
	}

	if _, err := documentationSource(source, "bad"); err == nil {
		t.Fatal("incomplete revisions must fail")
	}
}

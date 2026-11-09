package workbench

import (
	"strings"
	"testing"

	"github.com/lanej/charmed/story"
)

func TestMissingConsumerExampleNeverFallsBackToImplementation(t *testing.T) {
	_, err := exampleCode(t.TempDir(), story.Definition{ID: "chart/line", Source: "chart/native.go", Symbol: "NewLine"})
	if err == nil || !strings.Contains(err.Error(), "consumer usage example") {
		t.Fatalf("missing example should be explicit, got %v", err)
	}
}

func TestExampleExtractionPreservesLiteralData(t *testing.T) {
	source := "package snippets\nimport tea \"charm.land/bubbletea/v2\"\nfunc ExampleText() string {\n\tvalue := `first\n\tsecond`\n\treturn value\n}\n"

	example, err := parseExample("text.go", []byte(source))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(example.Code, "`first\n\tsecond`") {
		t.Fatalf("dedenting changed the string's value: %s", example.Code)
	}

	for _, scaffolding := range []string{"package snippets", "import ", "func ExampleText"} {
		if strings.Contains(example.Code, scaffolding) {
			t.Errorf("published example includes %s", scaffolding)
		}
	}

	if len(example.Imports) != 1 || example.Imports[0].Name != "tea" || example.Imports[0].Path != "charm.land/bubbletea/v2" {
		t.Fatalf("import alias was lost: %+v", example.Imports)
	}
}

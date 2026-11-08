package workbench

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/lanej/charmed/theme"
)

func TestBrandTokensAreDerivedFromNativeTheme(t *testing.T) {
	native := theme.New(theme.Dark)
	brand := newBrandPalette(native)
	css := brand.style()

	for _, want := range []string{
		"--bg:" + colorHex(native.Background),
		"--surface:" + colorHex(native.Surface),
		"--text:" + colorHex(native.Foreground),
		"--muted:" + colorHex(native.Muted),
		"--line:" + colorHex(native.Border),
		"--accent:" + colorHex(native.Accent),
		"--success:" + colorHex(native.Success),
		"--info:" + colorHex(native.Info),
		"--warning:" + colorHex(native.Warning),
		"--danger:" + colorHex(native.Danger),
		"--series-0:" + colorHex(native.SeriesColor(0)),
		"--series-1:" + colorHex(native.SeriesColor(1)),
	} {
		if !strings.Contains(css, want+";") {
			t.Errorf("generated CSS missing canonical theme token %q", want)
		}
	}
}

func TestStylesheetDoesNotOwnItsOwnColorPalette(t *testing.T) {
	css, err := os.ReadFile("site/style.css")
	if err != nil {
		t.Fatal(err)
	}

	if regexp.MustCompile(`#[0-9a-fA-F]{6}\b`).Match(css) {
		t.Fatal("site stylesheet contains hardcoded hex values; use theme semantic tokens")
	}
}

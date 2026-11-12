package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutIdentityHeader() string {
	headline := layout.IdentityHeader{
		Kind: "Request", ID: "01234567-89ab-cdef-0123-456789abcdef",
		Context: "POST /v2/shipments/shp_fixture/buy · HTTP 500",
	}

	return headline.View(64, theme.New(theme.Dark))
}

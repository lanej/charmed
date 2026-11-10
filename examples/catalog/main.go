// Command catalog renders the same consumer example displayed in the guide.
package main

import (
	"fmt"
	"os"

	"github.com/lanej/charmed/docs/snippets"
)

func main() {
	if len(os.Args) != 2 {
		if _, err := fmt.Fprintln(os.Stderr, "usage: go run ./examples/catalog <component-id>"); err != nil {
			os.Exit(1)
		}

		os.Exit(1)
	}

	frame, err := snippets.Render(os.Args[1])
	if err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(1)
		}

		os.Exit(1)
	}

	if _, err := fmt.Println(frame); err != nil {
		os.Exit(1)
	}
}

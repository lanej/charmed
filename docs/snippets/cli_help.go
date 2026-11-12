package snippets

import (
	"github.com/spf13/cobra"

	"github.com/lanej/charmed/cli"
	"github.com/lanej/charmed/theme"
)

func ExampleCLIHelp() string {
	show := &cobra.Command{
		Use: "release VERSION", Short: "Show the services in a release.",
		Example: "release v1.4 --format json",
	}
	show.Flags().StringP("format", "f", "table", "Output format")
	show.InitDefaultHelpFlag()

	formatter := cli.HelpFormatter{Width: 56, Theme: theme.New(theme.Dark)}

	return formatter.Format(show)
}

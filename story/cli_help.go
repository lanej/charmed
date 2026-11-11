package story

import (
	"github.com/spf13/cobra"

	"github.com/lanej/charmed/cli"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func cliHelpStory(e Environment) (runtime.Scene, error) {
	root := &cobra.Command{Use: "release"}
	root.PersistentFlags().StringP("profile", "p", "local", "Configuration profile to use")

	show := &cobra.Command{
		Use: "show VERSION", Short: "Show the services in a release.",
		Aliases: []string{"inspect"}, Example: "release show v1.4 --format json",
	}
	show.Flags().StringP("format", "f", "table", "Output format")
	show.Flags().Bool("changed", false, "Show only services that changed in this release")
	show.InitDefaultHelpFlag()
	root.AddCommand(show)

	return static(func(size layout.Size) string {
		return (cli.HelpFormatter{Width: size.Width, Theme: e.Config.Palette()}).Format(show)
	}), nil
}

package main

import (
	"fmt"
	"os"

	"charm.land/fang/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	mango "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"

	"github.com/lanej/charmed/theme"
)

// Keep the workbench's Fang diagnostics and manpage alongside the Cobra help hook.
func printCommandError(cmd *cobra.Command, err error) {
	palette := theme.New(theme.Dark)
	styles := fang.Styles{
		ErrorHeader: lipgloss.NewStyle().Foreground(palette.Background).Background(palette.Danger).Bold(true).Padding(0, 1).SetString("ERROR"),
		ErrorText:   lipgloss.NewStyle(),
		Program:     fang.Program{Flag: palette.Heading()},
	}
	fang.DefaultErrorHandler(colorprofile.NewWriter(cmd.ErrOrStderr(), os.Environ()), styles, err)
}

func manCommand() *cobra.Command {
	return &cobra.Command{
		Use: "man", Short: "Generates manpages", Hidden: true,
		DisableFlagsInUseLine: true, Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			page, err := mango.NewManPage(1, cmd.Root())
			if err != nil {
				return err
			}

			_, err = fmt.Fprint(cmd.OutOrStdout(), page.Build(roff.NewDocument()))

			return err
		},
	}
}

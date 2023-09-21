package cmd

import (
	"os"

	"github.com/adrianfinantyo/jkt48-showroom-cli/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Show information about this CLI",
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColMinWidth(0, 15)
		table.SetColMinWidth(1, 32)

		table.Append([]string{"Version", utils.AppVersion})
		table.Append([]string{"Author", utils.AppAuthor})
		table.Append([]string{"License", utils.AppLicense})
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}

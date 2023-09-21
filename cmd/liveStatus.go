package cmd

import (
	"os"

	"github.com/adrianfinantyo/jkt48-showroom-cli/utils"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var liveStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show all JKT48 members live status",
	Run: func(cmd *cobra.Command, args []string) {

		utils.LogInfo("💫 Getting all JKT48 members live status...")
		progressBar := progressbar.Default(0)

		member := utils.GetAllJKT48Rooms(progressBar)
		progressBar.ChangeMax(len(*member))

		utils.ClearScreen()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Name", "Live Status"})

		for _, data := range *member {
			var liveStatus string
			if data.IsLive {
				liveStatus = color.GreenString("Live")
			} else {
				liveStatus = color.RedString("Not Live")
			}
			table.Append([]string{data.Name, liveStatus})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(liveStatusCmd)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)

}

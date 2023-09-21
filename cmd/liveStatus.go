package cmd

import (
	"os"

	"github.com/adrianfinantyo/jkt48-showroom-cli/models"
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
		memberChan := make(chan *[]models.Room)

		utils.LogInfo("💫 Getting all JKT48 members live status...")
		progressBar := progressbar.NewOptions(len(utils.AddedRooms), progressbar.OptionSetWidth(35))

		go utils.GetAllJKT48Rooms(progressBar, memberChan)

		member := <-memberChan

		utils.ClearScreen()

		utils.PrintHeader("JKT48 Showroom CLI", "Live Status")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Name", "Live Status"})

		for _, data := range *member {
			var liveStatus string
			if data.IsLive {
				liveStatus = color.GreenString("ONLINE")
			} else {
				liveStatus = color.RedString("OFFLINE")
			}
			table.Append([]string{data.Name, liveStatus})
		}
		table.Render()
	},
}

func init() {
	utils.PrintHeader("JKT48 Showroom CLI", "Live Status")

	rootCmd.AddCommand(liveStatusCmd)
}

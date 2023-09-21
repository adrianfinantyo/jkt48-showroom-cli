package cmd

import (
	"fmt"
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
		fmt.Println()
		progressBar := progressbar.NewOptions(len(utils.AddedRooms), progressbar.OptionSetWidth(35), progressbar.OptionOnCompletion(func() {
			fmt.Fprint(cmd.OutOrStdout(), "\n")
		}))

		go utils.GetAllJKT48Rooms(progressBar, memberChan)

		member := <-memberChan

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"Name", "Live Status"})

		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER})

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
	rootCmd.AddCommand(liveStatusCmd)
}

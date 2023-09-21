package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"github.com/olekukonko/tablewriter"
)

func LogError(err error) {
	if err != nil {
		color.Red(err.Error())
	} else {
		color.Red("Something went wrong")
	}
}

func LogInfo(msg string) {
	color.Blue(msg)
}

func ClearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func PrintHeader(title string, desc string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{title})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeaderColor(
		tablewriter.Color(tablewriter.BgCyanColor, tablewriter.FgHiWhiteColor, tablewriter.Bold),
	)
	table.SetColMinWidth(0, 50)
	table.Append([]string{desc})
	table.SetAutoWrapText(false)
	table.Render()
	fmt.Println()
}

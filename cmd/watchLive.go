package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/adrianfinantyo/jkt48-showroom-cli/models"
	"github.com/adrianfinantyo/jkt48-showroom-cli/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var watchLive = &cobra.Command{
	Use:   "watch",
	Short: "Watch JKT48 member live stream",
	Run: func(cmd *cobra.Command, args []string) {
		liveRooms := utils.GetActiveRooms()
		if len(*liveRooms) == 0 {
			utils.LogInfo("🤭 Oops, seems like there's no active rooms right now")
			os.Exit(0)
		}

		selectedRoom := promptGetSelectMember(liveRooms)
		if selectedRoom.RoomId != 0 {
			streamURL := promptSelectQuality(selectedRoom.StreamURL)
			startVLCService(streamURL)
		} else {
			err := fmt.Errorf("⚠️ Error: Room ID is 0")
			utils.LogError(err)
			os.Exit(1)
		}
	},
}

func init() {
	utils.PrintHeader("JKT48 Showroom CLI", "Watch Live Stream")
	rootCmd.AddCommand(watchLive)
}

func promptGetSelectMember(liveRooms *[]models.LiveRoom) models.LiveRoom {
	var items []string
	for _, room := range *liveRooms {
		label := fmt.Sprintf("%s (ID: %d)", room.RoomKey, room.RoomId)
		items = append(items, label)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | cyan }}?",
		Active:   "\U0001F449 {{ . | green | underline }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "\U0001F449 {{ . | green | underline }}",
	}

	prompt := promptui.Select{
		Label:     "Select Member",
		Items:     items,
		Size:      10,
		Templates: templates,
	}

	i, _, err := prompt.Run()
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

	selectedRoom := (*liveRooms)[i]

	return selectedRoom
}

func promptSelectQuality(url string) string {
	streamKey, err := utils.GetStreamKey(url)
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

	template := &promptui.SelectTemplates{
		Label:    "{{ . | cyan }}?",
		Active:   "\U0001F449 {{ . | green | underline }}",
		Inactive: "  {{ . | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select Quality",
		Items:     []string{"Original", "Low Latency"},
		Templates: template,
	}

	i, _, err := prompt.Run()
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

	if i == 0 {
		return fmt.Sprintf("https://hls-ull.showroom-cdn.com/%s/source/chunklist.m3u8", streamKey)
	} else {
		return fmt.Sprintf("https://hls-origin249.showroom-cdn.com/liveedge/%s_low/chunklist.m3u8", streamKey)
	}
}

func startVLCService(url string) {
	cmd := exec.Command("vlc", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}
}

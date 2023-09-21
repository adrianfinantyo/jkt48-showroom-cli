package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/adrianfinantyo/jkt48-showroom-cli/models"
	"github.com/adrianfinantyo/jkt48-showroom-cli/utils"
	"github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var watchLive = &cobra.Command{
	Use:   "watch",
	Short: "Watch JKT48 member live stream",
	Run: func(cmd *cobra.Command, args []string) {
		watchLiveRun(cmd)
	},
}

func watchLiveRun(cmd *cobra.Command) {
	memberChan := make(chan *[]models.Room)

	utils.LogInfo("💫 step 1 of 2 | Getting all JKT48 members live status...")
	fmt.Println()
	progressBar1 := progressbar.NewOptions(len(utils.AddedRooms), progressbar.OptionSetWidth(35), progressbar.OptionOnCompletion(func() {
		fmt.Fprint(cmd.OutOrStdout(), "\n")
	}))

	go utils.GetAllJKT48Rooms(progressBar1, memberChan)
	member := <-memberChan

	fmt.Println()
	utils.LogInfo("💫 step 2 of 2 | Seaching their streaming url's")
	fmt.Println()
	progressBar2 := progressbar.NewOptions(len(*member), progressbar.OptionSetWidth(35), progressbar.OptionOnCompletion(func() {
		fmt.Fprint(cmd.OutOrStdout(), "\n")
	}))

	go func() {
		for i := 0; i < progressBar2.GetMax(); i++ {
			if i == progressBar2.GetMax()-1 && !progressBar2.IsFinished() {
				progressBar2.ChangeMax(progressBar2.GetMax() + 10)
			}
			progressBar2.Add(1)
			time.Sleep(1 * time.Second)
		}
	}()

	var liveMembers []models.LiveRoom
	for _, data := range *member {
		if data.IsLive {
			urlList := utils.GetActiveRoomsByMemberData(data.Id)
			liveMembers = append(liveMembers, models.LiveRoom{
				RoomId:        data.Id,
				RoomKey:       data.Name,
				StreamURLList: *urlList,
			})
		}
		progressBar2.Finish()
	}

	if len(liveMembers) == 0 {
		fmt.Println()
		utils.LogInfo("🤭 Oops, seems like there's no active rooms right now")
		os.Exit(0)
	}

	selectedRoom := promptGetSelectMember(&liveMembers)
	if selectedRoom.RoomId != 0 {
		streamURL := promptSelectQuality(&selectedRoom.StreamURLList)
		startVLCService(streamURL)
	} else {
		err := fmt.Errorf("⚠️ Error: Room ID is 0")
		utils.LogError(err)
		os.Exit(1)
	}
}

func init() {
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

func promptSelectQuality(urlList *[]models.StreamURLList) string {
	template := &promptui.SelectTemplates{
		Label:    "{{ . | cyan }}?",
		Active:   "\U0001F449 {{ . | green | underline }}",
		Inactive: "  {{ . | cyan }}",
	}

	var streamLabel []string
	for _, data := range *urlList {
		streamLabel = append(streamLabel, data.Label)
	}

	prompt := promptui.Select{
		Label:     "Select Quality",
		Items:     streamLabel,
		Templates: template,
	}

	i, _, err := prompt.Run()
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

	return (*urlList)[i].StreamURL
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

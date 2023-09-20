package main

import (
	"github.com/adrianfinantyo/jkt48-showroom-cli/cmd"
)

func main() {
	// member := utils.GetAllJKT48Rooms()
	// for _, data := range *member {
	// 	fmt.Println(data.Name)
	// }

	// liveRooms := utils.GetActiveRooms()
	// if len(*liveRooms) == 0 {
	// 	println("No active rooms")
	// } else {
	// 	for _, data := range *liveRooms {
	// 		println(data.RoomKey)
	// 	}
	// }

	cmd.Execute()
}

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/adrianfinantyo/jkt48-showroom-cli/models"
)

func GetAllJKT48Rooms() *[]models.Room {
	var jkt48Room []models.Room

	res, err := http.Get("https://campaign.showroom-live.com/akb48_sr/data/room_status_list.json")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var decodedData []models.Room
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedData); err != nil {
		fmt.Println(err)
	}

	for _, data := range decodedData {
		if strings.Contains(data.Name, "JKT48") {
			jkt48Room = append(jkt48Room, data)
		}
	}

	return &jkt48Room
}

func GetActiveRooms() *[]models.LiveRoom {
	var liveRooms []models.LiveRoom

	res, err := http.Get("https://www.showroom-live.com/api/live/onlives")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var decodedData interface{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedData); err != nil {
		fmt.Println(err)
	}

	for _, data := range decodedData.(map[string]interface{})["onlives"].([]interface{}) {
		if data.(map[string]interface{})["genre_name"] == "Idol" {
			for _, room := range data.(map[string]interface{})["lives"].([]interface{}) {
				roomKey := room.(map[string]interface{})["room_url_key"]
				if strings.Contains(roomKey.(string), "JKT48") {
					roomId := room.(map[string]interface{})["room_id"]
					streamURL := room.(map[string]interface{})["streaming_url_list"].([]interface{})[0].(map[string]interface{})["url"]
					liveRooms = append(liveRooms, models.LiveRoom{RoomId: int(roomId.(float64)), RoomKey: roomKey.(string), StreamURL: streamURL.(string)})
				}
			}
		}
	}

	return &liveRooms
}

func GetStreamKey(url string) (string, error) {
	print(url)
	pattern := `https://hls-origin\d+\.showroom-cdn\.com/liveedge/([A-Za-z0-9]+)_low/chunklist\.m3u8`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(url)
	if len(match) >= 2 {
		key := match[1]
		return key, nil
	} else {
		return "", fmt.Errorf("Cannot find stream key")
	}
}

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/adrianfinantyo/jkt48-showroom-cli/models"
	"github.com/schollz/progressbar/v3"
)

func GetAllJKT48Rooms(bar *progressbar.ProgressBar) *[]models.Room {
	var jkt48Room []models.Room

	res, err := http.Get(AKB48RoomURL)
	if err != nil {
		LogError(err)
	}
	defer res.Body.Close()

	var decodedData []models.Room
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedData); err != nil {
		LogError(err)
	}

	for _, data := range decodedData {
		if strings.Contains(data.Name, "JKT48") {
			jkt48Room = append(jkt48Room, data)
			bar.Add(1)
		}
	}

	for _, data := range AddedRooms {
		res, err := http.Get(fmt.Sprintf("%s/profile?room_id=%d", RoomApiURL, data.RoomId))
		if err != nil {
			LogError(err)
		}
		defer res.Body.Close()

		var decodedData interface{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&decodedData); err != nil {
			LogError(err)
		}

		newRoom := models.Room{
			Id:               data.RoomId,
			Name:             decodedData.(map[string]interface{})["room_name"].(string),
			URLKey:           decodedData.(map[string]interface{})["room_url_key"].(string),
			ImageURL:         decodedData.(map[string]interface{})["image"].(string),
			Description:      decodedData.(map[string]interface{})["description"].(string),
			FollowerNum:      int(decodedData.(map[string]interface{})["follower_num"].(float64)),
			IsLive:           decodedData.(map[string]interface{})["is_onlive"].(bool),
			IsParty:          decodedData.(map[string]interface{})["is_party_enabled"].(bool),
			NextLiveSchedule: 0,
		}

		jkt48Room = append(jkt48Room, newRoom)
		bar.Add(1)
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

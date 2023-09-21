package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/adrianfinantyo/jkt48-showroom-cli/models"
	"github.com/schollz/progressbar/v3"
)

func GetAllJKT48Rooms(bar *progressbar.ProgressBar, resultChan chan<- *[]models.Room) {
	var jkt48Room []models.Room

	res, err := http.Get(AKB48RoomURL)
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		LogError(fmt.Errorf("⚠️ Error: %s", res.Status))
		os.Exit(1)
	}

	var decodedData []models.Room
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedData); err != nil {
		LogError(err)
	}

	for _, data := range decodedData {
		if strings.Contains(data.Name, "JKT48") {
			jkt48Room = append(jkt48Room, data)
		}
	}

	bar.ChangeMax(bar.GetMax() + len(jkt48Room))
	bar.Add(len(jkt48Room))

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

	resultChan <- &jkt48Room
}

func GetActiveRoomsByMemberData(roomId int) *[]models.StreamURLList {
	var StreamURLList []models.StreamURLList

	res, err := http.Get(fmt.Sprintf("%s/streaming_url?room_id=%d", LiveApiURL, roomId))
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		LogError(fmt.Errorf("⚠️ Error: %s", res.Status))
		os.Exit(1)
	}

	var decodedData interface{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedData); err != nil {
		LogError(err)
	}

	for _, data := range decodedData.(map[string]interface{})["streaming_url_list"].([]interface{}) {
		StreamURLList = append(StreamURLList, models.StreamURLList{Label: data.(map[string]interface{})["label"].(string), StreamURL: data.(map[string]interface{})["url"].(string)})
	}

	return &StreamURLList
}

func GetStreamKey(url string) (*models.LiveStream, error) {
	print(url)
	pattern := `https://hls-origin\d+\.showroom-cdn\.com/liveedge/([A-Za-z0-9]+)_low/chunklist\.m3u8`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(url)
	if len(match) >= 2 {
		liveStreamData := models.LiveStream{
			// The unique id logic still not clear yet
			UniqueId: match[0],
			Key:      match[1],
		}
		return &liveStreamData, nil
	} else {
		return nil, fmt.Errorf("Cannot find stream key")
	}
}

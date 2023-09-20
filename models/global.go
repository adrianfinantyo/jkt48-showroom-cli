package models

type Room struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	URLKey           string `json:"url_key"`
	ImageURL         string `json:"image_url"`
	Description      string `json:"description"`
	FollowerNum      int    `json:"follower_num"`
	IsLive           bool   `json:"is_live"`
	IsParty          bool   `json:"is_party"`
	NextLiveSchedule int    `json:"next_live_schedule"`
}

type LiveRoom struct {
	RoomId    int
	RoomKey   string
	StreamURL string
}

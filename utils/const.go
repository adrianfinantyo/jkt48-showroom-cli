package utils

import "github.com/adrianfinantyo/jkt48-showroom-cli/models"

const (
	AppVersion     = "v0.1.0"
	AppName        = "JKT48 Showroom CLI"
	AppDescription = "JKT48 Showroom CLI is a CLI tool to get information about JKT48 members showrooms"
	AppAuthor      = "Adrian Finantyo"
	AppLicense     = "MIT"
	ShowroomApiURL = "https://www.showroom-live.com/api"
	RoomApiURL     = ShowroomApiURL + "/room"
	LiveApiURL     = ShowroomApiURL + "/live"
	AKB48RoomURL   = "https://campaign.showroom-live.com/akb48_sr/data/room_status_list.json"
)

var (
	AddedRooms = []models.CustomRoom{
		{Nick: "Amanda", RoomId: 400710},
		{Nick: "Lia", RoomId: 400713},
		{Nick: "Callie", RoomId: 400714},
		{Nick: "Ela", RoomId: 400715},
		{Nick: "Indira", RoomId: 400716},
		{Nick: "Lyn", RoomId: 400717},
		{Nick: "Raisha", RoomId: 400718},
		{Nick: "Alya", RoomId: 461451},
		{Nick: "Anin", RoomId: 461452},
		{Nick: "Cathy", RoomId: 461454},
		{Nick: "Chelsea", RoomId: 461458},
		{Nick: "Cynthia", RoomId: 461463},
		{Nick: "Elin", RoomId: 461475},
		{Nick: "Danella", RoomId: 461466},
		{Nick: "Daisy", RoomId: 461465},
		{Nick: "Gracie", RoomId: 461478},
		{Nick: "Greseel", RoomId: 461479},
		{Nick: "Gendis", RoomId: 461476},
		{Nick: "Jeane", RoomId: 461480},
		{Nick: "Michie", RoomId: 461481},
	}
)

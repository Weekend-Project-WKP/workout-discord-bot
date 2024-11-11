package models

type User struct {
	DiscordUserId int
	Username      string
	TeamId        int
}

type Team struct {
	Id       int
	TeamName string
}
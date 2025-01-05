package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	DiscordUserId int
	Username      string
	TeamId        int
}

type Team struct {
	Id       primitive.ObjectID	`bson:"_id,omitempty"`
	TeamName string 			`bson:"team_name"`
}
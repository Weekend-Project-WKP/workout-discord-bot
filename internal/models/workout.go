package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkoutCategory struct {
	Id                     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	CategoryName           string
	Points                 float64
	Measurement            float64
	MeasurementDescription string
}

type Workout struct {
	DiscordUserName   string
	DiscordGuildId    string
	WorkoutCategoryId primitive.ObjectID
	WorkoutEntryTime  primitive.DateTime
	MessageId         string
	Points            float64
	TeamName          string
	Description       string
}

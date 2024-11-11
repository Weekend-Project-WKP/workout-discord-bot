package models

type WorkoutCategory struct {
	Id                     int
	CategoryName           string
	Points                 int
	Measurement            int
	MeasurementDescription string
}

type Workout struct {
	Id                int
	DiscordUserId     int
	WorkoutCategory   int
	ActualWorkoutTime string
	WorkoutEntryTime  string
}
package models

type WorkoutCategory struct {
	_id                    string
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

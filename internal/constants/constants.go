package constants

const (
	// DB Constants
	DbName                    = "workoutbot"
	WorkoutCategoryCollection = "workout_categories"
	WorkoutsCollection        = "workouts"
	TeamsCollection           = "teams"
	// AI Constants
	AiPrompt = "Analyze a photo and/or message to group it the correct category with the correct duration or distance of the workout. Convert hours to minutes. Sometimes the format of time will be hours:minutes:seconds. Use the same sentence for each workout substituting XXX and YYY; \"Category='XXX' for Duration/Distance='YYY' miles/minutes\". XXX will be replaced with the category you determine. YYY will be replaced with the duration in miles or minutes. Add a newline between workouts. Make sure to add all the static text. It is also possible we won't send you an image and just will send you a sentence with what we did. If this is the case, we would also like that in your structured response. Start your message with: \"Workout Summary for 'USER_NAME' 'TEAM_NAME' on 'WORKOUT_DATETIME\". USER_NAME, TEAM_NAME, and WORKOUT_DATETIME will be replaced with the user name we gave you. Make sure the apostrophes are in your message. The total response should be less than 2000 characters. High Intensity Interval Training (HIIT) is considered strength. Yoga is considered stretching. Peloton is considered biking, only use the miles, do not track the time. "
	// TODO: Update this logic to reference the user by tag≈
	AiErrorMessage = "There was an issue with Google Geimini initializing. Use the /workout slash command to log your workout (keep AI emoji on picture so we know it failed)"
	// Application Constants
	PeriodStartDay      = "monday"
	PeriodEndDay        = "sunday"
	Prefix              = "!workoutbot"
	DavidGoginsAiPrompt = "Make fun of the workout or message in a way david goggins would."
	WorkoutGangServerId = "898225376642338846"
)

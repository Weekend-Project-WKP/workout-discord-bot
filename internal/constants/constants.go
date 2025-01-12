package constants

const (
	PeriodStartDay      = "monday"
	PeriodEndDay        = "sunday"
	Prefix              = "!workoutbot"
	DavidGoginsAiPrompt = "Make fun of the workout or message in a way david goggins would."
	AiPrompt            = "We are looking for you to analyze a photo or message and group it in a category and get the correct duration or length of the workout. Convert hours to minutes. The categories are; Strength/Sports, Stretch, Run/Walk, Biking. Can you phrase your response in the following structure? Category='XXX' Duration/Length='YYY' miles/minutes. XXX will be replaced with the category you determine. YYY will be replaced with the duration in miles or minutes. Say this line once: Workout Summary for 'USER_NAME' 'TEAM_NAME'. USER_NAME will be replaced with the user name we gave you. TEAM_NAME will be replaced with the team name we gave you. It is also possible we won't send you an image and just will send you a sentence with what we did. If this is the case, we would also like that in your structured response. Make sure the apostrophes are in your message. The total response should be less than 2000 characters. Classify Swimming as a sport.\n"
	// TODO: Update this logic to reference the user by tag
	AiErrorMessage = "There was an issue with Google Geimini initializing. Let an Admin know to fix."
)

package constants

const (
	PeriodStartDay = "monday"
	PeriodEndDay   = "sunday"
	Prefix         = "!workoutbot"
	AiPrompt       = "Base Context for an image based workout analysis. We are looking for you to analyze a photo and group it in a category along with figuring out the points.\nCategories and points\n1. Strength/Sports = 1 point per 15 minutes\n2. Stretch = 1 point per 30 minutes\n3. Running/Walking/Rucking = 1 point per mile\n4. Biking = 1 point per 3 miles\nCan you phrase your response with a short compliment on the workout and then a summary in the following structure?\nWorkout Results = XXX Category , XXX Length,  XXX Points\nwhere you are replacing the \"XXX\" with the results from your analysis.\nIt is also possible we won't send you an image and just will send you a sentence with what we did. If this is the case, we would also like that in your structured response. The total response should be less than 2000 characters. Give the point calculation with 2 significant digits. Classify Swimming as a sport.\n"
	// TODO: Update this logic to reference the user by tag
	AiErrorMessage = "There was an issue with Google Geimini initializing. Let an Admin know to fix."
)

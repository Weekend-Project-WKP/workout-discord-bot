package helpers

import (
	"strconv"
	"strings"
	"time"
	"workoutbot/internal/db"
	"workoutbot/internal/models"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Make this string structure not as hard coded dependent
/* String Structure Expected - Index 1 gives username and category name, index 3 gives team name and duration/length depending on if its first row or not
Workout Summary for 'User_name' 'Team 1'
Category="Run/Walk" Duration/Length="1" mile
Category="Sports" Duration/Length="15" minutes
*/
func CreateWorkoutsViaString(workoutString string, guildId string, messageId string, messageTs time.Time) ([]models.Workout, error) {
	workoutCategoryMap, err := db.WorkoutCategoryGetAll()
	var workouts []models.Workout
	if err == nil {
		username, teamname := "", ""
		for i, line := range strings.Split(workoutString, "\n") {
			// Blank Line skipping.
			// TODO - Figure out how to remove all the blank lines from AI
			if line != "" {
				lineItemSplit := strings.Split(line, "'")
				if strings.Contains(line, "Workout Summary") {
					// First Line for Info
					username = lineItemSplit[1]
					teamname = lineItemSplit[3]
				} else if strings.Contains(line, "Category") {
					// Following Lines for Workouts
					durationInt, _ := strconv.ParseFloat(lineItemSplit[3], 64)
					workouts = append(workouts, models.Workout{
						Points:            CalculatePoints(workoutCategoryMap[lineItemSplit[1]].Points, workoutCategoryMap[lineItemSplit[1]].Measurement, durationInt),
						DiscordUserName:   username,
						DiscordGuildId:    guildId,
						Description:       line,
						WorkoutCategoryId: workoutCategoryMap[lineItemSplit[1]].Id,
						MessageId:         messageId + "-" + strconv.Itoa(i),
						WorkoutEntryTime:  primitive.NewDateTimeFromTime(messageTs.UTC()),
						TeamName:          teamname})
				}
			}
		}
	}
	return workouts, err
}

func LogWorkouts(s *discordgo.Session, workouts []models.Workout, channelId string, teamName string) {
	_, err := db.WorkoutsInsertMany(workouts)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error collection") {
			s.ChannelMessageSend(channelId, "This workout was already logged, stop trying to cheat.")
		} else {
			s.ChannelMessageSend(channelId, "There was an issue logging this workout. Sounds like no points for you. Contact an admin.")
		}
	} else {
		runningPoints := 0.00
		for _, workout := range workouts {
			runningPoints += workout.Points
		}
		s.ChannelMessageSend(channelId, strconv.FormatFloat(runningPoints, 'f', -1, 64)+" points awarded to "+teamName)
	}
}

func CalculatePoints(categoryPoint float64, categoryPointInterval float64, workoutLength float64) float64 {
	return float64((workoutLength / categoryPointInterval) * categoryPoint)
}
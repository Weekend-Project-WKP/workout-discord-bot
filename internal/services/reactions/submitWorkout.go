package reactions

import (
	"fmt"
	"strconv"
	"strings"
	"workoutbot/internal/db"
	"workoutbot/internal/helpers"
	"workoutbot/internal/models"

	"github.com/bwmarrin/discordgo"
)

// Workout Summary for 'User_name' 'Team 1' index =1,3 size 5
// Category="Run/Walk" Duration/Length="1" mile index 1,3 size 5
// Category="Sports" Duration/Length="15" minutes
func SubmitWorkout(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	workoutCategoryMap, err := db.WorkoutCategoryGetAll()
	if err != nil {
		fmt.Println(err)
		s.ChannelMessageSend(r.ChannelID, "No Workout categories. Can't log anything. Sounds like no points for you. Contact an admin.")
	}
	message, _ := s.ChannelMessage(r.ChannelID, r.MessageID)
	username, teamname := "", ""
	runningPoints := 0.00
	var workouts []models.Workout
	// messageSplit := strings.Split(message.Content, "\n")
	for i, line := range strings.Split(message.Content, "\n") {
		// Blank Line skipping. TODO - Figure out how to remove all the blank lines from AI
		if line != "" {
			lineItemSplit := strings.Split(line, "'")
			if strings.Contains(line, "Workout Summary") {
				// First Line for Info
				username = lineItemSplit[1]
				teamname = lineItemSplit[3]
			} else if strings.Contains(line, "Category") {
				// Following Lines for Workouts
				durationInt, _ := strconv.Atoi(lineItemSplit[3])
				runningPoints += helpers.CalculatePoints(workoutCategoryMap[lineItemSplit[1]].Points, workoutCategoryMap[lineItemSplit[1]].Measurement, durationInt)
				workouts = append(workouts, models.Workout{
					Points:            helpers.CalculatePoints(workoutCategoryMap[lineItemSplit[1]].Points, workoutCategoryMap[lineItemSplit[1]].Measurement, durationInt),
					DiscordUserName:   username,
					DiscordGuildId:    r.GuildID,
					WorkoutCategoryId: workoutCategoryMap[lineItemSplit[1]].Id,
					MessageId:         r.MessageID + "-" + strconv.Itoa(i),
					WorkoutEntryTime:  message.Timestamp.UTC().String(),
					TeamName:          teamname})
			}
		}
	}
	if len(workouts) > 0 {
		err := helpers.CreateWorkouts(workouts)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key error collection") {
				s.ChannelMessageSend(r.ChannelID, "This workout was already logged, stop trying to cheat.")
			} else {
				s.ChannelMessageSend(r.ChannelID, "There was an issue logging this workout. Sounds like no points for you. Contact an admin.")
			}
		} else {
			s.ChannelMessageSend(r.ChannelID, strconv.FormatFloat(runningPoints, 'f', -1, 64)+" points awarded to "+teamname)
		}
	}
}

package reactions

import (
	"fmt"
	"workoutbot/internal/helpers"

	"github.com/bwmarrin/discordgo"
)

func SubmitWorkout(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	message, _ := s.ChannelMessage(r.ChannelID, r.MessageID)
	workouts, dbErr := helpers.CreateWorkoutsViaString(message.Content, r.GuildID, r.MessageID, message.Timestamp)
	if dbErr != nil {
		fmt.Println(dbErr)
		s.ChannelMessageSend(r.ChannelID, dbErr.Error())
		return
	} else if len(workouts) > 0 {
		helpers.LogWorkouts(s, workouts, r.ChannelID, workouts[0].TeamName)
	}
}

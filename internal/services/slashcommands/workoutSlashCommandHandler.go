package slashcommands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func WorkoutSlashCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate){
	//name := i.ApplicationCommandData().Options[0].StringValue()

	// Respond with a summary of the points added
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Workout Added!",
		},
	})
	
	if err != nil {
		log.Printf("Error responding to slash command: %v", err)
	}
}
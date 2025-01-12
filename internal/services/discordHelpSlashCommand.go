package services

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "hello" {
		name := i.ApplicationCommandData().Options[0].StringValue()

		// Respond with a greeting
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Hello, %s!", name),
			},
		})

		if err != nil {
			log.Printf("Error responding to slash command: %v", err)
		}
	}
}

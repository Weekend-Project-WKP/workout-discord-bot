package services

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func DiscordHelpSlashCommandHandler(session *discordgo.Session) {
	// Register the slash command
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			handleCommand(s, i)
		}
	})
}

func handleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

		// Register the command with Discord
		command := &discordgo.ApplicationCommand{
			Name:        "hello",
			Description: "Sends a greeting",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "name",
					Description: "Your name",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
			},
		}

		_, errDiscord := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if errDiscord != nil {
			log.Printf("Failed to load command: %v", errDiscord)
		}
		fmt.Println(errDiscord)
	}
}

package slashcommands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func WorkoutSlashCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate){
	//name := i.ApplicationCommandData().Options[0].StringValue()

		// Respond with a greeting
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Workout Added!",
			},
		})

		if err != nil {
			log.Printf("Error responding to slash command: %v", err)
		}

		// // Register the command with Discord
		// command := &discordgo.ApplicationCommand{
		// 	Name:        "workout",
		// 	Description: "Adds a new workout for the current day",
		// 	Options: []*discordgo.ApplicationCommandOption{
		// 		{
		// 			Name:        "Workout Category",
		// 			Description: "Type of workout completed",
		// 			Type:        discordgo.ApplicationCommandOptionString,
		// 			Required:    true,
		// 		},
		// 	},
		// }

		// _, errDiscord := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		// if errDiscord != nil {
		// 	log.Printf("Failed to load command: %v", errDiscord)
		// }
		// fmt.Println(errDiscord)
}
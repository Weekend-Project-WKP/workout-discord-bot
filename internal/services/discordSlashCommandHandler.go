package services

import (
	"fmt"
	"log"
	"workoutbot/internal/services/slashcommands"

	"github.com/bwmarrin/discordgo"
)

func DiscordSlashCommandHandler(session *discordgo.Session) {
	// Register the slash command
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handleCommand(s, i)

		// Define multiple slash commands
		commands := []*discordgo.ApplicationCommand{
			{
				Name:        "hello",
				Description: "Sends a greeting",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "name",
						Description: "Your name",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
					{
						Name:        "age",
						Description: "Your age",
						Type:        discordgo.ApplicationCommandOptionInteger,
						Required:    true,
					},
				},
			},
			{
				Name:        "workout",
				Description: "Adds a new workout for the current day",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "workout-category",
						Description: "Type of workout completed",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
					{
						Name:        "workout-duration-distance",
						Description: "Distance or Duration of the workout",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
				},
			},
		}

		// Register the commands
		for _, cmd := range commands {
			log.Printf("Command Name: %v", cmd.Name)
			_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
			if err != nil {
				log.Printf("Cannot create '%v' command: %v", cmd.Name, err)
			}
		}
	})
}

func handleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	
	switch i.ApplicationCommandData().Name{
		case "hello":
			log.Printf("Enter hello slash command")
			slashcommands.HelloSlashCommandHandler(s, i)
		case "workout":
			log.Printf("Enter workout slash command")
			slashcommands.WorkoutSlashCommandHandler(s, i)
		case "goodbye":
			log.Printf("Enter goodbye slash command")
			slashcommands.WorkoutSlashCommandHandler(s, i)
		default:
			fmt.Printf("No slash command found for %v\n", i.ApplicationCommandData().Name)
		}	
}

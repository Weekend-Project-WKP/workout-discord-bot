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
		// if i.Type == discordgo.InteractionApplicationCommand {
		// 	handleCommand(s, i)
		// }
		
		// Register the command with Discord
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

			// Define multiple slash commands
		// commands := []*discordgo.ApplicationCommand{
		// 	{
		// 		Name:        "hello",
		// 		Description: "Says hello!",
		// 	},
		// 	{
		// 		Name:        "goodbye",
		// 		Description: "Says goodbye!",
		// 	},
		// }

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
						Name:        "Workout Category",
						Description: "Type of workout completed",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
					{
						Name:        "Workout Duration/Distance",
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
	// if i.ApplicationCommandData().Name == "hello" {
	// 	name := i.ApplicationCommandData().Options[0].StringValue()

	// 	// Respond with a greeting
	// 	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
	// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 		Data: &discordgo.InteractionResponseData{
	// 			Content: fmt.Sprintf("Hello, %s!", name),
	// 		},
	// 	})

	// 	if err != nil {
	// 		log.Printf("Error responding to slash command: %v", err)
	// 	}

	// 	// Register the command with Discord
	// 	command := &discordgo.ApplicationCommand{
	// 		Name:        "hello",
	// 		Description: "Sends a greeting",
	// 		Options: []*discordgo.ApplicationCommandOption{
	// 			{
	// 				Name:        "name",
	// 				Description: "Your name",
	// 				Type:        discordgo.ApplicationCommandOptionString,
	// 				Required:    true,
	// 			},
	// 		},
	// 	}

	// 	_, errDiscord := s.ApplicationCommandCreate(s.State.User.ID, "", command)
	// 	if errDiscord != nil {
	// 		log.Printf("Failed to load command: %v", errDiscord)
	// 	}
	// 	fmt.Println(errDiscord)
	// }
}

package services

import (
	"fmt"
	"log"
	"workoutbot/internal/constants"
	"workoutbot/internal/db"
	"workoutbot/internal/services/slashcommands"

	"github.com/bwmarrin/discordgo"
)

func DiscordSlashCommandHandler() {
	// Register the slash command
	DiscordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		defer FatalSessionClosing()
		handleCommand(s, i)

		workoutCategories, err := db.WorkoutCategoryGetAll()
		if err != nil {
			log.Printf("Failed to get workout category: %v", err)
			return
		}

		// Get all current discord users
		// Fetch members in the guild
		members, err := s.GuildMembers(constants.WorkoutGangServerId, "", 1000) // Fetch up to 1000 members
		if err != nil {
			log.Printf("Error fetching guild members: %v", err)
		}

		// Filter out bots
		var userList []string
		for _, member := range members {
			if !member.User.Bot { // Exclude bots
				userList = append(userList, member.User.Username)
			}
		}

		// Define multiple slash commands
		commands := []*discordgo.ApplicationCommand{
			{
				Name:        "workout",
				Description: "Adds a new workout for the current day",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "workout-category",
						Description: "Type of workout completed",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
						Choices: func() []*discordgo.ApplicationCommandOptionChoice {
							var choices []*discordgo.ApplicationCommandOptionChoice
							for _, category := range workoutCategories {
								choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
									Name:  category.CategoryName,
									Value: category.CategoryName, // The value could differ if needed, e.g., an ID or slug
								})
							}
							return choices
						}(),
					},
					{
						Name:        "workout-duration-distance",
						Description: "Distance or Duration of the workout",
						Type:        discordgo.ApplicationCommandOptionString, // TODO: Use discordgo.ApplicationCommandOptionNumber so it forces a number
						Required:    true,
					},
					{
						Name: "user",
						Description: "User who completed the workout",
						Type: discordgo.ApplicationCommandOptionString,
						Required: false,
						Choices: func() []*discordgo.ApplicationCommandOptionChoice {
							var choices []*discordgo.ApplicationCommandOptionChoice
							for _, user := range userList {
								choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
									Name:  user,
									Value: user, // The value could differ if needed, e.g., an ID or slug
								})
							}
							return choices
						}(),
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
	switch i.ApplicationCommandData().Name {
	case "workout":
		log.Printf("Enter workout slash command")
		slashcommands.WorkoutSlashCommandHandler(s, i)
	default:
		fmt.Printf("No slash command found for %v\n", i.ApplicationCommandData().Name)
	}
}

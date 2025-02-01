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
	})
}

func RegisterCommands() error {
	// Fetch workout categories
	workoutCategories, err := db.WorkoutCategoryGetAll()
	if err != nil {
		return fmt.Errorf("failed to get workout category: %v", err)
	}

	// Fetch members in the guild
	members, err := DiscordSession.GuildMembers(constants.WorkoutGangServerId, "", 1000)
	if err != nil {
		return fmt.Errorf("error fetching guild members: %v", err)
	}

	// Filter out bots
	var userList []string
	for _, member := range members {
		if !member.User.Bot {
			userList = append(userList, member.User.Username)
		}
	}

	// Define commands
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
								Value: category.CategoryName,
							})
						}
						return choices
					}(),
				},
				{
					Name:        "workout-duration-distance",
					Description: "Distance or Duration of the workout",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
				},
				{
					Name:        "user",
					Description: "User who completed the workout",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
					Choices: func() []*discordgo.ApplicationCommandOptionChoice {
						var choices []*discordgo.ApplicationCommandOptionChoice
						for _, user := range userList {
							choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
								Name:  user,
								Value: user,
							})
						}
						return choices
					}(),
				},
			},
		},
	}

	// TODO: Might want to clear the commands out before the start of every bot to make sure they are updated here

	// Register commands (use a guild for instant updates)
	for _, cmd := range commands {
		_, err := DiscordSession.ApplicationCommandCreate(DiscordSession.State.User.ID, constants.WorkoutGangServerId, cmd) // Register as guild command
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", cmd.Name, err)
		}
	}

	log.Println("Commands registered successfully")
	return nil
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



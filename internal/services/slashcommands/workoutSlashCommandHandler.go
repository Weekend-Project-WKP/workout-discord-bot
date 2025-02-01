package slashcommands

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"workoutbot/internal/constants"
	"workoutbot/internal/db"
	"workoutbot/internal/helpers"
	"workoutbot/internal/models"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WorkoutSlashCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate){
	data := i.ApplicationCommandData()
	
	// Find the category
	category := FindOption(data.Options, "workout-category").StringValue()
	log.Printf("Slash Command Category: '%v'", category)

	// Find the measurement
	measurement, errMeasurment := strconv.ParseFloat(FindOption(data.Options, "workout-duration-distance").StringValue(), 64)
	if errMeasurment != nil {
		fmt.Println("Error converting measurement string to float:", errMeasurment)
		return
	}
	log.Printf("Slash Command Measurement: '%v'", measurement)

	var username string
	// Find the user
	choiceUser := FindOption(data.Options, "user")
	if choiceUser == nil {
		username = i.Member.User.Username
	} else {
		username = choiceUser.StringValue()
	}
	log.Printf("Slash Command Username: '%v'", username)

	// Get the Workout Category from the DB using the slash command option
	//category := i.ApplicationCommandData().Options[0].StringValue()
	workoutCategory, errWorkoutCategory:= db.WorkoutCategoryGetOne(category)
	if errWorkoutCategory != nil {
		log.Printf("Error gathering workout category '%v' 'WorkoutSlashCommandHandler' command: %v", category, errWorkoutCategory)
	}

	// Get the measurement from the 2nd slash command option
	//measurement, errMeasurment := strconv.ParseFloat(i.ApplicationCommandData().Options[1].StringValue(), 64) //TODO: Update to pull from a Number Value once the SlashComandHandler is updated
	

	// Get the user from the optional 3rd slash command option
	//commandUser, errCmdUser := i.ApplicationCommandData().Options[2].StringValue()

	log.Printf("Interaction Guild Id: '%v'", i.GuildID)
	log.Printf("Interaction Message Author Id: '%v'", i.Member.User.ID)

	userId, errUserId := GetUserIDByUsername(s, constants.WorkoutGangServerId, username)
	if errUserId != nil {
		log.Println("User not found:", errUserId)
		return
	} else {
		fmt.Println("User ID:", userId)
	}
	
	// Get the user team name which is set to the current user role (guild id)
	teamName, errTeamName := helpers.GetTeamName(s, i.GuildID, userId)
	if errTeamName != nil {
		s.ChannelMessageSend(i.ChannelID, "This user isn't assigned to a team. Need a team my guy/gal/they.")
	}

	// username := i.Member.User.Username
	points := helpers.CalculatePoints(workoutCategory.Points, workoutCategory.Measurement, measurement)

	description := fmt.Sprintf("User `%v` logged a workout \n Category='%v' \n Duration/Length='%v' \n Points: '%v' \n Team='%v'", username, category, measurement, points, teamName)

	// Respond with a summary of the points added
	errInteraction := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%v added a workout!", username),
		},
	})

	if errInteraction != nil {
		log.Printf("Error responding to slash command: %v", errInteraction)
	}

	// Fetch the message ID from the follow-up response
	followUpMessage, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: fmt.Sprintf("Workout Summary \n %v", description),
	})
	if err != nil {
		fmt.Println("Error creating follow-up message:", err)
		s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("'%v' There was an issue logging your workout. Please contact an admin for assistance", username))
		return
	}
	
	// Create a new team document
	workout := models.Workout{
		DiscordUserName: username,	
		DiscordGuildId: i.GuildID,
		WorkoutCategoryId: workoutCategory.Id,
		WorkoutEntryTime: primitive.NewDateTimeFromTime(time.Now().UTC()),
		MessageId: followUpMessage.ID,
		Points: points,
		TeamName: teamName,
		Description: description,
	}
	var workouts []models.Workout
	workouts = append(workouts, workout)

	// Log the workout to the DB
	//helpers.LogWorkouts(s, workouts, i.ChannelID, teamName)

	s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("Workout logged for User '%v' on Team '%v' \nStart a thread on your request and contact an admin if you need to adjust your workout", username, teamName))
}

// FindOption searches for an option by name in a slice of options.
func FindOption(options []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, opt := range options {
		if opt.Name == name {
			return opt
		}
	}
	return nil
}

// GetUserIDByUsername searches for a user by their username in a guild
func GetUserIDByUsername(s *discordgo.Session, guildID, username string) (string, error) {
	members, err := s.GuildMembersSearch(guildID, username, 100) // Max 100 results
	if err != nil {
		return "", err
	}

	// Loop through results and match exact username
	for _, member := range members {
		if member.User.Username == username {
			return member.User.ID, nil
		}
	}
	return "", fmt.Errorf("user not found")
}
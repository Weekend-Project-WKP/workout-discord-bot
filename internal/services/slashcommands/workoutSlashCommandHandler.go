package slashcommands

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"workoutbot/internal/db"
	"workoutbot/internal/helpers"
	"workoutbot/internal/models"
	"workoutbot/internal/services/reactions"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func WorkoutSlashCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate){
	// Get the Workout Category from the DB using the slash command option
	category := i.ApplicationCommandData().Options[0].StringValue()
	workoutCategory, errWorkoutCategory:= db.WorkoutCategoryGetOne(category)
	if errWorkoutCategory != nil {
		log.Printf("Error gathering workout category '%v' 'WorkoutSlashCommandHandler' command: %v", category, errWorkoutCategory)
	}

	// Get the measurement from the 2nd slash command option
	measurement, errMeasurment := strconv.ParseFloat(i.ApplicationCommandData().Options[1].StringValue(), 64)
	if errMeasurment != nil {
		fmt.Println("Error converting measurement string to float:", errMeasurment)
		return
	}

	log.Printf("Interaction Guild Id: '%v'", i.GuildID)
	log.Printf("Interaction Message Author Id: '%v'", i.Member.User.ID)
	
	// Get the user team name which is set to the current user role (guild id)
	teamName, errTeamName := reactions.GetTeamName(s, i.GuildID, i.Member.User.ID)
	if errTeamName != nil {
		fmt.Println(errTeamName)
		s.ChannelMessageSend(i.ChannelID, "This user isn't assigned to a team. Need a team my guy/gal/they.")
	}

	username := i.Member.User.Username
	points := helpers.CalculatePoints(workoutCategory.Points, workoutCategory.Measurement, measurement)

	description := fmt.Sprintf("Category='%v' for Duration/Length='%v' %v", category, measurement, workoutCategory.MeasurementDescription)

	// Respond with a summary of the points added
	errInteraction := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Workout Added! %v", description),
		},
	})

	// Fetch the message ID from the follow-up response
	followUpMessage, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: "Follow-up message to get the ID.",
	})
	if err != nil {
		fmt.Println("Error creating follow-up message:", err)
		return
	}

	// Print the message ID of the follow-up message
	fmt.Printf("Message ID of the follow-up: %s\n", followUpMessage.ID)
	
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

	s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("-----Workout Debug----- UserName: %v DiscordGuildId: %v WorkoutCategoryId: %v WorkoutEntryTime: %v MessageId: %v Points: %v TeamName: %v Description: %v", 
	workout.DiscordUserName, workout.DiscordGuildId, workout.WorkoutCategoryId, workout.WorkoutEntryTime, workout.MessageId, workout.Points, workout.TeamName, workout.Description))

	if errInteraction != nil {
		log.Printf("Error responding to slash command: %v", errInteraction)
	}
}
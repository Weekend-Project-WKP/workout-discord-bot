package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"workoutbot/internal/db"
	"workoutbot/internal/services"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Project init
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get the Discord token
	// Initialize the session and log fatal if error
	session, errDiscord := services.InitializeDiscordGo()
	// Initiailze AI - Currently Gemini 1.5
	model, ctx, client, errAi := services.InitializeAiPartner()

	if errDiscord != nil {
		log.Printf("Error initializing DiscordGo: %v", errDiscord)
	}

	// Initialize MongoDB
	db.MongoDBInit()
	defer func() {
		if err := db.MongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
		log.Println("Disconnected from MongoDB")
	}()


	// TODO: Remove example DB code when we implement the examples in one of the discord handlers
	// team, err := db.TeamsGetOne("Test Team")
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// log.Println(team.Id, team.TeamName)

	// db.TeamsGetAll()
	
	// id, err := db.TeamsSaveOne("Saved Team")
	// if err != nil{
	// 	log.Println(err)
	// }
	// log.Printf("Saved teams ID: %v", id)

	// Add Discord handlers
	services.DiscordAddReactionHandler(session, model, ctx, errAi)
	services.DiscordRemoveReactionHandler(session)
	services.DiscordMessageCreateHandler(session)
	services.DiscordHelpMessageHandler(session)

	// Set the intent
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Open the Session
	errDiscord = session.Open()
	if errDiscord != nil {
		log.Fatal((errDiscord))
	}
	defer session.Close()
	defer client.Close()
	fmt.Println("The bot is online!")

	// Create a channel to listen to system notifications in order to close up. Use CTRL + C to close
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

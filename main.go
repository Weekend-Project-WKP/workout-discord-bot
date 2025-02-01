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
	errDiscord := services.InitializeDiscordGo()
	// Initiailze AI - Currently Gemini 1.5
	services.InitializeAiPartner()

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

	// Add Discord handlers
	services.DiscordAddReactionHandler()
	services.DiscordRemoveReactionHandler()
	services.DiscordMessageCreateHandler()
	services.DiscordHelpMessageHandler()

	// Set the intent
	services.DiscordSession.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Open the Session
	errDiscord = services.DiscordSession.Open()
	if errDiscord != nil {
		log.Fatal((errDiscord))
	}
	defer services.DiscordSession.Close()
	defer services.AiClient.Close()
	defer func() {
		services.DiscordSession.ChannelMessageSend("898225507064250420", "The Bot is Offline from Main.go Ending")
	}()

	// This needs to happen after the session.Open() in order for commands to be registered.
	services.RegisterCommands()
	services.DiscordSlashCommandHandler()

	services.DiscordSession.ChannelMessageSend("898225507064250420", "The bot is online!")
	fmt.Println("The bot is online!")

	// Create a channel to listen to system notifications in order to close up. Use CTRL + C to close
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"workoutbot/internal/services"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Project init
	services.Hello()

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
	if errAi != nil {
		log.Fatal(errAi)
	}

	// Add Discord handlers
	services.DiscordAddReactionHandler(session, model, ctx)
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

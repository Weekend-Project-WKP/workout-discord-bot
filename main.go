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
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("DISCORD_TOKEN not set in environment")
	}		

	// Initialize the session and log fatal if error
	session, err := discordgo.New(discordToken)
	if err != nil {
		log.Fatal((err))
	}

	// Process incomming message request
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate){
		// If author of message is the same as the session author ignore the message
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Look for the word "hello" and reply "world!" in chat
		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
		}
	})

	// Set the intent
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Open the Session
	err = session.Open()
	if err != nil {
		log.Fatal((err))
	}
	defer session.Close()

	fmt.Println("The bot is online!")

	// Create a channel to listen to system notifications in order to close up. Use CTRL + C to close
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
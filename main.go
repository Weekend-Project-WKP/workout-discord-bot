package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"workoutbot/internal/constants"
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

	session.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd){
		fmt.Printf("%v reacted with %v\n", r.UserID, r.Emoji.Name)

		if r.Emoji.Name == "💪🏿" {
			s.GuildMemberRoleAdd(r.GuildID, r.UserID, "1311487436278337638")
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v has been added to the %v", r.UserID, "Workout Challenge role" ))
		}
	})

	session.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionRemove){
		fmt.Printf("%v removed reaction %v\n", r.UserID, r.Emoji.Name)

		if r.Emoji.Name == "💪🏿" {
			s.GuildMemberRoleRemove(r.GuildID, r.UserID, "1311487436278337638")
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v has been removed from %v", r.UserID, "Workout Challenge role" ))
		}
	})

	// Process incomming message request
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate){
		// If author of message is the same as the session author ignore the message
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Parse the incomming message and look for the Prefix
		args := strings.Split(m.Content, " ")

		// If prefix not found ignore the messageq
		if args[0] != constants.Prefix {
			return
		}

		// If the sub command after the prefix is "hello" respond "world!""
		if args[1] == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
		}

		// If the sub command after the prefix is "hello" respond "world!""
		if args[1] == "proverbs" {
			proverbs := []string{
				"Don't communicate by sharing memory, share memory by communicating.",
				"Concurrency is not parallelism.",
				"Channels orchestrate; mutexes serialize.",
				"The bigger the interface, the weaker the abstraction.",
				"Make the zero value useful.",
				"interface{} says nothing.",
				"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
				"A little copying is better than a little dependency.",
				"Syscall must always be guarded with build tags.",
				"Cgo must always be guarded with build tags.",
				"Cgo is not Go.",
				"With the unsafe package there are no guarantees.",
				"Clear is better than clever.",
				"Reflection is never clear.",
				"Errors are values.",
				"Don't just check errors, handle them gracefully.",
				"Design the architecture, name the components, document the details.",
				"Documentation is for users.",
				"Don't panic.",
			}

			// Get a random proverb index
			selection := rand.Intn(len(proverbs))

			// Create embed message author
			author := discordgo.MessageEmbedAuthor{
				Name: "Rob Pike",
				URL: "https://go-proverbs.github.io/",
			}

			// Create embed message content
			embed := discordgo.MessageEmbed{
				Title: proverbs[selection],
				Author: &author,
			}

			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			//s.ChannelMessageSend(m.ChannelID, proverbs[selection])
		}		

		// If the sub command is "chill" respond "cousin" 
		if args[1] == "chill" {
			s.ChannelMessageSend(m.ChannelID, "cousin!")
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
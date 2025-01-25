package services

import (
	"math/rand"
	"strings"
	"workoutbot/internal/constants"

	"github.com/bwmarrin/discordgo"
)

func DiscordMessageCreateHandler() {
	// Process incomming message request
	DiscordSession.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		defer FatalSessionClosing()
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
				URL:  "https://go-proverbs.github.io/",
			}

			// Create embed message content
			embed := discordgo.MessageEmbed{
				Title:  proverbs[selection],
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
}

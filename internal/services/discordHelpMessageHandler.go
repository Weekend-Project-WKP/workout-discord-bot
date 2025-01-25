package services

import (
	"strings"
	"workoutbot/internal/constants"

	"github.com/bwmarrin/discordgo"
)

func DiscordHelpMessageHandler() {
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
		if args[1] == "help" {

			// Show a list of all commands and example options
			// TODO: Pull this list from a constants file instead of inline here
			commands := []string{
				"list <date> \n  ex: \n   !workoutbot list 12/21/2024, (Leaving date blank defaults to today)",
				"add <workoutcategory> <duration> <date> \n ex: \n !workoutbot add strength 1hr5m 12/21/2024 (Leaving date blank defaults to today)",
			}

			title := ""
			for i := 0; i <= len(commands)-1; i++ {
				title += commands[i] + "\n \n"
			}

			// Create embed message author
			author := discordgo.MessageEmbedAuthor{
				Name: "Commands Help Page (Click Me for more info)",
				URL:  "https://github.com/Weekend-Project-WKP/workout-discord-bot/blob/main/README.md#commands",
			}

			// Create embed message content
			embed := discordgo.MessageEmbed{
				Title:  title,
				Author: &author,
				Color:  0xff5733,
			}
			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		}
	})
}

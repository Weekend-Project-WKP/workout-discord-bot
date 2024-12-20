package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func DiscordAddReactionHandler(session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd){
		fmt.Printf("%v reacted with %v\n", r.UserID, r.Emoji.Name)

		// TODO: Currently this only works for this specific muscle tone. We need to make this work for all muscle tones
		if r.Emoji.Name == "💪🏿" {
			s.GuildMemberRoleAdd(r.GuildID, r.UserID, "1311487436278337638")
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v has been added to the %v", r.UserID, "Workout Challenge role" ))
		}
	})
}

func DiscordRemoveReactionHandler(session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionRemove){
		fmt.Printf("%v removed reaction %v\n", r.UserID, r.Emoji.Name)

		// TODO: Currently this only works for this specific muscle tone. We need to make this work for all muscle tones
		if r.Emoji.Name == "💪🏿" {
			s.GuildMemberRoleRemove(r.GuildID, r.UserID, "1311487436278337638")
			s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v has been removed from %v", r.UserID, "Workout Challenge role" ))
		}
	})
}
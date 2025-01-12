package reactions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func AddWorkoutChallengeRole(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	s.GuildMemberRoleAdd(r.GuildID, r.UserID, "1311487436278337638")
	s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v has been added to the %v", r.UserID, "Workout Challenge role"))
}
func RemoveWorkoutChallengeRole(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	s.GuildMemberRoleRemove(r.GuildID, r.UserID, "1311487436278337638")
	s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v has been removed from %v", r.UserID, "Workout Challenge role"))
}

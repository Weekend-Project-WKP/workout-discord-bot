package services

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func InitializeDiscordGo() (*discordgo.Session, error) {
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("DISCORD_TOKEN not set in environment")
	}

	session, err := discordgo.New(discordToken)
	if err != nil {
		log.Fatal((err))
	}
	return session, err
}
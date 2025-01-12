package services

import (
	"context"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
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
func InitializeAiPartner() (*genai.GenerativeModel, context.Context, *genai.Client, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	model := client.GenerativeModel("gemini-1.5-flash")
	return model, ctx, client, err
}

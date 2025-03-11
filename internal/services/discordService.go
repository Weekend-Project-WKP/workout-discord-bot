package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"workoutbot/internal/db"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var DiscordSession *discordgo.Session
var AiModel *genai.GenerativeModel
var AiContext context.Context
var AiClient *genai.Client
var AiError error

func InitializeDiscordGo() error {
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("DISCORD_TOKEN not set in environment")
	}
	var err error
	DiscordSession, err = discordgo.New(discordToken)
	if err != nil {
		log.Fatal((err))
	}
	return err
}
func InitializeAiPartner() {
	AiContext = context.Background()
	AiClient, AiError = genai.NewClient(AiContext, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	AiModel = AiClient.GenerativeModel("gemini-2.0-flash")
}

func FatalSessionClosing() {
	if r := recover(); r != nil {
		rAsString := fmt.Sprintf("%v", r)
		DiscordSession.ChannelMessageSend("898225507064250420", "The bot is offline because of an error: "+rAsString)
		AiClient.Close()
		db.MongoClient.Disconnect(context.TODO())
		DiscordSession.Close()
		os.Exit(0)
	}
}

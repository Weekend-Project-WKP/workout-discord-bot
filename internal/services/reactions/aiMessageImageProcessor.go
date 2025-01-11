package reactions

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"workoutbot/internal/constants"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
)

func ProcessWorkout(s *discordgo.Session, r *discordgo.MessageReactionAdd, model *genai.GenerativeModel, ctx context.Context) {
	message, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		fmt.Println(err)
	}

	var parts []genai.Part
	textContext := constants.AiPrompt
	if message.Content != "" {
		textContext = textContext + message.Content
	}
	parts = append(parts, genai.Text(textContext))
	if len(message.Attachments) > 0 {
		for _, msg := range message.Attachments {
			fileData, err := getFile(msg.URL)
			if err != nil {
				fmt.Println(err)
			}
			parts = append(parts, genai.ImageData("jpeg", fileData))
		}
	}
	resp, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		log.Fatal(err)
	}
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				_, err := s.ChannelMessageSend(r.ChannelID, fmt.Sprintf("%v", part))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
func getFile(fileUrl string) ([]byte, error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil, err
	}
	defer resp.Body.Close() // Important: Close the response body when done

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Non-200 status code:", resp.StatusCode)
		return nil, err
	}

	// Read the response body (the file content)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	return body, nil
}

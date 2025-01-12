package reactions

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"workoutbot/internal/constants"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
)

func WhatWouldDavidGogginsSay(s *discordgo.Session, r *discordgo.MessageReactionAdd, model *genai.GenerativeModel, ctx context.Context) {
	message, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		fmt.Println(err)
	}
	textContext := constants.DavidGoginsAiPrompt
	if message.Content != "" {
		textContext = textContext + message.Content
	}
	processMessageAndAttachment(textContext, message, r.ChannelID, s, model, ctx)
}
func GetAiSummary(s *discordgo.Session, r *discordgo.MessageReactionAdd, model *genai.GenerativeModel, ctx context.Context) {
	message, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		fmt.Println(err)
	}

	teamName, errTeamName := getTeamName(s, r.GuildID, r.UserID)
	if errTeamName != nil {
		fmt.Println(errTeamName)
		s.ChannelMessageSend(r.ChannelID, "This user isn't assigned to a team. Need a team my guy/gal/they.")
	}

	textContext := constants.AiPrompt
	if message.Content != "" {
		textContext = textContext + message.Content
	}
	textContext = textContext + "The team name is " + teamName + ". "
	textContext = textContext + "The user name is " + message.Author.Username + "."
	processMessageAndAttachment(textContext, message, r.ChannelID, s, model, ctx)
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
func processMessageAndAttachment(textContext string, message *discordgo.Message, channelId string, s *discordgo.Session, model *genai.GenerativeModel, ctx context.Context) {
	var parts []genai.Part
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
				_, err := s.ChannelMessageSend(channelId, fmt.Sprintf("%v", part))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
func getTeamName(s *discordgo.Session, guildId string, userId string) (string, error) {
	serverRoleMap := make(map[string]string)
	roles, _ := s.GuildRoles(guildId)
	for _, serverRole := range roles {
		serverRoleMap[serverRole.ID] = serverRole.Name
	}
	member, _ := s.GuildMember(guildId, userId)
	for _, userRole := range member.Roles {
		if strings.Contains(serverRoleMap[userRole], "Team") {
			return serverRoleMap[userRole], nil
		}
	}
	return "", fmt.Errorf("no team found for this user")
}

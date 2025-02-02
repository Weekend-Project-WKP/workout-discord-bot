package reactions

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"workoutbot/internal/constants"
	"workoutbot/internal/db"
	"workoutbot/internal/helpers"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
)

func WhatWouldDavidGogginsSay(s *discordgo.Session, r *discordgo.MessageReactionAdd, model *genai.GenerativeModel, ctx context.Context, aiError error) {
	if !isAiAvailable(aiError, s, r.ChannelID) {
		return
	}
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
func GetAiSummary(s *discordgo.Session, r *discordgo.MessageReactionAdd, model *genai.GenerativeModel, ctx context.Context, aiError error) {
	if !isAiAvailable(aiError, s, r.ChannelID) {
		return
	}
	message, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		fmt.Println(err)
	}

	teamName, errTeamName := helpers.GetTeamName(s, r.GuildID, message.Author.ID)
	if errTeamName != nil {
		fmt.Println(errTeamName)
		s.ChannelMessageSend(r.ChannelID, "This user isn't assigned to a team. Need a team my guy/gal/they.")
	}

	textContext := constants.AiPrompt + getCategoriesForAiPrompt()
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
		panic(err)
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

func isAiAvailable(aiError error, s *discordgo.Session, channelId string) bool {
	if aiError != nil {
		s.ChannelMessageSend(channelId, constants.AiErrorMessage)
		return false
	}
	return true
}

func getCategoriesForAiPrompt() string {
	accumulatingCategoryPrompt := "The categories are:  "
	workoutCategoryMap, _ := db.WorkoutCategoryGetAll()
	for _, value := range workoutCategoryMap {
		accumulatingCategoryPrompt += value.CategoryName + " measured in " + value.MeasurementQuantification + ". "
	}
	return accumulatingCategoryPrompt
}

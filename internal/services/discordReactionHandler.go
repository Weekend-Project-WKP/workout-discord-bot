package services

import (
	"context"
	"fmt"
	"workoutbot/internal/services/reactions"

	"github.com/bwmarrin/discordgo"
	"github.com/google/generative-ai-go/genai"
)

func DiscordAddReactionHandler(session *discordgo.Session, model *genai.GenerativeModel, ctx context.Context, aiError error) {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
		fmt.Printf("%v reacted with %v\n", r.UserID, r.Emoji.Name)
		switch r.Emoji.Name {
		// TODO: Currently this only works for this specific muscle tone. We need to make this work for all muscle tones
		case "💪🏿":
			reactions.AddWorkoutChallengeRole(s, r)
		case "AI":
			reactions.GetAiSummary(s, r, model, ctx, aiError)
		case "Goggins":
			reactions.WhatWouldDavidGogginsSay(s, r, model, ctx, aiError)
		case "✅":
			reactions.SubmitWorkout(s, r)
		default:
			fmt.Printf("No Add Emoji Reaction Logic for %v\n", r.Emoji.Name)
		}
	})
}

func DiscordRemoveReactionHandler(session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
		fmt.Printf("%v removed reaction %v\n", r.UserID, r.Emoji.Name)
		switch r.Emoji.Name {
		// TODO: Currently this only works for this specific muscle tone. We need to make this work for all muscle tones
		case "💪🏿":
			reactions.RemoveWorkoutChallengeRole(s, r)
		default:
			fmt.Printf("No Remove Emoji Reaction Logic for %v\n", r.Emoji.Name)
		}
	})
}

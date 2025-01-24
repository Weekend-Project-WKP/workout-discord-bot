package services

import (
	"fmt"
	"workoutbot/internal/services/reactions"

	"github.com/bwmarrin/discordgo"
)

func DiscordAddReactionHandler() {
	DiscordSession.AddHandler(func(session *discordgo.Session, r *discordgo.MessageReactionAdd) {
		defer FatalSessionClosing()
		fmt.Printf("%v reacted with %v\n", r.UserID, r.Emoji.Name)
		switch r.Emoji.Name {
		// TODO: Currently this only works for this specific muscle tone. We need to make this work for all muscle tones
		case "💪🏿":
			reactions.AddWorkoutChallengeRole(session, r)
		case "AI":
			reactions.GetAiSummary(session, r, AiModel, AiContext, AiError)
		case "Goggins":
			reactions.WhatWouldDavidGogginsSay(session, r, AiModel, AiContext, AiError)
		case "✅":
			reactions.SubmitWorkout(session, r)
		default:
			fmt.Printf("No Add Emoji Reaction Logic for %v\n", r.Emoji.Name)
		}
	})
}

func DiscordRemoveReactionHandler() {
	DiscordSession.AddHandler(func(session *discordgo.Session, r *discordgo.MessageReactionRemove) {
		fmt.Printf("%v removed reaction %v\n", r.UserID, r.Emoji.Name)
		switch r.Emoji.Name {
		// TODO: Currently this only works for this specific muscle tone. We need to make this work for all muscle tones
		case "💪🏿":
			reactions.RemoveWorkoutChallengeRole(session, r)
		default:
			fmt.Printf("No Remove Emoji Reaction Logic for %v\n", r.Emoji.Name)
		}
	})
}

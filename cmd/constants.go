package cmd

import (
	"github.com/bwmarrin/discordgo"
)

var (
	SucessfulEventSetMsg = "Event (%s) successfully scheduled @%s :date: \nParticipants:\n"
	CreditsMsg = `CalendarAssistant is a helper bot made by a_walking_dead
	Discord: a_walking_dead#9013
	Instagram: a_w4lking_dead
	Telegram: @massive18dynamics
	`

	ErrInvalidTimeInput = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Please use proper 24 hour format.",
		},
	}
	ErrInvalidEmailInput = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Invalid gmail input.",
		},
	}
	ErrInvalidAttendeesInput = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Invalid participants input. Emails should be separated by dash (-)",
		},
	}
)
package cmd

import (
	"github.com/bwmarrin/discordgo"
)

var (
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
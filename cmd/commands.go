package cmd

import (
	"github.com/bwmarrin/discordgo"
)

var (
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			Name: "credits",
			Description: "About me",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"credits": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "calendar-bot is a helper bot made by a_walking_dead#9013.",
				},
			})
		},
	}
)

func InitializeBotCommands(s *discordgo.Session) []*discordgo.ApplicationCommand {
	var initializedCmds []*discordgo.ApplicationCommand
	for _, cmd := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			panic("couldn't initialize commands")
		}
		s.AddHandler(commandHandlers[cmd.Name])
		initializedCmds = append(initializedCmds, cmd)
	}

	return initializedCmds
}

func DeactivateBotCommands(s *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	for _, cmd := range commands {
		s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
	}
}
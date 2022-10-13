package cmd

import (
	"github.com/bwmarrin/discordgo"
)

var (
	defaultMemberPermissions int64   = discordgo.PermissionManageServer
	integerMinVal            float64 = 0

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "credits",
			Description: "About me",
		},
		{
			Name:        "new",
			Description: "Set up a new meeting",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "title",
					Description: "Main point of the event.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "description",
					Description: "Details of the event.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "day",
					Description: "Day of meeting (0 for today, 1 tomorrow, 2 day after tomorrow, and so on",
					MinValue:    &integerMinVal,
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "start",
					Description: "Use 24h format. (e.g. 13:45, 07:00)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "end",
					Description: "Use 24h format. (e.g. 13:45, 07:00)",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "participants",
					Description: "Emails separated by - | e.g. abc@gmail.com,def@gmail.com",
					Required:    true,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"credits": credits,
		"new": setNewCalEvent,
	}
)

func InitializeBotCommands(s *discordgo.Session) []*discordgo.ApplicationCommand {
	var initializedCmds []*discordgo.ApplicationCommand
	for _, cmd := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			panic(err)
		}
		initializedCmds = append(initializedCmds, cmd)
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	return initializedCmds
}

func DeactivateBotCommands(s *discordgo.Session, commands []*discordgo.ApplicationCommand) {
	for _, cmd := range commands {
		s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
	}
}

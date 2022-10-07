package cmd

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	cal "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/faramarz-hosseini/calendar-bot.git/Gcal"
)

var (
	calendar = Gcal.InitClient()
	calendarService, _ = cal.NewService(
		context.Background(), option.WithHTTPClient(calendar),
	)

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
					Required:   true,
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
		"credits": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			fmt.Println("I was called here")
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "CalendarAssistant is a helper bot made by a_walking_dead#9013.",
				},
			})
		},
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

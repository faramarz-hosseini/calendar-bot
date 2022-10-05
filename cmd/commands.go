package cmd

import (
	"context"
	"fmt"
	"time"
	"strings"

	"github.com/bwmarrin/discordgo"
	cal "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/faramarz-hosseini/calendar-bot.git/Gcal"
)

const (
	botCalendarID = "primary"
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
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "day",
					Description: "Day of meeting (0 for today, 1 tomorrow, 2 day after tomorrow, and so on",
					MinValue:    &integerMinVal,
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "time",
					Description: "Use 24h format. (e.g. 13:45, 07:00)",
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

func setNewCalEvent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	requestedDay := options[0].Value.(float64)
	// requestedTime := options[1].Value.(string)
	eventTime := time.Now().AddDate(0, 0, int(requestedDay)).Format(time.RFC3339)
	st := strings.Split(eventTime, "T")

	
	fmt.Println(st[0])
	// event := &cal.Event{
	// 	Attendees: []*cal.EventAttendee{
	// 		{Email: "faramarz.hosseini99@gmail.com"},
	// 	},
	// 	Start: &cal.EventDateTime{
	// 		DateTime: "",
	// 	},
	// }
	// calendarService.Events.Insert(botCalendarID, event)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "bit queen doos dari?",
		},
	})
}


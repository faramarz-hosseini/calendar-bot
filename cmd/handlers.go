package cmd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	cal "google.golang.org/api/calendar/v3"

	"github.com/faramarz-hosseini/calendar-bot.git/utils"
)

const (
	botCalendarID = "primary"

	iranTimezone = "Asia/Tehran"
)

func setNewCalEvent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	reqTitle, reqDesc := options[0].Value.(string), options[1].Value.(string)
	reqDay := options[2].Value.(float64)
	reqStartTime, reqEndTime := options[3].Value.(string), options[4].Value.(string)
	reqAttendees := options[5].Value.(string)

	if !utils.IsValidTimeString(reqStartTime) ||
	!utils.IsValidTimeString(reqEndTime) {
		s.InteractionRespond(i.Interaction, ErrInvalidTimeInput)
		return
	}
	if !utils.IsValidAttendeesInput(reqAttendees) {
		s.InteractionRespond(i.Interaction, ErrInvalidAttendeesInput)
		return
	}

	var attendees []*cal.EventAttendee
	attendeesEmails := strings.Split(reqAttendees, "-")
	for _, email := range attendeesEmails {
		if email == "" {
			continue
		}
		trimmedEmail := strings.TrimSpace(email)
		if !utils.IsValidGmailInput(trimmedEmail) {
			s.InteractionRespond(i.Interaction, ErrInvalidEmailInput)
			return
		}
		attendees = append(attendees, &cal.EventAttendee{Email: trimmedEmail})
	}

	dateString := utils.GenerateDateStringFromCmdInp(int(reqDay))
	reqStartTime += ":00"
	reqEndTime += ":00"
	eventStartTime := dateString+"T"+reqStartTime
	eventEndTime := dateString+"T"+reqEndTime

	event := &cal.Event{
		Summary: reqTitle,
		Description: reqDesc,
		Attendees: attendees,
		Start: &cal.EventDateTime{
			DateTime: eventStartTime,
			TimeZone: iranTimezone,
		},
		End: &cal.EventDateTime{
			DateTime: eventEndTime,
			TimeZone: iranTimezone,
		},
	}
	createdEvent, err := calendarService.Events.Insert(botCalendarID, event).Do()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(createdEvent)
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Event successfully set.",
		},
	})
}


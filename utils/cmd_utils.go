package utils

import (
	"regexp"
	"strings"
	"time"
)

const (
	twentyFourHourTimeFormatRegex = "^(2[0-3]|[01]?[0-9]):[0-5][0-9]" 
)

var (
	twentyFourHourValidator, _ = regexp.Compile(twentyFourHourTimeFormatRegex)
)

func GenerateDateStringFromCmdInp(days int) string {
	eventTime := time.Now().AddDate(0, 0, int(days)).Format(time.RFC3339)
	return strings.Split(eventTime, "T")[0]
}

func IsValidTimeString(timeStr string) bool {
	return twentyFourHourValidator.MatchString(timeStr)
}

func IsValidAttendeesInput(inp string) bool {
	return strings.Contains(inp, "-")
}

func IsValidGmailInput(inp string) bool {
	return strings.Contains(inp, "@gmail.com") 
}

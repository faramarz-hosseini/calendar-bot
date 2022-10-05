package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/faramarz-hosseini/calendar-bot.git/cmd"

	"github.com/bwmarrin/discordgo"
)


func main() {
	// t := time.Now().Format(time.RFC3339)
	// events, err := srv.Events.List("primary").ShowDeleted(false).
	// 		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	// if err != nil {
	// 		fmt.Printf("Unable to retrieve next ten of the user's events: %v", err)
	// }
	// fmt.Println("Upcoming events:")
	// if len(events.Items) == 0 {
	// 		fmt.Println("No upcoming events found.")
	// } else {
	// 		for _, item := range events.Items {
	// 				date := item.Start.DateTime
	// 				if date == "" {
	// 						date = item.Start.Date
	// 				}
	// 				fmt.Printf("%v (%v)\n", item.Summary, date)
	// 		}
	// }

	cli, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic("could not start client")
	}
	err = cli.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	commands := cmd.InitializeBotCommands(cli)
	defer cmd.DeactivateBotCommands(cli, commands)

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
    <-sc
	cli.Close()
}
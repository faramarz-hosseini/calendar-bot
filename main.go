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
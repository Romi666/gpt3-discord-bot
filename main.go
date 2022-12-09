package main

import (
	"fmt"
	"github.com/Romi666/gpt3-discord-bot/completion"
	"github.com/Romi666/gpt3-discord-bot/connection"
	discord_gpt "github.com/Romi666/gpt3-discord-bot/discord-gpt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env files : %v", err)
		return
	}
}
func main() {
	clientGpt := connection.InitConnectionGPT()
	dg, err := connection.InitConnectionDiscord()
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	comp := completion.NewCompletion(clientGpt)
	dgpt := discord_gpt.NewDiscordGPT(comp)

	dg.AddHandler(dgpt.CreateCompletion)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

package connection

import (
	"github.com/bwmarrin/discordgo"
	"os"
)

func InitConnectionDiscord() (s *discordgo.Session, err error) {
	var API_KEY = os.Getenv("DISCORD_API_KEY")
	dg, err := discordgo.New("Bot " + API_KEY)
	if err != nil {
		return nil, err
	}

	return dg, nil
}


package discord_gpt

import (
	"github.com/Romi666/gpt3-discord-bot/completion"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type DiscordGPT struct {
	ClientGPT	completion.Completion
}

func NewDiscordGPT(clientGPT completion.Completion) DiscordGPT {
	return DiscordGPT{
		ClientGPT: clientGPT,
	}
}

func (d *DiscordGPT) CreateCompletion(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if  strings.HasPrefix(m.Content, "!q") {
	//	@TODO
	}
}

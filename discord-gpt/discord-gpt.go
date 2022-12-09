package discord_gpt

import (
	"context"
	"github.com/Romi666/gpt3-discord-bot/completion"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type DiscordGPT struct {
	ClientGPT completion.Completion
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

	if strings.HasPrefix(m.Content, "!q") {
		//	@TODO
		args := strings.TrimLeft(m.Content, "!q")
		resp, err := d.ClientGPT.SendCompletionsRequest(context.Background(), args)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "error generate answer, try again")
			return
		}

		s.ChannelMessageSend(m.ChannelID, resp.Choices[0].Text)
	}
}

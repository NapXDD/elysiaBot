package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
        return
    }

    msg := strings.TrimSpace(strings.ToLower(m.Content))
    if msg == "3" {
        s.ChannelMessageSend(m.ChannelID, "6")
    }
}
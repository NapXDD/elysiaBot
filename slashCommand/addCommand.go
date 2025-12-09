package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)



func addCommand(session *discordgo.Session, GuildID string){
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}
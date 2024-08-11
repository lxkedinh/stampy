package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var cmds = []*discordgo.ApplicationCommand{
	Timestamp,
}

func InitCommands(s *discordgo.Session) {
	for _, c := range cmds {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", c)
		if err != nil {
			log.Panicf("Could not create '%v' command: %v", c.Name, err)
		}
	}
}

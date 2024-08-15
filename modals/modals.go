package modals

import (
	"github.com/bwmarrin/discordgo"
)

type Modal struct {
	Id            string
	OpenHandler   func(s *discordgo.Session, i *discordgo.InteractionCreate) error
	SubmitHandler func(s *discordgo.Session, i *discordgo.InteractionCreate) error
}

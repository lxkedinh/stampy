package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate) error

var commandHandlers = map[string]CommandHandler{
	Timestamp.Name: HandleTimestamp,
}

func InitCommandHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			err := handler(s, i)
			if err != nil {
				log.Println(err)
				handleCommandError(s, i)
			}
		}
	})
}

func handleCommandError(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Sorry, I could not respond to this command. Try again!",
		},
	})
}

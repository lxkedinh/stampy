package modals

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var modalHandlers = map[string]Modal{
	TimezoneModal.Id: TimezoneModal,
}

func InitModalSubmitHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionModalSubmit {
			return
		}
		if modal, ok := modalHandlers[i.ModalSubmitData().CustomID]; ok {
			err := modal.SubmitHandler(s, i)
			if err != nil {
				log.Println(err)
			}
		}
	})
}

package modals

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var TimezoneModal Modal = Modal{
	Id:            "modal_timezone",
	OpenHandler:   handleOpenTimezone,
	SubmitHandler: handleSubmitTimezone,
}

func handleOpenTimezone(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "modal_timezone",
			Title:    "Timezone Prompt",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "modal_timezone",
							Label:       "What is your local timezone?",
							Style:       discordgo.TextInputShort,
							Placeholder: "American/Los_Angeles",
							Required:    true,
							MaxLength:   50,
							MinLength:   3,
						},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func handleSubmitTimezone(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	modalData := i.ModalSubmitData()
	timezoneInput := modalData.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("You entered the timezone \"%s\"", timezoneInput),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

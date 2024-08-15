package commands

import (
	"bytes"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/lxkedinh/stampy/templates"
	"github.com/lxkedinh/stampy/timestamp"
)

var Timestamp = &discordgo.ApplicationCommand{
	Name:        "timestamp",
	Description: "Generate a timestamp message",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:         "datetime",
			Description:  "The date & time you want to generate a timestamp for. Defaults to the current time if left empty.",
			Type:         discordgo.ApplicationCommandOptionString,
			Required:     false,
			Autocomplete: true,
		},
		{
			Name:        "timezone",
			Description: "Enter the timezone you want to use. Defaults to your local timezone if left empty.",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
		},
	},
}

func HandleTimestampCmd(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	// commandOptions := i.Interaction.ApplicationCommandData().Options
	snowflake := i.ID
	t := timestamp.FromSnowflake(snowflake)
	response, err := execTimestampTemplate(t)
	if err != nil {
		return err
	}

	// TODO: ask for timezone instead if user's timezone cannot be found in redis db
	// modals.TimezoneModal.OpenHandler(s, i)
	// return nil

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: response,
		},
	})
	return err
}

func execTimestampTemplate(t timestamp.Timestamp) (string, error) {
	buf := &bytes.Buffer{}
	err := templates.Timestamp.Execute(buf, t)
	if err != nil {
		return "", errors.New("could not execute timestamp response message template")
	}
	return buf.String(), nil
}

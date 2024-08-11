package commands

import (
	"bytes"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/lxkedinh/stampy/timestamp"
	"github.com/lxkedinh/stampy/tmpl"
)

var Timestamp = &discordgo.ApplicationCommand{
	Name:        "timestamp",
	Description: "Generate a timestamp message",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:         "datetime",
			Description:  "The date/time you want to generate a timestamp for.",
			Type:         discordgo.ApplicationCommandOptionString,
			Required:     false,
			Autocomplete: true,
		},
	},
}

func HandleTimestamp(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	// commandOptions := i.Interaction.ApplicationCommandData().Options
	snowflake := i.ID
	t := timestamp.FromSnowflake(snowflake)
	response, err := execTimestampTemplate(t)
	if err != nil {
		return err
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: response},
	})
	return err
}

func execTimestampTemplate(t timestamp.Timestamp) (string, error) {
	buf := &bytes.Buffer{}
	err := tmpl.Timestamp.Execute(buf, t)
	if err != nil {
		return "", errors.New("could not execute timestamp response message template")
	}
	return buf.String(), nil
}

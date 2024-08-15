package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/lxkedinh/stampy/commands"
	"github.com/lxkedinh/stampy/db"
	"github.com/lxkedinh/stampy/env"
	"github.com/lxkedinh/stampy/modals"
)

var session *discordgo.Session

func init() {
	env.Load("./.env")
}

func init() {
	db.InitRedisClient()
}

func init() {
	var err error
	session, err = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panicf("Could not initialize Discord bot\n%v", err)
	}
	log.Println("Stampy bot initialized")
}

func init() {
	commands.InitCommandHandlers(session)
	modals.InitModalSubmitHandlers(session)
}

func main() {
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Stampy bot logged in")
	})
	err := session.Open()
	if err != nil {
		log.Panicf("Could not open a session\n%v", err)
	}
	defer session.Close()

	commands.InitCommands(session)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
	log.Println("Shutting down bot...")
}

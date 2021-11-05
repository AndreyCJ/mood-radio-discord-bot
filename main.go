package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(chatMessageHandler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Ready to GOOOOOOOOO")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func chatMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	channelId := m.ChannelID

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!listen" {
		const message = "Здарова нахуй!!!!!!!!!11"
		fmt.Println("message sent => ", message)
		s.ChannelMessageSend(channelId, message)
	}

	if m.Content == "!эй" {
		const message = "кок <:9716_Pepega:752249224736800778>"
		fmt.Println("message sent => ", message)
		s.ChannelMessageSend(channelId, message)
	}
}

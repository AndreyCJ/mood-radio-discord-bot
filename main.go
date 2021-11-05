package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	CommandPrefix string = "!"
	Token         string
	ClientId      string
	ChannelId     string
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

	dg.AddHandler(botReadyHandler)
	dg.AddHandler(chatMessageHandler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func chatMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	userMessage := m.Content
	userMessageBody := getUserMessageBody(userMessage)

	if isCommandMessage(userMessage, "listen") {
		const message = "Здарова нахуй!!!!!!!!!11"
		s.ChannelMessageSend(ChannelId, message)
	}

	if isCommandMessage(userMessage, "эй") {
		const message = "кок <:9716_Pepega:752249224736800778>"
		s.ChannelMessageSend(ChannelId, message)
	}

	if isCommandMessage(userMessage, "tts") && userMessageBody != "" {
		s.ChannelMessageSendTTS(ChannelId, userMessageBody)
	}
}

func botReadyHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	ChannelId = m.ChannelID
}

func isCommandMessage(message, command string) bool {
	return strings.HasPrefix(message, CommandPrefix+command)
}

func getUserMessageBody(message string) string {
	userMessageBody := strings.Split(message, " ")

	if len(userMessageBody) > 1 {
		return strings.Join(userMessageBody[1:], " ")
	}
	return ""
}

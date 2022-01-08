package bot

import (
	"fmt"
	"strings"

	"github.com/AndreyCJ/mood-radio-discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session
var ClientId string
var ChannelId string

func Start() {
	//creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Making our bot a user using User function .
	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(botReadyHandler)
	goBot.AddHandler(chatMessageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running! :)")
}

func Close() {
	goBot.Close()
}

func chatMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
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

	if isCommandMessage(userMessage, "умри") {
		const message = "покаааааааа"
		s.ChannelMessageSend(ChannelId, message)
		Close()
	}

	if isCommandMessage(userMessage, "tts") && userMessageBody != "" {
		s.ChannelMessageSendTTS(ChannelId, userMessageBody)
	}
}

func botReadyHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	ChannelId = m.ChannelID
}

func isCommandMessage(message, command string) bool {
	return strings.HasPrefix(message, config.CommandPrefix+command)
}

func getUserMessageBody(message string) string {
	userMessageBody := strings.Split(message, " ")

	if len(userMessageBody) > 1 {
		return strings.Join(userMessageBody[1:], " ")
	}
	return ""
}

package message

import (
	"discordbotgo/music"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "!"

// Handler for when a message is sent in chat,
// I think the name is important
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var err error = nil
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	//TODO This should probably be done in a more modular way
	switch m.Content {
	case prefix + "ping":
		_, err = s.ChannelMessageSend(m.ChannelID, "Pong!")
	case prefix + "pong":
		_, err = s.ChannelMessageSend(m.ChannelID, "Ping!")
	case prefix + "who":
		_, err = s.ChannelMessageSend(m.ChannelID, s.State.User.Username+"#"+s.State.User.Discriminator)
	case prefix + "play":
		music.Play(s, "buttman")
		_, err = s.ChannelMessageSend(m.ChannelID, "Play woo")
	}

	if err != nil {
		fmt.Println("Error sending message: ", err)
		return
	}
}

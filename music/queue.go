package music

import "github.com/bwmarrin/discordgo"

func Queue(s *discordgo.Session, query string) {
	// Check if there is a queue already for the guild
	// Create one if not
	// When play command is sent add that song to the queue
	// When song is over remove from queue
	q := make([]string, 100)

	q[0] = query
	println(len(q))
}

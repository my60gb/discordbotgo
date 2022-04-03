package main

import (
	"discordbotgo/message"
	"discordbotgo/music"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	go music.YT()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TESTTOKEN := os.Getenv("TESTTOKEN")

	if TESTTOKEN == "" {
		log.Fatal("Uhh you need a token you dork")
	}

	s, err := discordgo.New("Bot " + TESTTOKEN)
	if err != nil {
		log.Fatal("Issue with making session: ", err)
	}
	fmt.Println("Connected!")

	s.AddHandler(message.MessageCreate)

	s.Identify.Intents = discordgo.IntentsGuildMessages

	err = s.Open()

	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}
	//SetupCloseHandler(s, err)

	fmt.Println(s.State.User.Username + "#" + s.State.User.Discriminator)
	if err != nil {
		fmt.Println(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	fmt.Println()
	s.Close()

}

// Function to wait for os close signal and close connection when detected
func SetupCloseHandler(s *discordgo.Session, err error) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		s.Close()
		if err != nil {
			log.Fatal("Some error: ", err)
		}
		fmt.Println("\nClosed!")
		os.Exit(0)
	}()
}

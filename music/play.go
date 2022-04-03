package music

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/wader/goutubedl"
)

func Play(s *discordgo.Session, q string) {
	fmt.Println(q)
}

func YT() {
	result, err := goutubedl.New(context.Background(), "https://www.youtube.com/watch?v=jgVhBThJdXc", goutubedl.Options{})
	if err != nil {
		log.Fatal(err)
	}

	goutubedl.Path = "yt-dlp"
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		log.Fatal(err)
	}
	defer downloadResult.Close()
	f, err := os.Create("output")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, downloadResult)
}

package main

import (
	"github.com/knspriggs/go-twitch"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting to watch for changes...")
	session, err := twitch.NewSession(twitch.NewSessionInput{ClientID: "s5374c3j2vgdtl3gtbqxdla35gfvhv"})
	
	if err != nil {
		panic(err)
	}
	
	ticker := time.NewTicker(1 * time.Second)
	
	for {
		select {
		case <- ticker.C:
			checkStream(session)
		}
	}
}

// Checks if the stream is on.
func checkStream(session *twitch.Session) {
	stream, err := session.GetStream(&twitch.GetStreamsInputType{Channel: "ToaztyTV"})
	if err != nil {
		panic(err)
	}
	
	if stream.Total == 1 {
		fmt.Println("Channel is online, sending message to Discord channel...")
		// TODO: output message to discord
	}
}

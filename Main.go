package main

import (
    "fmt"
    "github.com/onestay/go-new-twitch"
    "github.com/bwmarrin/discordgo"
    "os"
    "time"
)

var (
    twitchToken    string
    discordToken   string
    twitchChannel  string
    discordChannel string
)


func main() {
    
    if len(os.Args) < 5 {
        fmt.Println("Too few arguments. Expecting: <twitchToken> <discordToken> <twitchChannel> <discordChannel>")
        return
    }
    
    twitchToken    = os.Args[1]
    discordToken   = os.Args[2]
    twitchChannel  = os.Args[3]
    discordChannel = os.Args[4]
    
    twitchSession       := twitch.NewClient(twitchToken)
    discordSession, err := discordgo.New("Bot " + discordToken)
    discordSession.Open()
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    defer discordSession.Close()
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
	fmt.Println("Starting to watch for changes...")
	ticker := time.NewTicker(10 * time.Second)
	
	for {
		select {
		case <- ticker.C:
			checkStream(twitchSession, discordSession)
		}
	}
}


// Checks if the stream is on and sends message to specified Discord channel.
func checkStream(tc *twitch.Client, ds *discordgo.Session) {
	streamData, err := tc.GetStreams(twitch.GetStreamsInput{UserLogin: []string{twitchChannel}})
	
	if err == nil {
        if len(streamData) >= 1 {
            fmt.Println("Channel is online, sending message to Discord channel...")
            ds.ChannelMessageSend(discordChannel, "Ich bin jetzt Live auf Twitch: \n`" + streamData[0].Title + "`\nhttps://www.twitch.tv/" + twitchChannel)
        }
    } else {
        fmt.Println(err)
        fmt.Println("Trying again in 5 seconds...")
        time.Sleep(5 * time.Second)
    }
}

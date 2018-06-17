package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/bwmarrin/discordgo"
)

var (
    Token string
)

func main()  {

    Token := os.Getenv("TOKEN")
    dg, err := discordgo.New("Bot " + Token)
    if err != nil {
        fmt.Println("error creating Discord session,", err)
        return
    }

    // Register the messageCreate func as a callback for MessageCreate events.
    dg.AddHandler(messageCreate)

    // Open a websocket connection to Discord and begin listening.
    err = dg.Open()
    if err != nil {
        fmt.Println("error opening connection,", err)
        return
    }

    // Wait here until CTRL-C or other term signal is received.
    fmt.Println("Bot is now running.  Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

    if m.Author.ID == s.State.User.ID {
		return
	}

    fmt.Println(m.Content)

    if m.Content == "ping" {
        s.ChannelMessageSend(m.ChannelID, "pong")
    }
}

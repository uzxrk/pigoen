package main

import (
    "log"
    "github.com/uzxrk/pigoen/internal/archiver"
    "github.com/uzxrk/pigoen/internal/slack"
)

func main() {
    config, err := archiver.LoadConfig("configs/config.yaml")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    client := slack.NewClient(config.SlackToken)
	archiver := archiver.NewArchiver(client)
	filePath := "messages.txt"

    err = archiver.ArchiveMessages(config.ChannelID, filePath)
    if err != nil {
        log.Fatalf("Error archiving messages: %v", err)
    }

    log.Printf("Messages archived to %s", filePath)
}
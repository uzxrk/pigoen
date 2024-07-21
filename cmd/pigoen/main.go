package main

import (
	"log"

	"github.com/uzxrk/pigoen/internal/archiver"
	"github.com/uzxrk/pigoen/internal/cli"
	"github.com/uzxrk/pigoen/internal/slack"
)

func main() {
	config, err := archiver.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := slack.NewClient(config.SlackToken)
	archiver := archiver.NewArchiver(client)
	encryptionKey := []byte(config.EncryptionKey) // 32 bytes for AES-256

	if len(encryptionKey) != 32 {
		log.Fatalf("Encryption key must be 32 bytes long for AES-256")
	}

	cliConfig, err := cli.ParseFlags()
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
	}

	if cliConfig.Mode == "archive" {
		err := archiver.ArchiveMessages(cliConfig.Channel, cliConfig.FilePath, encryptionKey)
		if err != nil {
			log.Fatalf("Error archiving messages: %v", err)
		}

		log.Printf("Messages archived to %s", cliConfig.FilePath)
	} else if cliConfig.Mode == "decrypt" {
		startTime, err := cli.ParseTime(cliConfig.Start)
		if err != nil {
			log.Fatalf("Error parsing start time: %v", err)
		}
		endTime, err := cli.ParseTime(cliConfig.End)
		if err != nil {
			log.Fatalf("Error parsing end time: %v", err)
		}

		err = archiver.DecryptMessages(cliConfig.FilePath, encryptionKey, cliConfig.User, startTime, endTime)
		if err != nil {
			log.Fatalf("Error decrypting messages: %v", err)
		}
	}
}

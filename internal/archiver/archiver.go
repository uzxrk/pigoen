package archiver

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/uzxrk/pigoen/internal/slack"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SlackToken    string `yaml:"slack_token"`
	ChannelID     string `yaml:"channel_id"`
	EncryptionKey string `yaml:"encryption_key"`
}

type Archiver struct {
	client *slack.Client
}

func NewArchiver(client *slack.Client) *Archiver {
	return &Archiver{client: client}
}

func (a *Archiver) ArchiveMessages(channelID, filePath string, encryptionKey []byte) error {
	messages, err := a.client.ReadMessages(channelID)
	if err != nil {
		return fmt.Errorf("error reading messages: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	userCache := make(map[string]string)

	for _, message := range messages {
		userName, ok := userCache[message.User]
		if !ok {
			user, err := a.client.GetUser(message.User)
			if err != nil {
				return fmt.Errorf("error fetching user info: %w", err)
			}
			userName = user.Name
			userCache[message.User] = userName
		}

		timestampParts := strings.Split(message.Timestamp, ".")
		if len(timestampParts) == 0 {
			return fmt.Errorf("invalid timestamp: %s", message.Timestamp)
		}

		seconds, err := strconv.ParseInt(timestampParts[0], 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing timestamp: %w", err)
		}

		timestamp := time.Unix(seconds, 0).Format(time.RFC3339)
		line := fmt.Sprintf("User: %s, Text: %s, Timestamp: %s\n", userName, message.Text, timestamp)
		encryptedLine, err := Encrypt([]byte(line), encryptionKey)
		if err != nil {
			return fmt.Errorf("error encrypting message: %w", err)
		}

		if _, err := file.WriteString(encryptedLine + "\n"); err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	return nil
}

func (a *Archiver) DecryptMessages(filePath string, encryptionKey []byte, userFilter string, startTime, endTime time.Time) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		encryptedLine := scanner.Text()
		decryptedLine, err := Decrypt(encryptedLine, encryptionKey)
		if err != nil {
			return fmt.Errorf("error decrypting message: %w", err)
		}

		if userFilter != "" {
			if !strings.Contains(decryptedLine, fmt.Sprintf("User: %s,", userFilter)) {
				continue
			}
		}

		if !startTime.IsZero() || !endTime.IsZero() {
			timestampStr := extractTimestamp(decryptedLine)
			if timestampStr != "" {
				timestamp, err := time.Parse(time.RFC3339, timestampStr)
				if err == nil {
					if !startTime.IsZero() && timestamp.Before(startTime) {
						continue
					}
					if !endTime.IsZero() && timestamp.After(endTime) {
						continue
					}
				}
			}
		}

		fmt.Println(decryptedLine)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

func extractTimestamp(message string) string {
	parts := strings.Split(message, ",")
	for _, part := range parts {
		if strings.HasPrefix(strings.TrimSpace(part), "Timestamp:") {
			return strings.TrimSpace(strings.TrimPrefix(part, "Timestamp:"))
		}
	}
	return ""
}

func LoadConfig(filepath string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}

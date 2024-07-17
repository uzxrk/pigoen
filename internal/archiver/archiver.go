package archiver

import (
    "fmt"
    "os"
    "github.com/uzxrk/pigoen/internal/slack"
    "strconv"
    "strings"
    "time"

	"gopkg.in/yaml.v2"
)

type Config struct {
    SlackToken string `yaml:"slack_token"`
	ChannelID string `yaml:"channel_id"`
}

type Archiver struct {
    client *slack.Client
}

func NewArchiver(client *slack.Client) *Archiver {
    return &Archiver{client: client}
}

func (a *Archiver) ArchiveMessages(channelID, filePath string) error {
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
        if _, err := file.WriteString(line); err != nil {
            return fmt.Errorf("error writing to file: %w", err)
        }
    }

    return nil
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
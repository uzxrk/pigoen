package cli

import (
	"fmt"
	"time"

	flag "github.com/spf13/pflag"
)

type Config struct {
	Mode     string
	Channel  string
	FilePath string
	User     string
	Start    string
	End      string
}

func ParseFlags() (*Config, error) {
	config := &Config{}

	flag.StringVarP(&config.Mode, "mode", "m", "", "Mode of operation: archive or decrypt")
	flag.StringVarP(&config.Channel, "channel", "c", "", "Slack channel ID (required for archive mode)")
	flag.StringVarP(&config.FilePath, "file", "f", "messages.txt", "File path for input/output")
	flag.StringVarP(&config.User, "user", "u", "", "Filter messages by user (optional for decrypt mode)")
	flag.StringVarP(&config.Start, "start", "s", "", "Start time in RFC3339 format (optional for decrypt mode)")
	flag.StringVarP(&config.End, "end", "e", "", "End time in RFC3339 format (optional for decrypt mode)")

	flag.Parse()

	if config.Mode != "archive" && config.Mode != "decrypt" {
		flag.Usage()
		return nil, fmt.Errorf("mode must be either 'archive' or 'decrypt'")
	}

	if config.Mode == "archive" && config.Channel == "" {
		flag.Usage()
		return nil, fmt.Errorf("channel is required in archive mode")
	}

	return config, nil
}

func ParseTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, timeStr)
}

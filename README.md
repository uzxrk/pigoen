# Pigoen - A Slack Message Archiver

A Go application that archives and decrypts Slack messages from specific channels. This app can archive messages from Slack channels, encrypt them, and provide functionality to decrypt and filter messages through a command-line interface (CLI).

## Features

- Archive messages from Slack channels.
- Encrypt archived messages using AES-256 encryption.
- Decrypt archived messages.
- Filter decrypted messages by user, timestamp, or a range of timestamps.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or later)
- Slack API token with permissions for reading messages (e.g., `channels:history`, `groups:history`, `im:history`, and `mpim:history`).

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/uzxrk/pigoen.git
   cd pigoen
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build -o bin/pigoen cmd/pigoen/main.go
   ```

## Configuration

Create a configuration file `config.yaml` in the `configs/` directory. This file should contain your Slack API token and the encryption key.

```yaml
slack_token: "xoxb-your-slack-token"
encryption_key: "your-32-byte-long-encryption-key-here" # 32 bytes for AES-256 encryption
channel_id: "your-channel-id" # optional, as you can also add the channel-id through the CLI.
```

## Usage

The application can be run with different modes using CLI flags. The modes are `archive` (to archive and encrypt messages) and `decrypt` (to decrypt and display messages).

### Archive Messages

To archive messages from a Slack channel and save them to a file, use the following command:

```bash
./pigoen --mode=archive --channel="C1234567890" --file="messages.txt"
```

- **Required Flags:**
  - `--mode`: Operation mode (`archive` or `decrypt`).
  - `--channel`: Slack channel ID to archive messages from.
  
- **Optional Flags:**
  - `--file`: File path for storing the archived messages (default: `messages.txt`).

### Decrypt Messages

To decrypt messages from the archived file and display them, use the following command:

```bash
./pigoen --mode=decrypt --file="messages.txt"
```

- **Required Flags:**
  - `--mode`: Operation mode (`archive` or `decrypt`).

- **Optional Flags:**
  - `--file`: Path to the archived file (default: `messages.txt`).
  - `--user`: Filter messages by Slack user name.
  - `--start`: Filter messages after this start timestamp (RFC3339 format, e.g., `2024-01-01T00:00:00Z`).
  - `--end`: Filter messages before this end timestamp (RFC3339 format).

### Examples

1. **Archive Messages from a Channel**:
   ```bash
   ./pigoen --mode=archive --channel="C1234567890" --file="messages.txt"
   ```

2. **Decrypt and Filter Messages by User**:
   ```bash
   ./pigoen --mode=decrypt --file="messages.txt" --user="john.doe"
   ```

3. **Decrypt Messages Between a Timestamp Range**:
   ```bash
   ./pigoen --mode=decrypt --file="messages.txt" --start="2024-01-01T00:00:00Z" --end="2024-01-31T23:59:59Z"
   ```

## How it Works

### Archiving Messages

When running in `archive` mode, the app:
- Fetches messages from the specified Slack channel.
- Archives them to a text file.
- Encrypts the messages using AES-256 encryption for secure storage.

### Decrypting Messages

In `decrypt` mode, the app:
- Decrypts the archived message file.
- Filters messages based on the specified user or timestamp range.
- Displays the decrypted messages in the console.

## Dependencies

- [pflag](https://github.com/spf13/pflag): Used for command-line flag parsing.
- [Slack API](https://github.com/slack-go/slack): Go client for Slack's Web API.
- [Go stdlib crypto package](https://pkg.go.dev/crypto): Used for encryption and decryption.

## Contributing

Feel free to open issues or submit pull requests for improvements, bug fixes, or new features.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---
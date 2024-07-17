# Pigoen - A Slack Message Archiver

A simple Go application to archive Slack messages.

## Getting Started

### Prerequisites

- Go 1.20 or later
- A Slack API token

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/uzxrk/pigoen.git
    cd pigoen
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Configure your Slack token in `configs/config.yaml`:
    ```yaml
    slack_token: "YOUR_SLACK_API_TOKEN"
    ```

### Usage

To run the application:
```sh
go run cmd/archiver/main.go
```

### License

This project is licensed under the MIT License - see the LICENSE file for details.

#### `Makefile`

```makefile
.PHONY: build test

build:
    go build -o bin/archiver cmd/archiver/main.go

test:
    go test ./...

```
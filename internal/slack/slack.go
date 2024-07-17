package slack

import (
    "github.com/slack-go/slack"
)

type Client struct {
    api *slack.Client
}

func NewClient(token string) *Client {
    return &Client{
        api: slack.New(token),
    }
}

// ReadMessages reads messages from the specified Slack channel
func (c *Client) ReadMessages(channelID string) ([]slack.Message, error) {
    var allMessages []slack.Message
    params := &slack.GetConversationHistoryParameters{
        ChannelID: channelID,
        Limit:     100, // Adjust the limit as needed
    }

    for {
        history, err := c.api.GetConversationHistory(params)
        if err != nil {
            return nil, err
        }

        allMessages = append(allMessages, history.Messages...)

        if !history.HasMore {
            break
        }

        params.Cursor = history.ResponseMetaData.NextCursor
    }

    return allMessages, nil
}

func (c *Client) GetUser(userID string) (*slack.User, error) {
    user, err := c.api.GetUserInfo(userID)
    if err != nil {
        return nil, err
    }
    return user, nil
}
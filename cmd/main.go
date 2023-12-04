package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(".env")
	token := os.Getenv("SLACK_TOKEN")
	userId := os.Getenv("SLACK_USER_ID")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api := slack.New(token, slack.OptionDebug(true))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	user, err := api.GetUserInfo(userId)
	fmt.Println(user)
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}

package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New("xoxe.xoxp-1-Mi0yLTMwNTg1ODM3OTQxNTAtMzA2OTA3Nzk4NTM5Ny02MjkwMDE0NzU4ODM0LTYyNzU0ODE2MTkxNzUtNjM3ZDExNmJmMzY0YmZiOGFlNTNkZjAwY2YzNDViMWJlMzgyYWVhMjdlY2JiZDdjZjY3YmYzMTQxZGI4YTMyZA", slack.OptionDebug(true))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	user, err := api.GetUserInfo("T031QH5PC4E")
	fmt.Println(user)
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}

package main

import (
	events "dondarrion91/github-activity/internal"
	"dondarrion91/github-activity/internal/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 0 || os.Args[1] == "" {
		panic("Missing username as first argument")
	}

	username := os.Args[1]

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	activity, err := utils.GetEvents(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, action := range activity {
		event, err := events.GetEvent(action)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(event)
	}
}

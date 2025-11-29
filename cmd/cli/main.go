package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	custom "github-activity/infraestructure/http"
	model "github-activity/internal/domain"
	usecase "github-activity/internal/usecase/events"
)

const (
	MinArgs = 1
)

func main() {
	if len(os.Args) > MinArgs {
		url := strings.ReplaceAll("https://api.github.com/users/<username>/events", "<username>", os.Args[1])

		client := custom.NewClient(url, custom.Config{})

		resp, err := client.Do()
		if err != nil {
			fmt.Println("err trying to fetch data, usr: %s, err: %v", os.Args[1], err)
			return
		}

		var data []model.GitHubEvent
		if err := json.Unmarshal(resp.Body(), &data); err != nil {
			fmt.Println("cannot unmarshal response body: %v", err)
			return
		}

		eventHandler := usecase.NewEventHandler()

		for _, info := range data {
			eventHandler.Handle(info)
		}
	}
	return
}
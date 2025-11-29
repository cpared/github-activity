package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	MinArgs = 1
)


func main() {
	if len(os.Args) > MinArgs {
		url := strings.ReplaceAll("https://api.github.com/users/<username>/events", "<username>", os.Args[1])
		var data []GitHubEvent
		_, err := resty.New().R().SetResult(&data).Get(url)
		if err != nil {
			fmt.Println("err trying to fetch data, usr: %s, err: %v", os.Args[1], err)
			return
		}

		fmt.Println(data)
	}
	return
}
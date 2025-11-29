package usecase

import (
	"fmt"
	model "github-activity/internal/domain"
)

type PushStrategy struct {}

func (cs *PushStrategy) handle(event model.GitHubEvent) {
	fmt.Printf("- %s has pushed into %s \n", event.Actor.Login, event.Repo.URL)
}
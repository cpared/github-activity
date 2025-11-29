package usecase

import (
	"fmt"
	model "github-activity/internal/domain"
)

type CreateStrategy struct {}

func (cs *CreateStrategy) handle(event model.GitHubEvent) {
	fmt.Printf("- %s has created a new repository, url: %s \n", event.Actor.Login, event.Repo.URL)
}
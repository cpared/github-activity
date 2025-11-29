package usecase

import (
	model "github-activity/internal/domain"
)

type EventStrategy interface {
	handle(event model.GitHubEvent)
}

type EventHandler struct{
	strategies map[string]EventStrategy
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		strategies: map[string]EventStrategy {
			"CreateEvent": &CreateStrategy{},
			"PushEvent": &PushStrategy{},
		},
	}
}

func (e *EventHandler) Handle(event model.GitHubEvent) {
	if strategy, found := e.strategies[event.Type]; found {
		strategy.handle(event)
	}
}
package app

import (
	"errors"

	"github.com/junaid1460/analytics-software-engineer-assignment/packages/actors"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/commits"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/events"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/repositories"
)

type GithubData struct {
	// Data
	actors  []*actors.Actor
	commits []*commits.Commit
	repos   []*repositories.Repository
	events  []*events.Event

	// Indices
	actorByIdIndex map[actors.ID]*actors.Actor
	repoByIdIndex  map[repositories.ID]*repositories.Repository
	eventByIdIndex map[events.ID]*events.Event
}

func NewDB() *GithubData {
	return &GithubData{
		actors:         make([]*actors.Actor, 0),
		commits:        make([]*commits.Commit, 0),
		repos:          make([]*repositories.Repository, 0),
		events:         make([]*events.Event, 0),
		actorByIdIndex: make(map[actors.ID]*actors.Actor),
		repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
		eventByIdIndex: make(map[events.ID]*events.Event),
	}
}

func (database *GithubData) getEvent(id events.ID) (*events.Event, error) {
	event, found := database.eventByIdIndex[id]
	if !found {
		return nil, errors.New("event not found")
	}

	return event, nil
}

func (database *GithubData) getRepoByEvent(event *events.Event) (*repositories.Repository, error) {

	repo, found := database.repoByIdIndex[event.RepoID]
	if !found {
		return nil, errors.New("repo not found")
	}

	return repo, nil
}

func (database *GithubData) getActorByEvent(event *events.Event) (*actors.Actor, error) {

	actor, found := database.actorByIdIndex[event.ActorId]
	if !found {
		return nil, errors.New("repo not found")
	}

	return actor, nil
}

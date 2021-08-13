package app

import (
	"io"
	"os"

	"github.com/junaid1460/analytics-software-engineer-assignment/packages/actors"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/commits"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/events"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/repositories"
)

/**
* * Rebuild indices
* * builds Indices and populates derived fields
 */
func (database *GithubData) BulkRebuildIndices() error {
	// Drop indices
	database.actorByIdIndex = make(map[actors.ID]*actors.Actor)
	database.repoByIdIndex = make(map[repositories.ID]*repositories.Repository)
	database.eventByIdIndex = make(map[events.ID]*events.Event)

	// Repopulate data for indices
	// Process actor info
	for _, actorStruct := range database.actors {
		database.actorByIdIndex[actorStruct.Id] = actorStruct
	}

	// Process repo info
	for _, repo := range database.repos {
		database.repoByIdIndex[repo.Id] = repo
	}

	// Process generic events
	for _, event := range database.events {

		database.eventByIdIndex[event.Id] = event
		switch event.Type {
		case events.WatchEvent:
			repo, err := database.getRepoByEvent(event)
			if err != nil {
				return err
			}
			repo.WatchCount += 1
		case events.PullRequestEvent:
			actor, err := database.getActorByEvent(event)

			if err != nil {
				return err
			}

			actor.PullRequestCount += 1
		}
	}

	// Process commits
	for _, commit := range database.commits {
		// Get event
		event, err := database.getEvent(commit.EventId)

		if err != nil {
			return err
		}

		// Update user commit count
		actor, err := database.getActorByEvent(event)
		if err != nil {
			return err
		}

		actor.CommitCount += 1

		// Update repo commit count
		repo, err := database.getRepoByEvent(event)
		if err != nil {
			return err
		}

		repo.CommitCount += 1
	}
	return nil
}

type BulkReadDataFromFilesInput struct {
	actorsFilePath  string
	reposFilePath   string
	eventsFilePath  string
	commitsFilePath string
}

/**
* Reads CSV files to populate data
* * On success replaces existing data and
* * builds Indices and populates dericed fields
 */
func (database *GithubData) BulkReadDataFromFiles(files BulkReadDataFromFilesInput) error {
	actorsFile := openFile(files.actorsFilePath)
	defer actorsFile.Close()

	repoFile := openFile(files.reposFilePath)
	defer repoFile.Close()

	commitsFile := openFile(files.commitsFilePath)
	defer commitsFile.Close()

	eventsFile := openFile(files.eventsFilePath)
	defer eventsFile.Close()

	return database.bulkReadDataFromOpenFiles(
		actorsFile,
		repoFile,
		commitsFile,
		eventsFile,
	)

}

func openFile(path string) *os.File {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func (database *GithubData) bulkReadDataFromOpenFiles(
	actorsFile io.Reader,
	repoFile io.Reader,
	commitsFile io.Reader,
	eventsFile io.Reader,
) error {
	actorsSlice, err := actors.FromCSV(actorsFile)
	if err != nil {
		return err
	}

	reposSlice, err := repositories.FromCSV(repoFile)
	if err != nil {
		return err
	}

	commitsSlice, err := commits.FromCSV(commitsFile)
	if err != nil {
		return err
	}

	eventsSlice, err := events.FromCSV(eventsFile)
	if err != nil {
		return err
	}

	database.actors = actorsSlice
	database.commits = commitsSlice
	database.events = eventsSlice
	database.repos = reposSlice

	database.BulkRebuildIndices()

	return nil
}

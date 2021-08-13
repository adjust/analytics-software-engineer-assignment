package app

import (
	"fmt"
	"strings"

	"github.com/junaid1460/analytics-software-engineer-assignment/packages/actors"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/repositories"
)

func RunFromFiles(directory string) {
	database := NewDB()

	database.BulkReadDataFromFiles(BulkReadDataFromFilesInput{
		actorsFilePath:  fmt.Sprintf("%s/actors.csv", directory),
		commitsFilePath: fmt.Sprintf("%s/commits.csv", directory),
		reposFilePath:   fmt.Sprintf("%s/repos.csv", directory),
		eventsFilePath:  fmt.Sprintf("%s/events.csv", directory),
	})

	top10UsersByPRandCommits := database.TopUsersByPRsAndCommits(10)
	printTopUsersByPRsAndCommits(top10UsersByPRandCommits)

	top10ReposByCommitCount := database.TopReposByCommitCount(10)
	printTopReposByCommits(top10ReposByCommitCount)

	top10ReposByWatchCount := database.TopReposByWatchCount(10)
	printTopReposByWatchCount(top10ReposByWatchCount)
}

func printTopReposByWatchCount(repos []*repositories.Repository) {
	println("\n • Top 10 repositories sorted by amount of watch events")
	println(strings.Repeat("-", 100))
	println("Watch Count \t\tName")
	println(strings.Repeat("-", 100))
	for _, item := range repos {
		println(item.WatchCount, "\t\t", item.Name)
	}
	println(strings.Repeat("-", 100))
	println(strings.Repeat(" ", 100))
}

func printTopReposByCommits(repos []*repositories.Repository) {
	println("\n • Top 10 repositories sorted by amount of commits pushed")
	println(strings.Repeat("-", 100))
	println("Commits \tName")
	println(strings.Repeat("-", 100))
	for _, item := range repos {
		println(item.CommitCount, "\t\t", item.Name)
	}
	println(strings.Repeat("-", 100))
}

func printTopUsersByPRsAndCommits(actors []*actors.Actor) {
	println("\n • Top 10 active users sorted by amount of PRs created and commits pushed")
	println(strings.Repeat("-", 100))
	println("PRs \t Commits \t Username")
	println(strings.Repeat("-", 100))
	for _, item := range actors {
		println(item.PullRequestCount, "\t", item.CommitCount, "\t", item.Username)
	}
	println(strings.Repeat("-", 100))
}

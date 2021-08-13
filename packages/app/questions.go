package app

import (
	"sort"

	"github.com/junaid1460/analytics-software-engineer-assignment/packages/actors"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/repositories"
)

func (database *GithubData) TopUsersByPRsAndCommits(count int) []*actors.Actor {
	sort.Slice(database.actors, func(p, q int) bool {

		left := database.actors[p]
		right := database.actors[q]

		if right.PullRequestCount < left.PullRequestCount {
			return true
		} else if right.PullRequestCount == left.PullRequestCount && right.CommitCount < left.CommitCount {
			return true
		} else {
			return false
		}
	})
	return database.actors[:count]
}

func (database *GithubData) TopReposByCommitCount(count int) []*repositories.Repository {
	sort.Slice(database.repos, func(p, q int) bool {

		left := database.repos[p]
		right := database.repos[q]

		if right.CommitCount < left.CommitCount {
			return true
		} else {
			return false
		}
	})
	return database.repos[:count]
}

func (database *GithubData) TopReposByWatchCount(count int) []*repositories.Repository {
	sort.Slice(database.repos, func(p, q int) bool {

		left := database.repos[p]
		right := database.repos[q]

		if right.WatchCount < left.WatchCount {
			return true
		} else {
			return false
		}
	})

	return database.repos[:count]
}

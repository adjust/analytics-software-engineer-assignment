package app

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/junaid1460/analytics-software-engineer-assignment/packages/actors"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/commits"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/events"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/repositories"
)

func TestGithubData_TopUsersByPRsAndCommits(t *testing.T) {
	type fields struct {
		actors         []*actors.Actor
		commits        []*commits.Commit
		repos          []*repositories.Repository
		events         []*events.Event
		actorByIdIndex map[actors.ID]*actors.Actor
		repoByIdIndex  map[repositories.ID]*repositories.Repository
		eventByIdIndex map[events.ID]*events.Event
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*actors.Actor
	}{
		{
			name: "Success",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 3, Username: "j9", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 1,
					},
				},
				repos: []*repositories.Repository{
					{
						Id:          20,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
					{
						Id:          21,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
				},
				events: []*events.Event{
					{
						Type:    events.WatchEvent,
						Id:      1,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      2,
						ActorId: 1,
						RepoID:  20,
					},
					{
						Type:    events.DeleteEvent,
						Id:      3,
						ActorId: 1,
						RepoID:  21,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			args: args{count: 1},
			want: []*actors.Actor{{Id: 1, Username: "junaid", PullRequestCount: 1, CommitCount: 1}},
		}, {
			name: "Success: where users are by default not sorted",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 1,
					},
				},
				repos: []*repositories.Repository{
					{
						Id:          20,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
					{
						Id:          21,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
				},
				events: []*events.Event{
					{
						Type:    events.WatchEvent,
						Id:      1,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      2,
						ActorId: 1,
						RepoID:  20,
					},
					{
						Type:    events.PullRequestEvent,
						Id:      2,
						ActorId: 2,
						RepoID:  20,
					},
					{
						Type:    events.DeleteEvent,
						Id:      3,
						ActorId: 1,
						RepoID:  21,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			args: args{count: 1},
			want: []*actors.Actor{{Id: 1, Username: "junaid", PullRequestCount: 1, CommitCount: 1}, {Id: 1, Username: "junaid", PullRequestCount: 1, CommitCount: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := &GithubData{
				actors:         tt.fields.actors,
				commits:        tt.fields.commits,
				repos:          tt.fields.repos,
				events:         tt.fields.events,
				actorByIdIndex: tt.fields.actorByIdIndex,
				repoByIdIndex:  tt.fields.repoByIdIndex,
				eventByIdIndex: tt.fields.eventByIdIndex,
			}
			database.BulkRebuildIndices()

			if got := database.TopUsersByPRsAndCommits(tt.args.count); !reflect.DeepEqual(got, tt.want) {

				println("Expected =>")
				for index, p := range tt.want {
					fmt.Printf("%d: %+v\n", index, p)
				}
				println("Got =>")
				for index, p := range got {
					fmt.Printf("%d: %+v\n", index, p)
				}
			}
		})
	}
}

func TestGithubData_TopReposByCommitCount(t *testing.T) {
	type fields struct {
		actors         []*actors.Actor
		commits        []*commits.Commit
		repos          []*repositories.Repository
		events         []*events.Event
		actorByIdIndex map[actors.ID]*actors.Actor
		repoByIdIndex  map[repositories.ID]*repositories.Repository
		eventByIdIndex map[events.ID]*events.Event
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*repositories.Repository
	}{
		{
			name: "Success",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 3, Username: "j9", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 1,
					},
				},
				repos: []*repositories.Repository{
					{
						Id:          20,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
					{
						Id:          21,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
				},
				events: []*events.Event{
					{
						Type:    events.WatchEvent,
						Id:      1,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      2,
						ActorId: 1,
						RepoID:  20,
					},
					{
						Type:    events.DeleteEvent,
						Id:      3,
						ActorId: 1,
						RepoID:  21,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			args: args{count: 1},
			want: []*repositories.Repository{{
				Id:          20,
				Name:        "repo",
				CommitCount: 1,
				WatchCount:  1,
			}},
		},
		{
			name: "Success: by default repos are not sorted",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 3, Username: "j9", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 2,
					},
				},
				repos: []*repositories.Repository{
					{
						Id:          21,
						Name:        "repo2",
						CommitCount: 0,
						WatchCount:  0,
					},
					{
						Id:          20,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
				},
				events: []*events.Event{
					{
						Type:    events.WatchEvent,
						Id:      1,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      2,
						ActorId: 1,
						RepoID:  20,
					},
					{
						Type:    events.DeleteEvent,
						Id:      3,
						ActorId: 1,
						RepoID:  21,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			args: args{count: 1},
			want: []*repositories.Repository{{
				Id:          20,
				Name:        "repo",
				CommitCount: 1,
				WatchCount:  1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := &GithubData{
				actors:         tt.fields.actors,
				commits:        tt.fields.commits,
				repos:          tt.fields.repos,
				events:         tt.fields.events,
				actorByIdIndex: tt.fields.actorByIdIndex,
				repoByIdIndex:  tt.fields.repoByIdIndex,
				eventByIdIndex: tt.fields.eventByIdIndex,
			}

			database.BulkRebuildIndices()
			if got := database.TopReposByCommitCount(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GithubData.TopReposByCommitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubData_TopReposByWatchCount(t *testing.T) {
	type fields struct {
		actors         []*actors.Actor
		commits        []*commits.Commit
		repos          []*repositories.Repository
		events         []*events.Event
		actorByIdIndex map[actors.ID]*actors.Actor
		repoByIdIndex  map[repositories.ID]*repositories.Repository
		eventByIdIndex map[events.ID]*events.Event
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*repositories.Repository
	}{
		{
			name: "Success",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 3, Username: "j9", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 1,
					},
				},
				repos: []*repositories.Repository{
					{
						Id:          20,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
					{
						Id:          21,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
				},
				events: []*events.Event{
					{
						Type:    events.WatchEvent,
						Id:      1,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      2,
						ActorId: 1,
						RepoID:  20,
					},
					{
						Type:    events.DeleteEvent,
						Id:      3,
						ActorId: 1,
						RepoID:  21,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			args: args{count: 1},
			want: []*repositories.Repository{{
				Id:          20,
				Name:        "repo",
				CommitCount: 1,
				WatchCount:  1,
			}},
		},
		{
			name: "Success: by default repos are not sorted",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 3, Username: "j9", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 1,
					},
				},
				repos: []*repositories.Repository{

					{
						Id:          21,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
					{
						Id:          20,
						Name:        "repo",
						CommitCount: 0,
						WatchCount:  0,
					},
				},
				events: []*events.Event{
					{
						Type:    events.WatchEvent,
						Id:      1,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      2,
						ActorId: 1,
						RepoID:  20,
					},
					{
						Type:    events.DeleteEvent,
						Id:      3,
						ActorId: 1,
						RepoID:  21,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			args: args{count: 1},
			want: []*repositories.Repository{{
				Id:          20,
				Name:        "repo",
				CommitCount: 1,
				WatchCount:  1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := &GithubData{
				actors:         tt.fields.actors,
				commits:        tt.fields.commits,
				repos:          tt.fields.repos,
				events:         tt.fields.events,
				actorByIdIndex: tt.fields.actorByIdIndex,
				repoByIdIndex:  tt.fields.repoByIdIndex,
				eventByIdIndex: tt.fields.eventByIdIndex,
			}

			database.BulkRebuildIndices()
			if got := database.TopReposByWatchCount(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GithubData.TopReposByWatchCount() = %v, want %v", got[0], tt.want[0])
			}
		})
	}
}

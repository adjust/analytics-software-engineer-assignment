package app

import (
	"io"
	"strings"
	"testing"

	"github.com/junaid1460/analytics-software-engineer-assignment/packages/actors"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/commits"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/events"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/repositories"
)

func TestGithubData_BulkRebuildIndices(t *testing.T) {
	type fields struct {
		actors         []*actors.Actor
		commits        []*commits.Commit
		repos          []*repositories.Repository
		events         []*events.Event
		actorByIdIndex map[actors.ID]*actors.Actor
		repoByIdIndex  map[repositories.ID]*repositories.Repository
		eventByIdIndex map[events.ID]*events.Event
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 10,
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
						Id:      10,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      10,
						ActorId: 1,
						RepoID:  20,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			wantErr: false,
		},
		{
			name: "Repo not found while processing events",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 10,
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
						Id:      10,
						ActorId: 1,
						RepoID:  23,
					}, {
						Type:    events.PullRequestEvent,
						Id:      10,
						ActorId: 1,
						RepoID:  20,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			wantErr: true,
		},
		{
			name: "Actor not found while processing events",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
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
						Id:      10,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      10,
						ActorId: 12,
						RepoID:  20,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			wantErr: true,
		},
		{
			name: "Event not found while processing commits",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 103,
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
						Id:      10,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      102,
						ActorId: 1,
						RepoID:  20,
					},
					{
						Type:    events.DeleteEvent,
						Id:      103,
						ActorId: 1,
						RepoID:  202,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			wantErr: true,
		},
		{
			name: "Event not found while processing commits",
			fields: fields{
				actors: []*actors.Actor{
					{Id: 1, Username: "junaid", PullRequestCount: 0, CommitCount: 0},
					{Id: 2, Username: "blah", PullRequestCount: 0, CommitCount: 0},
				},
				commits: []*commits.Commit{
					{
						SHA:     "121312",
						Message: "test",
						EventId: 108,
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
						Id:      10,
						ActorId: 1,
						RepoID:  20,
					}, {
						Type:    events.PullRequestEvent,
						Id:      102,
						ActorId: 1,
						RepoID:  20,
					},
					{
						Type:    events.DeleteEvent,
						Id:      103,
						ActorId: 1,
						RepoID:  202,
					},
				},
				actorByIdIndex: make(map[actors.ID]*actors.Actor),
				repoByIdIndex:  make(map[repositories.ID]*repositories.Repository),
				eventByIdIndex: make(map[events.ID]*events.Event),
			},
			wantErr: true,
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
			if err := database.BulkRebuildIndices(); (err != nil) != tt.wantErr {
				t.Errorf("GithubData.BulkRebuildIndices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGithubData_bulkReadDataFromOpenFiles(t *testing.T) {

	type args struct {
		actorsFile  io.Reader
		repoFile    io.Reader
		commitsFile io.Reader
		eventsFile  io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				repoFile:    strings.NewReader("id,name\n13,junaid"),
				eventsFile:  strings.NewReader("type,id,actor_id,repo_id\nPublicEvent,12,13,14"),
				commitsFile: strings.NewReader("sha,message,event_id\njunaid,ok,12"),
				actorsFile:  strings.NewReader("id,username\n12,junaid"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := NewDB()
			if err := database.bulkReadDataFromOpenFiles(tt.args.actorsFile, tt.args.repoFile, tt.args.commitsFile, tt.args.eventsFile); (err != nil) != tt.wantErr {
				t.Errorf("GithubData.bulkReadDataFromOpenFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

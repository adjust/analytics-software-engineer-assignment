package events

import (
	"encoding/csv"
	"errors"
	"io"

	"github.com/jszwec/csvutil"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/actors"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/repositories"
)

const (
	CommitCommentEvent            = 10
	CreateEvent                   = 20
	DeleteEvent                   = 30
	ForkEvent                     = 50
	GollumEvent                   = 60
	IssueCommentEvent             = 70
	IssuesEvent                   = 80
	MemberEvent                   = 90
	PublicEvent                   = 100
	PullRequestEvent              = 110
	PullRequestReviewCommentEvent = 120
	PushEvent                     = 130
	ReleaseEvent                  = 140
	WatchEvent                    = 150
)

var eventTypes = map[string]int{
	"CommitCommentEvent":            CommitCommentEvent,
	"CreateEvent":                   CreateEvent,
	"DeleteEvent":                   DeleteEvent,
	"ForkEvent":                     ForkEvent,
	"GollumEvent":                   GollumEvent,
	"IssueCommentEvent":             IssueCommentEvent,
	"IssuesEvent":                   IssuesEvent,
	"MemberEvent":                   MemberEvent,
	"PublicEvent":                   PublicEvent,
	"PullRequestEvent":              PullRequestEvent,
	"PullRequestReviewCommentEvent": PullRequestReviewCommentEvent,
	"PushEvent":                     PushEvent,
	"ReleaseEvent":                  ReleaseEvent,
	"WatchEvent":                    WatchEvent,
}

type EventType int
type ID int64
type Event struct {
	Type       EventType
	TypeString string          `csv:"type"`
	Id         ID              `csv:"id"`
	ActorId    actors.ID       `csv:"actor_id"`
	RepoID     repositories.ID `csv:"repo_id"`
}

func FromCSV(file io.Reader) ([]*Event, error) {
	csvReader := csv.NewReader(file)

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, err
	}

	var events []*Event
	for {
		var event Event
		if err := dec.Decode(&event); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	for _, event := range events {
		eventType, found := eventTypes[event.TypeString]

		if !found {
			return nil, errors.New("invalid event type")
		}
		event.Type = EventType(eventType)

	}
	return events, nil
}

package commits

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
	"github.com/junaid1460/analytics-software-engineer-assignment/packages/events"
)

type Commit struct {
	SHA     string    `csv:"sha"`
	Message string    `csv:"message"`
	EventId events.ID `csv:"event_id"`
}

func FromCSV(file io.Reader) ([]*Commit, error) {
	csvReader := csv.NewReader(file)

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, err
	}

	var users []*Commit
	for {
		var user Commit
		if err := dec.Decode(&user); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

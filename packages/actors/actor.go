package actors

import (
	"encoding/csv"
	"io"

	csvutil "github.com/jszwec/csvutil"
)

type ID int64
type Actor struct {
	Id               ID     `csv:"id"`
	Username         string `csv:"username"`
	PullRequestCount int
	CommitCount      int
}

func FromCSV(file io.Reader) ([]*Actor, error) {
	csvReader := csv.NewReader(file)

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, err
	}

	var users []*Actor
	for {
		var user Actor
		if err := dec.Decode(&user); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

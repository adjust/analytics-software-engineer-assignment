package repositories

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

type ID int64
type Repository struct {
	Id          ID     `csv:"id"`
	Name        string `csv:"name"`
	CommitCount int
	WatchCount  int
}

func FromCSV(file io.Reader) ([]*Repository, error) {
	csvReader := csv.NewReader(file)

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, err
	}

	var users []*Repository
	for {
		var user Repository
		if err := dec.Decode(&user); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

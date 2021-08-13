package commits

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestFromCSV(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []*Commit
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				file: strings.NewReader("sha,message,event_id\njunaid,ok,12"),
			},
			want:    []*Commit{{SHA: "junaid", Message: "ok", EventId: 12}},
			wantErr: false,
		},
		{
			name: "IO reader failed",
			args: args{
				file: errReader(1),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error while reading",
			args: args{
				file: strings.NewReader("id,username\n\n"),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromCSV(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
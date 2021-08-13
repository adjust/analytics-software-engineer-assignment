package actors

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/pkg/errors"
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
		want    []*Actor
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				file: strings.NewReader("id,username\n12,junaid"),
			},
			want:    []*Actor{{Id: 12, Username: "junaid"}},
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

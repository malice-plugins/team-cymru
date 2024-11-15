package main

import (
	"reflect"
	"testing"
)

func TestLookupHash(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name string
		args args
		want ResultsData
	}{
		{
			name: "Test LookupHash",
			args: args{
				hash: "1d31bd48b2e864c773ca6a3b9fd0019416809066",
			},
			want: ResultsData{
				Found:     true,
				LastSeen:  "2017-01-12",
				Detection: "86%",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LookupHash(tt.args.hash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LookupHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

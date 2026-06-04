package main

import (
	"reflect"
	"testing"
)

func Test_getResults(t *testing.T) {
	tests := []struct {
		name    string
		queries [][]int
		want    []bool
	}{
		{
			name:    "test case one",
			queries: [][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}},
			want:    []bool{false, true, true},
		}, {
			name:    "test case two",
			queries: [][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}},
			want:    []bool{true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getResults(tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

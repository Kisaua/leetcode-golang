package main

import (
	"reflect"
	"testing"
)

func Test_rotateTheBox(t *testing.T) {
	tests := []struct {
		name    string
		boxGrid [][]byte
		want    [][]byte
	}{
		{
			name:    "test case one",
			boxGrid: [][]byte{{'#', '.', '#'}},
			want: [][]byte{
				{'.'},
				{'#'},
				{'#'},
			},
		}, {
			name: "test case two",
			boxGrid: [][]byte{
				{'#', '.', '*', '.'},
				{'#', '#', '*', '.'},
			},
			want: [][]byte{
				{'#', '.'},
				{'#', '#'},
				{'*', '*'},
				{'.', '.'},
			},
		}, {
			name: "test case three",
			boxGrid: [][]byte{
				{'#', '#', '*', '.', '*', '.'},
				{'#', '#', '#', '*', '.', '.'},
				{'#', '#', '#', '.', '#', '.'},
			},
			want: [][]byte{
				{'.', '#', '#'},
				{'.', '#', '#'},
				{'#', '#', '*'},
				{'#', '*', '.'},
				{'#', '.', '*'},
				{'#', '.', '.'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateTheBox(tt.boxGrid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateTheBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_stringIndices(t *testing.T) {
	type args struct {
		wordsContainer []string
		wordsQuery     []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case one",
			args: args{
				wordsContainer: []string{"abcd", "bcd", "xbcd"},
				wordsQuery:     []string{"cd", "bcd", "xyz"},
			},
			want: []int{1, 1, 1},
		}, {
			name: "test case two",
			args: args{
				wordsContainer: []string{"abcdefgh", "poiuygh", "ghghgh"},
				wordsQuery:     []string{"gh", "acbfgh", "acbfegh"},
			},
			want: []int{2, 0, 2},
		}, {
			name: "test case three",
			args: args{
				wordsContainer: []string{"akfgaffgkg", "fagagu", "fuagkkkfu"},
				wordsQuery:     []string{"gagffu", "fgufkukauf", "fafgggk", "gkggufuag"},
			},
			want: []int{2, 1, 1, 0},
		}, {
			name: "test case four",
			args: args{
				wordsContainer: []string{"abcde", "abcde"},
				wordsQuery:     []string{"abcde", "bcde", "cde", "de", "e"},
			},
			want: []int{0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringIndices(tt.args.wordsContainer, tt.args.wordsQuery); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}

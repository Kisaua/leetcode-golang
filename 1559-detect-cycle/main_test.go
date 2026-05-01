package main

import "testing"

func Test_containsCycle(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want bool
	}{
		{
			name: "test case one",
			grid: [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}},
			want: true,
		}, {
			name: "test case two",
			grid: [][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}},
			want: true,
		}, {
			name: "test case three",
			grid: [][]byte{{'a', 'b', 'b'}, {'b', 'z', 'b'}, {'b', 'b', 'a'}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsCycle(tt.grid); got != tt.want {
				t.Errorf("containsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

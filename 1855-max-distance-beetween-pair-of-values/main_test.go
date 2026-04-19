package main

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case one",
			args: args{
				nums1: []int{55, 30, 5, 4, 2},
				nums2: []int{100, 20, 10, 10, 5},
			},
			want: 2,
		}, {
			name: "test case two",
			args: args{
				nums1: []int{2, 2, 2},
				nums2: []int{10, 10, 1},
			},
			want: 1,
		}, {
			name: "test case three",
			args: args{
				nums1: []int{30, 29, 19, 5},
				nums2: []int{25, 25, 25, 25, 25},
			},
			want: 2,
		}, {
			name: "test case four",
			args: args{
				nums1: []int{9820, 8937, 7936, 4855, 4830, 4122, 2327, 1342, 1167, 815, 414},
				nums2: []int{9889, 9817, 9800, 9777, 9670, 9646, 9304, 8977, 8974, 8802, 8626, 8622, 8456},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

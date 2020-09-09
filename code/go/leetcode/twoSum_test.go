package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "0,2", args: args{nums: []int{2, 3, 4}, target: 6}, want: []int{0, 2}},
		{name: "0,1", args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
		{name: "1,2", args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

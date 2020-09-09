package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 4, Next: n3}
	n1 := &ListNode{Val: 2, Next: n2}
	n6 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 6, Next: n6}
	n4 := &ListNode{Val: 5, Next: n5}
	n9 := &ListNode{Val: 8}
	n8 := &ListNode{Val: 0, Next: n9}
	n7 := &ListNode{Val: 7, Next: n8}
	n55 := &ListNode{Val: 5}
	n552 := &ListNode{Val: 5}
	n11 := &ListNode{Val: 1}
	n10 := &ListNode{Val: 0, Next: n11}
	n188 := &ListNode{Val: 8}
	n18 := &ListNode{Val: 1, Next: n188}
	n0 := &ListNode{Val: 0}
	n1a := &ListNode{Val: 1}
	n99 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	n100 := &ListNode{Val: 0, Next: &ListNode{Val: 1}}
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "807", args: args{l1: n1, l2: n4}, want: n7},
		{name: "55", args: args{l1: n55, l2: n552}, want: n10},
		{name: "18", args: args{l1: n18, l2: n0}, want: n18},
		{name: "1,99", args: args{l1: n1a, l2: n99}, want: &ListNode{Val: 0, Next: n100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			if !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("addTwoNumbers() = %v, want %v", got.Val, tt.want.Val)
			}
			if !reflect.DeepEqual(got.Next.Val, tt.want.Next.Val) {
				t.Errorf("addTwoNumbers() = %v, want %v", got.Next.Val, tt.want.Next.Val)
			}
		})
	}
}

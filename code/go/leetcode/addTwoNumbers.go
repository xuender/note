package leetcode

// ListNode list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	result := root
	flag := 0
	for l1 != nil || l2 != nil || flag > 0 {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		flag = (v1 + v2 + flag) / 10
		result.Next = &ListNode{Val: (v1 + v2 + flag) % 10}
		result = result.Next
	}
	return root.Next
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	result := root
	flag := 0
	for l1 != nil || l2 != nil || flag > 0 {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + flag
		flag = sum / 10
		result.Next = &ListNode{Val: sum % 10}
		result = result.Next
	}
	return root.Next
}
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	l2.Val += l1.Val
	up(l2)
	if l1.Next != nil {
		if l2.Next == nil {
			l2.Next = &ListNode{Val: 0}
		}
		addTwoNumbers1(l1.Next, l2.Next)
	}
	return l2
}

func up(l *ListNode) {
	if l == nil {
		return
	}
	if l.Val > 9 {
		if l.Next == nil {
			l.Next = &ListNode{Val: 0}
		}
		l.Next.Val++
		l.Val -= 10
		up(l.Next)
	}
}

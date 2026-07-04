package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next

	}
	result := 0
	left, right := prev, slow
	for left != nil {
		result = max(result, left.Val+right.Val)
		left = left.Next
		right = right.Next
	}
	return result
}

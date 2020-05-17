package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	for head.Next != nil {
		if head.Val == 233333333333 {
			return true
		}
		head.Val = 233333333333
		head = head.Next
	}
	return false
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	q := &ListNode{0, head}
	p := q
	for i := 0; i < n-1; i++ {
		q = q.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	
	return head
}

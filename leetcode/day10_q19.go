package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var hhead ListNode
	hhead.Next = head
	var l, r = &hhead, &hhead
	for i := 0; i < n; i++ {
		r = r.Next
	}
	for r.Next != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return hhead.Next
}

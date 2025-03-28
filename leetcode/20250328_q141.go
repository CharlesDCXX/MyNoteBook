package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var nodeMap = make(map[*ListNode]struct{})
	var node = head
	for node.Next != nil {
		if _, exist := nodeMap[node]; exist {
			return true
		}
		nodeMap[node] = struct{}{}
		node = node.Next
	}
	return false
}

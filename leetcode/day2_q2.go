package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*  非递归

	func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		var Result ListNode
		var tl = &Result
		var temp = 0
		var erwei = 0
		for l1 != nil || l2 != nil || erwei != 0 {
			if l1 != nil {
				temp += l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				temp += l2.Val
				l2 = l2.Next
			}
			temp += erwei
			erwei = temp / 10
			temp %= 10
			tl.Next = &ListNode{Val: temp}
			tl = tl.Next
			temp = 0
		}
		return Result.Next
	}

*
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return digui(l1, l2, 0)
}
func digui(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		}
		return &ListNode{Val: carry}
	}
	if l1 == nil {
		l1, l2 = l2, l1
	}
	sum := carry + l1.Val
	if l2 != nil {
		sum = carry + l2.Val
		l2 = l2.Next
	}
	carry = sum / 10
	sum = sum % 10
	l1.Val = sum % 10                    // 每个节点保存一个数位
	l1.Next = digui(l1.Next, l2, sum/10) // 进位
	return l1
}
func main1() {
	l1 := &ListNode{9, &ListNode{9, &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}
	l2 := &ListNode{9, &ListNode{9, &ListNode{Val: 9}}}
	re := addTwoNumbers(l1, l2)
	fmt.Println(re.Val, re.Next, re.Next.Next)
}

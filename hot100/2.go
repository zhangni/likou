package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{-1, nil}
	var next = head
	var sum = 0
	var jinwei = 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			sum = l2.Val + jinwei
			l2 = l2.Next
		} else if l2 == nil {
			sum = l1.Val + jinwei
			l1 = l1.Next
		} else {
			sum = l1.Val + l2.Val + jinwei
			l1 = l1.Next
			l2 = l2.Next
		}
		jinwei = sum / 10
		sum = sum % 10
		next.Next = &ListNode{sum, nil}
		next = next.Next
	}
	if jinwei != 0 {
		next.Next = &ListNode{jinwei, nil}
	}
	return head.Next

}

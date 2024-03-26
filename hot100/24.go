package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := new(ListNode)
	node.Next = head
	first, second := node, head
	for second.Next != nil && first.Next != nil {
		first.Next = second.Next
		first = first.Next
		second.Next = first.Next
		first.Next = second
		if second.Next == nil {
			return node.Next
		}
		first = first.Next
		second = second.Next
	}
	return node.Next
}

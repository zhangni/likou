package hot100

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var fast = head.Next
	var slow = head
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	head2 := slow.Next
	slow.Next = nil
	var left = sortList(head)
	var right = sortList(head2)
	var newHead = &ListNode{-1, nil}
	var res = newHead
	for {
		if left == nil {
			newHead.Next = right
			break
		}
		if right == nil {
			newHead.Next = left
			break
		}
		if left.Val > right.Val {
			newHead.Next = right
			right = right.Next
		} else {
			newHead.Next = left
			left = left.Next
		}
		newHead = newHead.Next
	}
	return res.Next
}

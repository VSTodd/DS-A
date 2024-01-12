func deleteDuplicates(head *ListNode) *ListNode {
	decoy := &ListNode{
		Val:  -101,
		Next: head,
	}

	if head == nil {
		return head
	}

	prev := decoy
	curr := head
	next := curr.Next
	for next != nil {
		if curr.Val == next.Val {
			for curr != nil && next != nil && curr.Val == next.Val {
				next = next.Next
			}
			prev.Next = next
		} else {
			prev = prev.Next
		}
		curr = next
		if curr == nil {
			break
		}

		next = curr.Next
	}

	return decoy.Next
}
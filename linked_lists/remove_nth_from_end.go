func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	sentinel := &ListNode{Next: head}
	slower := sentinel

	for i := 1; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		slower = slower.Next
	}

	slower.Next = slow.Next

	return sentinel.Next
}
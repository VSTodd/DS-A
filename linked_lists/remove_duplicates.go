package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	next := curr.Next
	for next != nil {
		if curr.Val == next.Val {
			next = next.Next
		} else {
			curr.Next = next
			curr = next
			next = curr.Next
		}
	}
	curr.Next = nil
	return head
}

func main() {

}

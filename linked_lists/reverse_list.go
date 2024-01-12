/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode
	prev = nil
	curr := head
	next := head.Next

	for next != nil {
		curr.Next = prev
		prev = curr
		curr = next
		next = next.Next
	}

	curr.Next = prev

	return curr
}
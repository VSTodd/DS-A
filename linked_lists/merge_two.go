package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var currentMerged *ListNode
	var head *ListNode
	currList1 := list1
	currList2 := list2

	if currList1 == nil && currList2 == nil {
		return nil
	} else if currList1 == nil {
		return currList2
	} else if currList2 == nil {
		return currList1
	} else if currList1.Val > currList2.Val {
		head = currList2
		currentMerged = currList2
		currList2 = currList2.Next
	} else {
		head = currList1
		currentMerged = currList1
		currList1 = currList1.Next
	}

	for {
		if currList1 == nil {
			currentMerged.Next = currList2
			break
		} else if currList2 == nil {
			currentMerged.Next = currList1
			break
		} else if currList1.Val < currList2.Val {
			currentMerged.Next = currList1
			currentMerged = currentMerged.Next
			currList1 = currList1.Next
		} else {
			currentMerged.Next = currList2
			currentMerged = currentMerged.Next
			currList2 = currList2.Next
		}
	}

	return head
}

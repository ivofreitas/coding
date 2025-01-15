package main

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	right := head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	dummy := &ListNode{Next: head}
	left := dummy
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next

	return dummy.Next
}

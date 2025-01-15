package main

func hasCycle(head *ListNode) bool {

	slowPointer := head
	fastPointer := head

	for fastPointer != nil && fastPointer.Next != nil {

		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

		if fastPointer == slowPointer {
			return true
		}
	}

	return false
}

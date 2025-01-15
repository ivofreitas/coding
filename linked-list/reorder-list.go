package main

import "fmt"

func ReorderList(head *ListNode) {
	slowPointer, fastPointer := head, head.Next

	/*
		move slow pointer to the half of the list
	*/
	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	/*
		inverse second part
	*/
	secondHalf := slowPointer.Next
	slowPointer.Next = nil
	var prev *ListNode
	for secondHalf != nil {
		tmp := secondHalf.Next
		secondHalf.Next = prev
		prev = secondHalf
		secondHalf = tmp
	}

	/*
		reorder list
	*/

	firstHalf, secondHalf := head, prev
	for secondHalf != nil {
		tmp1, tmp2 := firstHalf.Next, secondHalf.Next
		firstHalf.Next = secondHalf
		secondHalf.Next = tmp1
		firstHalf, secondHalf = tmp1, tmp2
	}

	fmt.Println("a")
}

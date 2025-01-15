package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := new(ListNode)
	node1.Val = 2
	node2 := new(ListNode)
	node2.Val = 4
	node1.Next = node2
	node3 := new(ListNode)
	node3.Val = 3
	node2.Next = node3

	node21 := new(ListNode)
	node21.Val = 5
	node22 := new(ListNode)
	node22.Val = 6
	node21.Next = node22
	node23 := new(ListNode)
	node23.Val = 4
	node22.Next = node23

	printList(addTwoNumbers(node1, node21))
}

func printList(list *ListNode) {
	s := make([]string, 0)
	for list != nil {
		s = append(s, strconv.Itoa(list.Val))
		list = list.Next
	}
	fmt.Println(strings.Join(s, " -> "))
}

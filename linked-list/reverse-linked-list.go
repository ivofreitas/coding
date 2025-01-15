package main

type LinkedList struct {
	head *ListNode
}

func (l *LinkedList) Insert(val int) {
	if l.head == nil {
		node := &ListNode{Val: val}
		l.head = node
		return
	}

	node := &ListNode{Val: val, Next: l.head}
	l.head = node
}

func ReverseList(head *ListNode) *ListNode {
	reverse := new(LinkedList)

	current := head
	for current != nil {
		reverse.Insert(current.Val)
		current = current.Next
	}

	return reverse.head
}

package twopointers

import "fmt"

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func (this *EduLinkedListNode) String() string {
	aux := this
	var list string
	for aux != nil {
		list += fmt.Sprintf("%v", aux.data)
		aux = aux.next
	}
	return list
}

func removeNthLastNode(head *EduLinkedListNode, n int) *EduLinkedListNode {
	if head == nil {
		return nil
	}
	left, right := head, head
	for i := 0; i < n; i++ {
		if right == nil {
			return nil
		}
		right = right.next
	}
	if right == nil {
		head = head.next
		return head
	}
	for right != nil && right.next != nil {
		right = right.next
		left = left.next
	}
	left.next = left.next.next
	return head
}

package twopointers

import "testing"

func TestRemoveNthLastNode(t *testing.T) {
	node := &EduLinkedListNode{data: 69}
	node.next = &EduLinkedListNode{data: 8}
	node.next.next = &EduLinkedListNode{data: 49}
	node.next.next.next = &EduLinkedListNode{data: 106}
	node.next.next.next.next = &EduLinkedListNode{data: 116}
	node.next.next.next.next.next = &EduLinkedListNode{data: 112}
	want := "849106116112"
	node = removeNthLastNode(node, 6)
	if got := node.String(); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

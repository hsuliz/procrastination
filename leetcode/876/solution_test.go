package _876

import (
	"testing"
)

func TestMiddleNode(t *testing.T) {
	list5 := ListNode{4, nil}
	list4 := ListNode{4, &list5}
	list3 := ListNode{3, &list4}
	list2 := ListNode{2, &list3}
	list1 := ListNode{1, &list2}
	got := *middleNode(&list1)
	if got != list3 {
		t.Errorf("got: %v, want %v", got, list3)
	}
}

package _206

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	list5 := ListNode{5, nil}
	list4 := ListNode{4, &list5}
	list3 := ListNode{3, &list4}
	list2 := ListNode{2, &list3}
	list1 := ListNode{1, &list2}

	got := reverseList(&list1)
	for got != nil {
		fmt.Println(got)
		got = got.Next
	}
}

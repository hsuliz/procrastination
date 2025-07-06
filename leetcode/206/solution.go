package _206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	tmp := head
	var prev *ListNode = nil
	var next *ListNode = nil
	for tmp != nil {
		next = tmp.Next
		tmp.Next = prev
		prev = tmp
		tmp = next
	}
	return prev
}

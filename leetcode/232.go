package leetcode

type MyQueue struct {
	Store []int
}

func Constructor() MyQueue {
	return MyQueue{make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.Store = append(this.Store, x)
}

func (this *MyQueue) Pop() int {
	res := this.Store[0]
	this.Store = this.Store[1:]
	return res
}

func (this *MyQueue) Peek() int {
	return this.Store[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.Store) == 0
}

package _102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var currRes []int
		queueLen := len(queue)
		for range queueLen {
			currNode := queue[0]
			queue = queue[1:]
			if currNode != nil {
				queue = append(queue, currNode.Left)
				queue = append(queue, currNode.Right)
				currRes = append(currRes, currNode.Val)
			}
		}
		if len(currRes) != 0 {
			res = append(res, currRes)
		}
	}

	return res
}

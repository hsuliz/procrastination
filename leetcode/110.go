package leetcode

import "math"

func isBalanced(root *TreeNode) bool {
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftDepth := depth(root.Left)
		if leftDepth == -1 {
			return -1
		}

		rightDepth := depth(root.Right)
		if rightDepth == -1 {
			return -1
		}

		if math.Abs(float64(rightDepth-leftDepth)) > 1 {
			return -1
		}

		return max(leftDepth, rightDepth) + 1
	}
	return depth(root) >= 0
}

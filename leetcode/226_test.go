package leetcode

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	/*tree := TreeNode{
		4,
		&TreeNode{
			2,
			&TreeNode{1, nil, nil},
			&TreeNode{3, nil, nil},
		},
		&TreeNode{
			7,
			&TreeNode{6, nil, nil},
			&TreeNode{9, nil, nil},
		},
	}*/

	tree := TreeNode{
		2,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	fmt.Printf("%v<-%v->%v\n", tree.Left.Val, tree.Val, tree.Right.Val)
	res := invertTree(&tree)
	fmt.Printf("%v<-%v->%v\n", res.Left.Val, res.Val, res.Right.Val)
}

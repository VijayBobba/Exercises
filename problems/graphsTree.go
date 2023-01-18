package problems

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GnTCaller() {
	var one, two, root *TreeNode
	one = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: two,
	}
	two = &TreeNode{
		Val:   2,
		Left:  one,
		Right: nil,
	}
	//[1, 3, 1, 2]
	root = &TreeNode{3, one, two}
	total := countNodes(root)
	fmt.Println(total)
	//[2, 1, 3, 1]
	root = invertTree(root)
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

func countNodes(root *TreeNode) int {
	if root != nil {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	} else {
		return 0
	}
}

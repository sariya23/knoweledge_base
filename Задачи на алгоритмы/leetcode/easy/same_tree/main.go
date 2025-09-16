package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var nodes1, nodes2 []int
	walkTreeAndCollectNodeValues(p, &nodes1)
	walkTreeAndCollectNodeValues(q, &nodes2)

	if len(nodes1) != len(nodes2) {
		return false
	}

	for i := 0; i < len(nodes1); i++ {
		if nodes1[i] != nodes2[i] {
			return false
		}
	}
	return true
}

func walkTreeAndCollectNodeValues(node *TreeNode, nodes *[]int) {
	if node == nil {
		*nodes = append(*nodes, 10_000_000)
		return
	}
	*nodes = append(*nodes, node.Val)
	walkTreeAndCollectNodeValues(node.Left, nodes)
	walkTreeAndCollectNodeValues(node.Right, nodes)
}

func walkTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(root.Val)
	walkTree(root.Left)
	walkTree(root.Right)
}

func main() {
	root1 := TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, nil},
		Right: nil,
	}

	root2 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{2, nil, nil},
	}

	fmt.Println(isSameTree(&root1, &root2))

	walkTree(&root1)
	fmt.Println()
	walkTree(&root2)
}

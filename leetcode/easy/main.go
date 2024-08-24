package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// func PrintPaths(root Node) {
// 	var path []int
// 	printHelper(root, path, 0)
// }

// func printHelper(node Node, path []int, pathLen int) {
// 	if node
// }

func insert(t *Node, v int) *Node {
	if t == nil {
		return &Node{v, nil, nil}
	}
	if v < t.Data {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

var steps []int

func getSteps(n int) []int {
	if n <= 1 {
		steps = append(steps, 1)
		return steps
	}
	getSteps(n - 1)
	getSteps(n - 2)
	return nil
}

func main() {
	fmt.Println(getSteps(4))
}

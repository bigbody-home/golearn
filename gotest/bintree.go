package gotest

import "fmt"

type tree struct {
	value int
	left  *tree
	right *tree
}

func PrintTree() {
	slicetree := []int{6, 7, 8, 2, 1, 4, 5, 10}
	root := GetTree(slicetree)
	//fmt.Println(root.value)
	//fmt.Println(root.left.value)
	//fmt.Println(root.right.value)
	middleLoop(root)

}
func middleLoop(tree *tree) {
	if tree != nil {
		fmt.Println(tree.value)
		middleLoop(tree.left)
		middleLoop(tree.right)
	}

}
func GetTree(n []int) *tree {
	var root *tree
	for _, v := range n {
		root = buildTree(root, v)
	}
	return root
}
func buildTree(root *tree, num int) *tree {
	if root == nil {
		root = new(tree)
		root.value = num
		return root
	}
	if root.value > num {
		root.left = buildTree(root.left, num)
	} else {
		root.right = buildTree(root.right, num)

	}
	return root

}

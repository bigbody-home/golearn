package bintree

import "fmt"

type Container interface {
	Put(value int)
	Get(k int) int
	Delete(k int)
	Contains(k int) bool
	IsEmpty() bool
	Size() int
	Max() int
	Min() int
}
type Node struct {
	val   int
	left  *Node
	right *Node
}
type BinTree struct {
	root *Node
	size int
	min  *Node
	max  *Node
}

func (b *BinTree) Put(value int) {
	var lastnode *Node
	if b.root == nil {
		n := new(Node)
		n.val = value
		b.size++
		b.max = n
		b.min = n
		b.root = n
		return
	}
	cur_node := b.root
	new_node := new(Node)
	new_node.val = value
	for cur_node != nil {
		lastnode = cur_node

		if cur_node.val > new_node.val {

			cur_node = cur_node.left

		} else if cur_node.val < new_node.val {

			cur_node = cur_node.right
		}

	}
	if lastnode.val > new_node.val {
		lastnode.left = new_node
		b.size++
		return
	}
	if lastnode.val < new_node.val {
		lastnode.right = new_node
		b.size++
	}

}

func (b *BinTree) Get(k int) int {
	var n *Node
	n = b.root
	for n != nil {
		fmt.Println(n)
		if n.val == k {
			return n.val

		}
		if n.val > k {
			n = n.left
			continue
		}
		if n.val < k {
			n = n.right
			continue
		}

	}
	return -1
}

func (b *BinTree) Delete(k int) {
	//TODO implement me
	panic("implement me")
}

func (b *BinTree) Contains(k int) bool {
	//TODO implement me
	if b.Get(k) > 0 {
		return true
	}
	return false
}

func (b *BinTree) IsEmpty() bool {
	if b.Size() > 0 {
		return false
	}
	return true
}

func (b *BinTree) Size() int {
	return b.size
}

func (b *BinTree) Max() int {
	maxNode := b.root
	var pre *Node
	for maxNode != nil {
		pre = maxNode
		maxNode = maxNode.right

	}
	return pre.val
}

func (b *BinTree) Min() int {
	minNode := b.root
	var pre *Node
	for minNode != nil {
		pre = minNode
		minNode = minNode.left
	}
	return pre.val
}

package main

import (
	"fmt"
)

type Bag struct {
	root  *Node
	first *Node
	tail  *Node
	len   uint
}

func NewBag() *Bag {

	return &Bag{}
}

type Node struct {
	val  uint
	next *Node
}

func (b *Bag) Add(n *Node) {
	if b.root == nil {
		b.root = n
		b.first = n
		b.tail = n
		b.len++
		return
	}
	tail := b.tail
	tail.next = n
	b.tail = n
	b.len++
}
func (b *Bag) PrintNode() {
	cur := b.root
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

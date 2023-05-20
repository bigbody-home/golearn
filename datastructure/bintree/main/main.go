package main

import (
	"fmt"
	"golearn/datastructure/bintree"
)

func main() {
	tree := new(bintree.BinTree)
	tree.Put(4)
	tree.Put(3)
	tree.Put(5)
	tree.Put(42)
	tree.Put(31)
	tree.Put(52)
	fmt.Println(tree.Size(), tree.Get(5), tree.Max(), tree.Min())
}

package main

import (
	"fmt"
	"sort"
)

type Node struct {
	val  interface{}
	next *Node
}

func NewNode(val interface{}, next *Node) *Node {
	return &Node{val: val, next: next}
}

type Link struct {
	first *Node
	last  *Node
	root  *Node
	size  int
}

func NewLink() *Link {
	return &Link{}
}

func (l *Link) AddHeader(node *Node) {
	if l.root == nil {
		l.first = node
		l.last = node
		l.root = node
		l.size++
		return
	}
	old := l.root
	l.first = node
	l.root = node
	node.next = old
	l.size++
}
func (l *Link) AddTail(node *Node) {
	if l.last == nil {
		l.first = node
		l.last = node
		l.root = node
		l.size++
		return
	}
	old := l.last
	l.last = node
	old.next = node
	l.size++
}

func (l *Link) GetSize() int {
	return l.size
}
func (l *Link) RemoveHeader() *Node {
	var header *Node
	if l.root != nil {
		header = l.first
		l.root = l.root.next
		l.size--
	}

	return header
}

func (l *Link) RemoveTail() *Node {
	var tail *Node
	if l.last != nil {
		tail = l.last
		for {
			if l.root.next == tail {
				l.root.next = nil
				l.last = l.root
				l.size--
				break
			} else {
				l.root = l.root.next
			}

		}
	}

	return tail
}

func (l Link) Loop() {
	for {
		//fmt.Println(l.root.val)
		if l.root != nil {
			fmt.Printf("当前的链表val==>%v\n", l.root.val)
			l.root = l.root.next
		} else {
			break
		}

	}
}

func main() {
	node1 := NewNode("to", nil)
	node2 := NewNode("be", nil)
	node3 := NewNode("not", nil)
	power := NewPower()
	power.AddTail(node1)
	power.AddTail(node2)
	power.AddTail(node3)
	fmt.Printf("第三个元素是%v\n", power.Get(2).val)
	s := make([]string, 0)
	//s = append(s, 5, 2, 3, 1, 6, 22, 11, 23, 666, 21)
	s = append(s, "a", "sss", "bab", "ss")
	//s = sort.StringSlice{}
	sort.Strings(s)
	fmt.Println(s)
	//var so1 []int
	//so1 = append(so1, 23, 4, 5, 6)
	//so := IntSlice(so1)
	//sort.Sort(so)

	sort1 := SortNum{31, 21, 2, 1, 7}
	HeapSort(sort1)
	//sort1.SelectSort()
	//sort1.InsertSort()
	//sort1.ShellSort()
	//res := MergeSort(sort1)
	//fmt.Println("归并排序", res)
	//res1 := QuickSort(sort1)
	//fmt.Println("快速排序", res1)
	//link := NewLink()
	//link.AddHeader(node1)
	//link.AddTail(node2)
	//link.AddTail(node3)
	//fmt.Printf("link header %v, link tail %v\n", link.root.val, link.last.val)
	//link.Loop()
	//fmt.Printf("current link has %v item,root %v\n", link.GetSize(), link.root.val)
	//link.RemoveHeader()
	//fmt.Printf("current link has %v item\n", link.GetSize())
	//link.Loop()
	//n := link.RemoveTail()
	//fmt.Printf("current link has %v item deletion value is %v\n", link.GetSize(), n.val)

}

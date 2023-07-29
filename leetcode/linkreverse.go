package main

import "fmt"

type Node1 struct {
	val  int
	next *Node1
}

// 链表:1->2->3->4->5->6->7->8->null, K = 3。那么 6->7->8，3->4->5，1->2各位一组。调整后：1->2->5->4->3->8->7->6->null。其中 1，2不调整，因为不够一组。
func main() {
	n1 := &Node1{val: 1, next: nil}
	n2 := &Node1{val: 2, next: nil}
	n3 := &Node1{val: 3, next: nil}
	n4 := &Node1{val: 4, next: nil}
	n5 := &Node1{val: 5, next: nil}
	n6 := &Node1{val: 6, next: nil}
	n7 := &Node1{val: 7, next: nil}
	n8 := &Node1{val: 8, next: nil}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	n5.next = n6
	n6.next = n7
	n7.next = n8
	rnode := reverse(n1)
	res := split(rnode, 3)
	res = reverse(res)
	for res != nil {
		fmt.Print(res.val)
		res = res.next
		if res != nil {
			fmt.Print("->")
		}
	}

}
func split(head *Node1, k int) *Node1 {

	cur := head
	for i := 1; i < k && cur.next != nil; i++ {
		cur = cur.next
	}

	if cur.next == nil {
		return head
	}
	tmp := cur.next
	cur.next = nil
	newhead := reverse(head) //这里已经反转，所以head从头变到尾部
	newtmp := split(tmp, k)
	head.next = newtmp //这里get到的是反转之后的尾部head
	return newhead
}

func reverse(head *Node1) *Node1 {
	if head == nil || head.next == nil {
		return head
	}
	beg := head
	end := beg.next
	for end != nil {
		beg.next = end.next //remove middle node and set head's next to end's next
		end.next = head     //put remove node to head position
		head = end          // setting head node is removed node
		end = beg.next      // setting end position to next node
		// loop until end node is nil
	}
	return head

}

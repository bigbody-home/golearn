package main

import "fmt"

func main() {
	var a []int
	a = append(a, 9, 8, 9, 0)
	fmt.Println(plusOne2(a))

	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{4, nil}
	n9 := &Node{9, nil}
	n1.next = n2
	n2.next = n3
	n3.next = n9

	n4 := &Node{1, nil}
	n5 := &Node{3, nil}
	n6 := &Node{4, nil}
	n4.next = n5
	n5.next = n6

	res := MergeTwoLink(n1, n4)
	for res.next != nil {
		fmt.Print(res.next.val)
		res = res.next
	}
	fmt.Println()
	bn1 := &BinNode{Val: 1}
	bn2 := &BinNode{Val: 2}
	bn3 := &BinNode{Val: 3}
	bn1.Right = bn2
	bn2.Left = bn3
	resarr := make([]int, 0)
	fmt.Println(inorderTraversal(bn1, resarr))

}

func Sum(arr []int, target int) []int {
	mp := make(map[int]int)
	var res []int
	for i, v := range arr {
		if _, ok := mp[target-v]; !ok {
			mp[v] = i
			//res = append(res, i)
		} else {
			res = append(res, i, mp[target-v])

		}

	}
	return res
}

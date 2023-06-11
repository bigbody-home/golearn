package main

type Node struct {
	val  int
	next *Node
}

func (n *Node) Add(s *Node) {
	n.next = s
}

func MergeTwoLink(first *Node, second *Node) *Node {
	res := &Node{}
	prev := res
	for ; first != nil && second != nil; prev = prev.next {
		if first.val >= second.val {
			prev.next = second
			second = second.next
		} else {
			prev.next = first
			first = first.next
		}

	}
	if first != nil {
		prev.next = first
	}
	if second != nil {
		prev.next = second
	}
	return res
}

// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

//你可以假设除了整数 0 之外，这个整数不会以零开头。

func plusOne(digits []int) []int {
	res := make([]int, len(digits))
	size := len(digits)
	lastNum := digits[size-1]
	b := size - 1
	if lastNum+1 < 10 {
		copy(res, digits)
		res[size-1] = lastNum + 1
	} else {
		copy(res, digits)
		for b >= 0 && res[b]+1 == 10 {
			res[b] = 0
			b--
		}
		if b >= 0 {
			res[b] = res[b] + 1
		} else {
			res[0] = 1
			res = append(res, 0)
		}

	}
	return res

}

func plusOne2(digits []int) []int {
	b := len(digits) - 1
	size := len(digits)
	lastNum := digits[size-1]
	if lastNum+1 < 10 {
		digits[size-1] = lastNum + 1
	} else {
		for b >= 0 && digits[b]+1 == 10 {
			digits[b] = 0
			b--
		}
		if b >= 0 {
			digits[b] = digits[b] + 1
		} else {
			digits[0] = 1
			digits = append(digits, 0)
		}

	}

	return digits
}

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历
type BinNode struct {
	Val   int
	Left  *BinNode
	Right *BinNode
}

func inorderTraversal(root *BinNode, res []int) []int {
	//fmt.Println(root.Val, "===")
	if root == nil {
		return nil
	}

	if root.Left != nil {
		//fmt.Println(1)
		res = inorderTraversal(root.Left, res)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		//fmt.Println(2)
		res = inorderTraversal(root.Right, res)
	}
	return res

}

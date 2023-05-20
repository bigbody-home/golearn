package main

import (
	"fmt"
)

type SortNum []int

func (n SortNum) SelectSort() {
	fmt.Println("以第一个数为标准，默认最小，在后面的所有数找出最小的值，然后跟第一个数比较，如果比第一个小，交换位置，最小指针右移动，直到排序完成")
	for i := 0; i < len(n); i++ {
		min := i //最小数的指针，最开始是0
		//比较后面的所有数，找到最小的那个值
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[min] {
				n[min], n[j] = n[j], n[min]
			}

		}
	}
	fmt.Println(n)

}
func (n SortNum) InsertSort() {
	fmt.Println("以第一个值为默认排好序的值，从排好序的那个值开始排序后一位数，找到前面排好序的合适位置插入")
	for i := 1; i < len(n); i++ {
		//for j := i; j > 0; j-- {
		//	if n[j] < n[j-1] {
		//		n[j], n[j-1] = n[j-1], n[j]
		//	}
		//}
		j := i
		for j > 0 {
			//fmt.Println(j)
			if n[j] < n[j-1] {
				n[j], n[j-1] = n[j-1], n[j]
			}
			j--
		}

	}
	fmt.Println(n)
}
func QuickSort(n SortNum) SortNum {
	if len(n) <= 1 {
		return n
	}
	//p := getPartitionNum(n, 0, len(n)-1)
	leftIndex := 0
	rightIndex := len(n) - 1
	value := n[leftIndex]
	for leftIndex < rightIndex {
		if n[leftIndex+1] > value {
			n[leftIndex+1], n[rightIndex] = n[rightIndex], n[leftIndex+1]
			rightIndex--
		} else if n[leftIndex+1] < value {
			n[leftIndex+1], n[leftIndex] = n[leftIndex], n[leftIndex+1]
			leftIndex++
		} else {
			leftIndex++
		}
	}
	p := leftIndex
	//fmt.Println("p is", p)
	QuickSort(n[:p])
	QuickSort(n[p+1:])
	return n
}
func HeapSort(n SortNum) {
	length := len(n)
	for i := length / 2; i >= 1; i-- {
		sink(n, i, length)
	}
	for length > 1 {
		n[1], n[length-1] = n[length-1], n[1]
		sink(n, 1, length)
		length--
	}
	fmt.Println("堆排序:", n)
}
func sink(s SortNum, n int, length int) {

	for n*2 < length {
		fmt.Println("length is ", length, "n is", n)
		j := 2 * n
		if j+1 < length && s[j] < s[j+1] {
			j++
		}
		s[j], s[n] = s[n], s[j]
		n = j
	}
}
func getPartitionNum(n SortNum, lo, hi int) int {
	//num := 0
	//fmt.Println("------>", n, n[lo])
	v := n[lo]
	//fmt.Println("v is", v)[12 21 5 32 42]
	leftIndex := lo + 1
	rightIndex := hi
	for {
		if n[leftIndex] < v {
			leftIndex++
			if leftIndex > hi {
				break
			}
		}
		if n[rightIndex] > v {
			rightIndex--
			if rightIndex < lo {
				break
			}
		}
		if leftIndex >= rightIndex {
			break
		}
		n[leftIndex], n[rightIndex] = n[rightIndex], n[leftIndex]
	}
	fmt.Println("p is", rightIndex)
	n[lo], n[rightIndex] = n[rightIndex], n[lo]
	return rightIndex
}
func MergeSort(n SortNum) SortNum {
	//length := len(n)

	if len(n) <= 1 {
		return n
	}
	left := MergeSort(n[:len(n)/2])
	right := MergeSort(n[len(n)/2:])
	return Merge(left, right)
}
func Merge(left, right SortNum) SortNum {
	var tmp SortNum
	i, j := 0, 0
	total := len(left) + len(right)
	for k := 0; k < total; k++ {
		//处理left元素完成的情况
		if i > len(left)-1 {
			tmp = append(tmp, right[j])
			j++
			continue
		} //处理right的元素完成的情况
		if j > len(right)-1 {
			tmp = append(tmp, left[i])
			i++
			continue
		} //如果左边的元素大于右边，则取右边
		if left[i] > right[j] {
			tmp = append(tmp, right[j])
			j++
			continue
		} //如果左边的元素小于右边，则取左边
		if left[i] <= right[j] {
			tmp = append(tmp, left[i])
			i++
			continue
		}
	}

	return tmp

}
func (n SortNum) ShellSort() {
	h := len(n)
	for h >= 1 {
		for i := h; i < len(n); i++ {
			for j := i; j > 0; j = -h {
				if n[j] < n[j-1] {
					n[j], n[j-1] = n[j-1], n[j]
				}
			}

		}
		h = h / 3
	}
	fmt.Println(n)

}

type Power struct {
	read  map[int]*Node
	head  *Node
	tail  *Node
	size  int
	index int
}

func NewPower() *Power {
	return &Power{}
}
func (l *Power) Get(index int) *Node {
	return l.read[index+1]
}
func (l *Power) AddTail(node *Node) {
	if l.read == nil {
		l.read = make(map[int]*Node)
	}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.size++
		index := l.size
		l.read[index] = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
	l.size++
	index := l.size
	l.read[index] = node
}

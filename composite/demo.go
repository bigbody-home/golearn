package main

import (
	"fmt"
	"golearn/gotest"
)

type Level int

const (
	High Level = iota
	Middle
	Low
)

func main() {
	var a = [3]int{1, 3, 4}
	fmt.Println(a[0])
	for _, v := range a {
		go func(val int) {
			fmt.Println(val)
		}(v)
	}
	//time.Sleep(1 * time.Second)
	l := [...]string{High: "zhangsan", Middle: "lisi", Low: "wangwu"}
	for i, v := range l {
		fmt.Printf("当前的级别是%v,对应的人是%v\n", i, v)
		fmt.Println("aaa")
	}
	s := []string{"hello", "", "world", "", "李四"}
	fmt.Println(nonempty2(s), len(nonempty2(s)))

	var s1 []string
	stack := gotest.NewStack(s1)
	stack.Push("zhangsan")
	stack.Push("lisi")
	stack.Print()
	val := stack.Pop()
	fmt.Println(val)
	stack.Print()
	rmstr := []string{"1", "2", "3"}
	r := remove(rmstr, 1)
	fmt.Println(r)
	iarr := []int{1, 2, 3, 4, 5}
	rarr := reverse(iarr)
	fmt.Println(rarr)

}
func nonempty(str []string) []string {
	i := 0
	for _, v := range str {
		if v != "" {
			str[i] = v
			i++
		}
	}
	return str[:i]

}
func nonempty2(str []string) []string {
	var res []string
	for _, v := range str {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func remove(str []string, i int) []string {
	copy(str[i:], str[i+1:])
	return str[:len(str)-1]

}
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]

	}
	return s

}

func reverse2(s *[7]int) []int {
	s1 := s[:]
	r := reverse(s1)
	return r
}

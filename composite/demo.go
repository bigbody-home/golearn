package main

import (
	"fmt"
	"golearn/gotest"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Level int

const (
	High Level = iota
	Middle
	Low
)

func main() {
	//MapTest()
	//testMiddleLoop()
	//RpcTest()
	fmt.Println("=========test append method=========")
	s1 := []string{"aa", "vv", "cc"}
	s2 := []string{"bb", "bv"}
	s1 = append(s1, s2...)
	fmt.Println(s1)
	fmt.Println("==========test copy method===========")
	copy(s1[:1], s2[:1])
	fmt.Println(s1)

	f := sum
	fmt.Println(f(3, 6))
	sint := []int{2, 3, 4, 5, 3, 2, 4, 6, 7, 7}
	fmt.Println(myslice(sint...))
}
func Mytest() {
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

	stack := gotest.NewStack()
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
func RpcTest() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello2", "luke", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
func MapTest() {
	mp := make(map[string]int)
	str := []string{"hello", "hello", "world", "world", "world", "nihao"}
	for _, v := range str {
		if v1, ok := mp[v]; ok {
			v1++
			mp[v] = v1
		} else {
			mp[v] = 1
		}
	}
	num := count("nihao", mp)
	fmt.Println(num)

}
func count(str string, mp map[string]int) int {
	num := mp[str]

	return num
}

func testMiddleLoop() {
	gotest.PrintTree()

}
func sum(x, y int) int {
	return x + y
}
func myslice(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

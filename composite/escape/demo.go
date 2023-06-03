package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"unsafe"
)

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆

	s.Name = name
	s.Age = age

	return s
}

func main() {
	//StudentRegister("Jim", 18)
	res, err := http.Get("https://www.baidu.com")
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}
	r, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(r))
	fmt.Println(unsafe.Sizeof("l"))
}

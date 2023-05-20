package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type SortPath struct {
	index int
	path  string
}

func main() {
	r, _ := ioutil.ReadDir("/Users/lukezhang/GolandProjects/golearn/composite/")
	var p []SortPath
	for _, f := range r {

		ids := fmt.Sprintf("%d", f.Name()[0])
		id, _ := strconv.Atoi(ids)
		sp := SortPath{
			index: id,
			path:  f.Name(),
		}
		p = append(p, sp)

	}
	SortPrint(p)
}
func SortPrint(s []SortPath) {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s)-1; j++ {
			if s[j].index > s[j+1].index {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	fmt.Println(s)
}

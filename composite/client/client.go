package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	mp := make(map[string]int, 10)
	mp["zhangsan"] = 9
	fmt.Println(mp["lisi"])
	delete(mp, "ki")
	fmt.Println(mp)
}

func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}

}

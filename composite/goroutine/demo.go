package main

import (
	"fmt"
	"time"
)

func main() {
	files := []string{"ss.txt", "sdf.py", "计划报表.txt"}
	DoFiles(files)
	//natural := make(chan int)
	//sq := make(chan int)
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		natural <- i
	//		//time.Sleep(time.Second * 1)
	//	}
	//	close(natural)
	//}()
	//
	//go func() {
	//	for {
	//		num, ok := <-natural
	//		if ok {
	//			sq <- num * num
	//		} else {
	//			break
	//		}
	//
	//	}
	//
	//}()
	//for {
	//	select {
	//	case num := <-sq:
	//		fmt.Println(num)
	//	}
	//
	//}

}

func DoFiles(files []string) {
	ch := make(chan struct{})
	for _, v := range files {
		go func(f string, ch chan struct{}) {
			fmt.Println("File ", f, " Finished", time.Now().Nanosecond())
			ch <- struct{}{}
		}(v, ch)
	}
	for range files {
		<-ch
	}
}

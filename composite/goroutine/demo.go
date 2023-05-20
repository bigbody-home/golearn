package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"sync"
	"time"
)

type S struct {
	Name string
	Size int64
}

func NewS() *S {
	return &S{}
}
func WorkDir(dir fs.FileInfo, c chan S, basepath string, wg *sync.WaitGroup) chan S {
	defer wg.Done()
	if dir.IsDir() {

		//WorkDir(dir, c)
		basepath = basepath + "/" + dir.Name()
		res, err := ioutil.ReadDir(basepath)
		if err != nil {
			panic(err)
		}
		for _, r := range res {
			wg.Add(1)
			go WorkDir(r, c, basepath, wg)
		}
	} else {
		tmp := NewS()
		tmp.Name = dir.Name()
		tmp.Size = dir.Size()
		c <- *tmp
	}

	return c

}
func main() {
	basePath := "/Users/lukezhang/GolandProjects/"
	ch := make(chan S)
	res, err := ioutil.ReadDir(basePath)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, r := range res {
		wg.Add(1)
		go WorkDir(r, ch, basePath, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}
	//files := []string{"ss.txt", "sdf.py", "计划报表.txt"}
	//DoFiles(files)
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

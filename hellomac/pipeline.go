package main

import "fmt"

func main() {
	done := make(chan interface{})
	defer close(done)
	var s []int = []int{1, 2, 3, 4}
	p1 := func(done chan interface{}, inStream []int) <-chan int {
		res := make(chan int)
		go func() {
			defer close(res)
			for _, v := range inStream {
				select {
				case <-done:
					return
				case res <- v:
				}
			}
		}()
		return res
	}

	p2 := func(done chan interface{}, input <-chan int) <-chan int {
		res := make(chan int)
		go func() {
			defer close(res)
			for i := range input {
				select {
				case <-done:
					return
				case res <- i * 2:
				}

			}
		}()
		return res

	}
	for v := range p2(done, p2(done, p1(done, s))) {
		fmt.Println(v)
	}

}

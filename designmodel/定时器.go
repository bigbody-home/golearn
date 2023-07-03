package main

import (
	"fmt"
	"time"
)

func proc() {
	panic("dead")
}
func main() {
	// go func() {
	// 	t := time.NewTimer(1 * time.Second)

	// 	for {
	// 		select {
	// 		case <-t.C:
	// 			go func() {
	// 				defer func() {
	// 					if err := recover(); err != nil {
	// 						fmt.Println(err)
	// 					}
	// 				}()
	// 				proc()
	// 			}()
	// 			t.Reset(1 * time.Second)

	// 		}
	// 	}
	// }()
	CreateEcs()
	for {
	}
	//time.Sleep(7 * time.Second)
}
func CreateEcs() {
	t := time.NewTimer(3 * time.Second)
	res := make(chan interface{})
	go func() {
		time.Sleep(4 * time.Second)
		res <- struct{}{}
	}()

	// if _, ok := <-res; ok {
	// 	fmt.Println("done")
	// 	return
	// }
	select {

	case <-res:
		fmt.Println("create ecs done!")
		return
	case <-t.C:
		fmt.Println("create ecs timeout..")
		t.Reset(3 * time.Second)
		return
	}

}

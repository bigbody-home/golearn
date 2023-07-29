package main

import (
	"fmt"
	"time"
)

func main() {
	catchan := make(chan interface{})
	dogchan := make(chan interface{})
	fishchan := make(chan interface{})

	pcat := func(fish chan interface{}, dogch chan interface{}) {

		select {
		case <-fish:
			fmt.Println("cat")
			dogch <- struct{}{}

		}

	}

	pdog := func(cat chan interface{}, fish chan interface{}) {

		select {
		case <-cat:
			fmt.Println("dog")
			fish <- struct{}{}
		}

	}
	pfish := func(dog chan interface{}, cat chan interface{}) {

		select {

		case <-dog:
			fmt.Println("fish")

			cat <- struct{}{}

		}

	}
	for i := 0; i < 10; i++ {
		go pcat(fishchan, dogchan)
		go pdog(catchan, fishchan)
		go pfish(dogchan, catchan)
	}

	fishchan <- struct{}{}

	arr := []int{1, 2, 3}
	for k, v := range arr {

		fmt.Println(k, "====", v)

	}
	time.Sleep(1 * time.Second)
	
}

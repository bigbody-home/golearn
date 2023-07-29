package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//ProducConsumer()
	test()
}

func test() {
	//map有一定几率报并发写或者读写问题，不一定100%会出现，是非并发安全的，sync.map针对了读多写少场景做了优化
	// mp := &sync.Map{}
	// go func() {
	// 	mp.Store("a", "b")

	// }()
	// go func() {
	// 	mp.Store("b", 1)
	// }()
	mp := make(map[string]interface{})

	go func(mp map[string]interface{}) {
		//time.Sleep(1 * time.Second)
		mp["a"] = 2

	}(mp)
	go func(mp map[string]interface{}) {
		mp["b"] = 2
		fmt.Println(mp["a"])

	}(mp)
	time.Sleep(1 * time.Second)
	fmt.Println(mp["a"])
}
func ProducConsumer() {
	cond := sync.NewCond(new(sync.Mutex))
	condition := 0

	// 消费者
	go func() {
		for {
			// 消费者开始消费时，锁住
			cond.L.Lock()
			// 如果没有可消费的值，则等待
			for condition == 0 {
				cond.Wait()
			}
			// 消费
			condition--
			fmt.Printf("Consumer: %d\n", condition)

			// 唤醒一个生产者
			cond.Signal()
			// 解锁
			cond.L.Unlock()
		}
	}()

	// 生产者
	for {
		// 生产者开始生产
		cond.L.Lock()

		// 当生产太多时，等待消费者消费
		for condition == 100 {
			cond.Wait()
		}
		// 生产
		condition++
		fmt.Printf("Producer: %d\n", condition)
		time.Sleep(2 * time.Second)
		// 通知消费者可以开始消费了
		cond.Signal()
		// 解锁
		cond.L.Unlock()
	}
}

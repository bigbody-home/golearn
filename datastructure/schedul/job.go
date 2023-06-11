package main

import "fmt"

type Work func() bool
type WorkQuene []Job
type Job struct {
	priority int
	work     Work
}

func (worker WorkQuene) Submit(job Job) WorkQuene {
	worker = append(worker, job)
	return worker
}

func NewWorkerQuene() WorkQuene {
	return make(WorkQuene, 0)
}

func (worker WorkQuene) Run() {
	for len(worker) > 0 {
		for _, v := range worker {
			//fmt.Println(i, "---", len(worker))
			//v := worker[i]
			max := worker.getMax()
			if v.priority == max {
				v.work()
				worker = worker[:len(worker)-1]
			}
		}
	}

}
func (w WorkQuene) getMax() int {
	max := 0
	for i, v := range w {
		v = w[i]
		if v.priority > max {
			max = v.priority
		}
	}
	return max
}

func main() {
	w := NewWorkerQuene()
	f1 := func() bool {
		fmt.Println("do f1 job")
		return true
	}
	f2 := func() bool {
		fmt.Println("do f2 job")
		return true
	}
	f3 := func() bool {
		fmt.Println("do f3 job")
		return true
	}
	j1 := Job{priority: 1, work: f1}
	j2 := Job{priority: 2, work: f2}
	j3 := Job{priority: 3, work: f3}
	w = w.Submit(j1)
	w = w.Submit(j2)
	w = w.Submit(j3)
	w.Run()

}

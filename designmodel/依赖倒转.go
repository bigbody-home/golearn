package main

import (
	"fmt"
	"sync"
	"time"
)

type Jupyter interface {
	Create()
	Update()
	Get()
	Delete()
}
type WriterLog interface {
	Write()
}
type Alicloud struct {
	createBy       string
	specifications string
	expiretime     string
	updateBy       string
	deleteBy       string
	id             int
	teamId         int
}

func (a *Alicloud) Create() {
	fmt.Printf("用户%s,%v在阿里云创建了一台jupyter\n", a.createBy, time.Now())
}
func (a *Alicloud) Update() {
	fmt.Println("更新jupyter状态")
}
func (a *Alicloud) Get() {
	fmt.Println("get jupyter")
}
func (a *Alicloud) Delete() {
	fmt.Println("delete jupyter")
}
func (a *Alicloud) Write() {
	fmt.Println("记录日志到文件中和数据库")
}
func CreateJupyter(p Jupyter, w WriterLog) {
	p.Create()
	w.Write()
}

func main() {
	a := Alicloud{createBy: "45142572", specifications: "small", expiretime: "5 day", updateBy: "", teamId: 117}
	CreateJupyter(&a, &a)
	p := sync.Pool{}
	p.New = func() any {
		return &Alicloud{}
	}
	for i := 0; i < 10; i++ {

		go func() {
			t := p.Get().(*Alicloud)			
			t.Delete()
			p.Put(t)
		}()
	}
	time.Sleep(1 * time.Second)
}

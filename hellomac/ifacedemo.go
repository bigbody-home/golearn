package main

import "fmt"

type Fight interface {
	FightOne()
}
type Fix interface {
	FixCar()
}
type MixFF interface {
	Fight
	Fix
	Read()
}
type zhangsan struct{}

func (z *zhangsan) FightOne() {
	//TODO implement me
	fmt.Println("I am zhangsan i can fight one")
}

func (z *zhangsan) FixCar() {
	fmt.Println("i am zhangsan i can fix car")
}

func (z *zhangsan) Read() {
	//TODO implement me
	fmt.Println("i am zhangsan i also can read")
}
func DoFixCarAndRead(f MixFF) {
	f.FixCar()
	f.Read()
}
func DoFightAndRead(f MixFF) {
	f.FightOne()
	f.Read()
}
func main() {
	zs := &zhangsan{}
	DoFixCarAndRead(zs)
	DoFightAndRead(zs)

}

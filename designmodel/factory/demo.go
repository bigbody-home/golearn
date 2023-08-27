package main

import "fmt"

type Transport interface {
	Run(from string, to string) error
	Fee() float32
}
type TranspotTool struct {
	name  string
	way   string
	fee   float32
	goods string
}

func (t *TranspotTool) Run(from string, to string) error {
	fmt.Printf("运输工具%v从%v运输%v到%v", t.name, from, t.goods, to)
	return nil
}
func (t *TranspotTool) Fee() float32 {

	return t.fee
}

type Car struct {
	TranspotTool
}

func NewCar(goods string) *Car {
	return &Car{
		TranspotTool: TranspotTool{
			name:  "big car",
			way:   "mainland",
			fee:   1087.98,
			goods: goods,
		},
	}
}

type Ship struct {
	TranspotTool
}

func NewShip(goods string) *Ship {
	return &Ship{
		TranspotTool: TranspotTool{
			name:  "Big Boat",
			way:   "Sea",
			fee:   2034.8,
			goods: goods,
		},
	}
}
func FactoryGet(ptype string, goods string) Transport {
	if ptype == "land" {
		return NewCar(goods)
	}
	if ptype == "water" {
		return NewShip(goods)
	}
	return nil
}

func main() {
	methord := "land"
	goods := "门窗32吨"
	product := FactoryGet(methord, goods)
	product.Run("西安", "北京")
	fmt.Printf("花费%.2f\n", product.Fee())

	methord2 := "water"
	goods2 := "三十吨鱼"
	product2 := FactoryGet(methord2, goods2)
	product2.Run("日本", "上海")
	fmt.Printf("花费%.2f", product2.Fee())
}

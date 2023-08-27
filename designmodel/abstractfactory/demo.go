package main

import "fmt"

type ICheiken interface {
	SaleChiken()
}
type IJuice interface {
	SaleJuice()
}
type AbstractFactory interface {
	MakeChikenFactory() ICheiken
	MakeJuiceFactory() IJuice
}
type Chiken struct {
	ctype string
	name  string
}

func (c *Chiken) SaleChiken() {
	fmt.Println("麦当劳售卖炸鸡")
}

type McDonald struct{}

func (m *McDonald) MakeChikenFactory() ICheiken {
	return &Chiken{
		ctype: "鸡",
		name:  "大油炸鸡",
	}
}

func main() {

}

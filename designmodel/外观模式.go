package main

type Food interface {
	Make()
}
type Coffee struct{}

func (c *Coffee) Make() {
	println("制作咖啡")
}

type Milk struct{}

func (c *Milk) Make() {
	println("制作牛奶")
}

type Facade struct {
	food Food
}

func NewFacade() *Facade {
	return new(Facade)
}
func (f *Facade) GetFood(food Food) {
	f.food = food
	f.food.Make()
}

func main() {
	milk := new(Milk)
	f := NewFacade()
	f.GetFood(milk)
	cafe := new(Coffee)
	f.GetFood(cafe)

}

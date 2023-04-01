package main

import "fmt"

type Circle struct {
	X      int
	Y      int
	Radium int
}
type Wheel struct {
	Circle
	Color int
}

func main() {
	fmt.Println("hello")
	fmt.Println("hello fmt")
	var wheel Wheel
	wheel.X = 3
	wheel.Y = 4
	wheel.Radium = 6
	wheel.Color = 8
	fmt.Println(wheel)
	wheel2 := &Wheel{
		Circle: Circle{3, 4, 6},
		Color:  10,
	}
	fmt.Println(wheel2)
}

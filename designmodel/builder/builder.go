package main

import "fmt"

type House struct {
	wall   string
	flool  int
	window string
}

type Builder interface {
	BuildWall()
	BuildFlool()
	InstallWindow()
	GetHouse() House
}

func GetHouse(htype string) Builder {
	if htype == "wood" {
		return NewWoodHouse()
	}
	return nil
}

func main() {
	wood := GetHouse("wood")

	d := NewDirector(&wood)
	h := d.Do()
	fmt.Print(h)
}

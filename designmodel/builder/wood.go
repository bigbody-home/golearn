package main

type WoodHouse struct {
	wall   string
	flool  int
	window string
}

func NewWoodHouse() *WoodHouse {
	return &WoodHouse{}
}
func (w *WoodHouse) BuildWall() {
	w.wall = "wood wall"
}
func (w *WoodHouse) BuildFlool() {
	w.flool = 1
}
func (w *WoodHouse) InstallWindow() {
	w.window = "尼罗姆"
}
func (w *WoodHouse) GetHouse() House {
	return House{
		wall:   w.wall,
		flool:  w.flool,
		window: w.window,
	}
}

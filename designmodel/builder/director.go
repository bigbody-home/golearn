package main

type Director struct {
	builder Builder
}

func (d *Director) SetBUilder(b Builder) {
	d.builder = b
}

func NewDirector(b *Builder) *Director {
	return &Director{
		builder: *b,
	}
}

func (d *Director) Do() House {
	d.builder.BuildWall()
	d.builder.BuildFlool()
	d.builder.InstallWindow()
	return d.builder.GetHouse()
}

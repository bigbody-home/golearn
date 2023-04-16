package main

type Power struct {
	read  map[int]*Node
	head  *Node
	tail  *Node
	size  int
	index int
}

func NewPower() *Power {
	return &Power{}
}
func (l *Power) Get(index int) *Node {
	return l.read[index+1]
}
func (l *Power) AddTail(node *Node) {
	if l.read == nil {
		l.read = make(map[int]*Node)
	}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.size++
		index := l.size
		l.read[index] = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
	l.size++
	index := l.size
	l.read[index] = node
}

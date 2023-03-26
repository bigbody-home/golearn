package gotest

import "fmt"

type Stack struct {
	container []string
}

func NewStack(container []string) *Stack {
	return &Stack{container: container}
}

func (s *Stack) Pop() string {
	//TODO implement me
	res := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return res
}

func (s *Stack) Push(val string) {
	//TODO implement me
	s.container = append(s.container, val)

}

func (s *Stack) Count() int {
	//TODO implement me
	return len(s.container)
}

func (s *Stack) GetVal(i int) string {
	//TODO implement me
	if i < 0 || i > len(s.container) {
		return ""
	}
	return s.container[i]
}
func (s *Stack) Print() {
	//TODO implement me
	fmt.Println(s.container)
}

type InterfaceStack interface {
	Pop() string
	Push(val string)
	Count() int
	GetVal(i int) string
	Print()
}

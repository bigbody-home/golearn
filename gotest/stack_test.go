package gotest

import (
	"fmt"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	container := []string{"hello", "world"}
	stack := NewStack(container)
	val := stack.Pop()
	if val != "world" {
		t.Errorf("test failed")
	}
}
func TestStack_Push(t *testing.T) {
	container := []string{"hello", "world"}
	stack := NewStack(container)
	stack.Push("ok")
	val := stack.Pop()
	if val != "ok" {
		t.Errorf("test failed")
	}
}
func TestStack_Count(t *testing.T) {
	container := []string{"hello", "world"}
	stack := NewStack(container)
	stack.Push("ok")
	stack.Pop()
	fmt.Println(len(stack.container))
	if len(stack.container) != 2 {
		t.Errorf("test failed")
	}
}
func TestStack_GetVal(t *testing.T) {
	container := []string{"hello", "world"}
	stack := NewStack(container)
	stack.Push("ok")
	val := stack.container[1]
	if val != "world" {
		t.Errorf("test failed")
	}
}

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Num interface {
	int | int16 | int32 | int64 | int8 | float32 | float64
}

func Add[T constraints.Ordered](a T, b T) T {

	return a + b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(3.2, 3.3))
}

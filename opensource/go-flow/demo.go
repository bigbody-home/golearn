package main

import (
	"fmt"

	goflow "github.com/kamildrazkiewicz/go-flow"
)

func main() {
	f1 := func(res map[string]interface{}) (interface{}, error) {
		fmt.Println(1)
		return 1, nil
	}
	f2 := func(res map[string]interface{}) (interface{}, error) {
		fmt.Println(2)
		return 2, nil
	}
	f3 := func(res map[string]interface{}) (interface{}, error) {
		fmt.Println(3)
		return 3, nil
	}
	f := goflow.New().Add("first", nil, f1).Add("second", []string{"first"}, f2).Add("third", []string{"second"}, f3)

	f.Do()
}

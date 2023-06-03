package main

import (
	"fmt"
	"reflect"
)

type Card struct {
	Id    int
	Owner string
}

func main() {
	var a int64 = 33
	res := reflect.TypeOf(a)
	fmt.Println(res.Name(), res.String(), res.Kind())
	c := Card{}
	c.Owner = "luke"
	c.Id = 1234
	res = reflect.TypeOf(c)
	fmt.Println(res.String())
	r := reflect.ValueOf(c)
	fmt.Println(r.FieldByName("Id"))
}

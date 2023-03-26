package gotest

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	s := "123456789"
	res := Comma(s)
	fmt.Println(res)
	if res != "123,456,789" {
		t.Errorf("结果预期不符")
	}
}
func TestCompareStr(t *testing.T) {
	s1 := "hello i am luke"
	s2 := "hello i am luke"
	if !CompareStr(s1, s2) {
		t.Errorf("it not good")
	}
	t.Log("it is good job! congratulations!")
}

package main

import (
	"encoding/json"
	"fmt"
	"golearn/dsw/factory"
)

func main() {
	uploader := factory.GetOssUploader("45142572", "oss-bucket-01", "osspath/01/02/sv.text")
	res := factory.CommonCharge(uploader)
	j, _ := json.Marshal(res)
	fmt.Println(string(j))
	a := 12111112
	b := int16(a)
	println(b)
	println(a)
	println("min func res is", min(23, 55))
	fmt.Println(doubleScore(0))
	fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))
}
func min(a, b int) int {
	//min :=0
	min := copy(make([]struct{}, a), make([]struct{}, b))
	return min
}
func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	return source * 2
}

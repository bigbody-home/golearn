package main

import "fmt"

func main() {
	// 2 7 1 22 --9
	test := []int{2, 11, 7, 5}
	target := 9
	fmt.Println(Sum(test, target))
}

func Sum(arr []int, target int) []int {
	mp := make(map[int]int)
	var res []int
	for i, v := range arr {
		if _, ok := mp[target-v]; !ok {
			mp[v] = i
			//res = append(res, i)
		} else {
			res = append(res, i, mp[target-v])

		}

	}
	return res
}

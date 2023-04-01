package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Author []string
}

func main() {

	var movie = []*Movie{
		&Movie{
			Title: "龙门客栈系列",
			Name:  "龙门客栈",
			Year:  20,
			Author: []string{
				"张曼玉",
				"甄子丹",
				"梁朝伟",
			},
		},
		&Movie{
			Title: "开心鬼系列",
			Name:  "开心鬼放暑假",
			Year:  20,
			Author: []string{
				"黄百鸣",
				"袁洁莹",
			},
		},
	}
	//res, _ := json.Marshal(movie)
	res1, _ := json.MarshalIndent(movie, "", "  ")

	fmt.Printf("%s\n", res1)

	//log.Fatalln(err)
	//fmt.Println(res)
}

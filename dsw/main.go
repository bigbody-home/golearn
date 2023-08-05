package main

import (
	"fmt"
	"golearn/dsw/factory"
)

func main() {
	uploader := factory.GetOssUploader("45142572", "oss-bucket-01", "osspath/01/02/sv.text")
	res := uploader.Upload()
	fmt.Println(res)

}

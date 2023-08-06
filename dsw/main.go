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

}

package upload

import (
	"fmt"
	"golearn/dsw/common"
	"testing"
)

func TestOssUploader_Upload(t *testing.T) {
	obj := NewOssUploader("123", "aaa", "/a/b/s")
	res := obj.Upload()
	res1 := res.(*common.ResCommonHandler).Code
	fmt.Println(res1)
	if res1 == 200 {
		t.Log("test successful")
	} else {
		fmt.Println("Failed")
		t.Fatal("Failed")

	}
}

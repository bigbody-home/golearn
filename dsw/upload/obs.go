package upload

import (
	"fmt"
	"golearn/dsw/common"
	"gorm.io/gorm"
)

type Uploader interface {
	Upload() common.Handler
	GetRes() common.Handler
}

type OssUploader struct {
	*gorm.Model
	CreateBy    string
	Bucket      string
	FilePath    string
	ProductType string
}

func NewOssUploader(createBy string, bucket string, filePath string) *OssUploader {
	return &OssUploader{CreateBy: createBy, Bucket: bucket, FilePath: filePath, ProductType: "OSS"}
}

func (o *OssUploader) Upload() common.Handler {

	fmt.Printf("%v upload file %v to %v OSS bucket", o.CreateBy, o.Bucket, o.FilePath)
	fmt.Println("step1 upload to oss")
	fmt.Println("step2 record db if success")
	fmt.Println("commit res")
	res := common.NewDSWHandler(200, "success")
	return res
}

func (o *OssUploader) GetRes() common.Handler {

	fmt.Println("get res done")
	s := &struct {
		Name    string
		Age     int
		Address string
	}{Name: "zhangsan", Age: 12, Address: "Beijing"}
	res := common.NewResWithData(1, 10, 200, s, 5)
	return res
}

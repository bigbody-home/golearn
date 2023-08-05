package upload

import (
	"fmt"
	"golearn/dsw/common"

	"gorm.io/gorm"
)

type Uploader interface {
	Upload() error
}

type OssUploader struct {
	*gorm.Model
	CreateBy    string
	Bucket      string
	FilePath    string
	ProductType string
}

func NewOssUploader(createBy string, bucket string, filePath string, productType string) *OssUploader {
	return &OssUploader{CreateBy: createBy, Bucket: bucket, FilePath: filePath, ProductType: productType}
}

func (o *OssUploader) Upload() error {

	fmt.Printf("%v upload one file to oss, ossfile info: %v,%v\n", o.CreateBy, o.FilePath, o.Bucket)
	println("step1->upload to oss")
	println("step2->record res to database")
	println("step3->return res to user")
	res := common.NewDSWError(200, "Success")

	return res

}

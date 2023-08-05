package factory

import "golearn/dsw/upload"

func GetOssUploader(createBy string, Bucket string, filepath string) *upload.OssUploader {

	return upload.NewOssUploader(createBy, Bucket, filepath, "OSS")
}

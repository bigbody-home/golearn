package factory

import (
	"golearn/dsw/common"
	"golearn/dsw/upload"
)

func CommonCharge(o upload.Uploader) common.Handler {
	res := o.Upload()
	return res
}

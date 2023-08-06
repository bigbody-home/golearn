package upload

import "golearn/dsw/common"

func CommonCharge(o Uploader) common.Handler {
	res := o.Upload()
	return res
}

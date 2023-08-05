package upload

func CommonCharge(o Uploader) error {
	res := o.Upload()
	return res
}

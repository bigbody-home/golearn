package gotest

import "crypto/sha256"

func Comma(str string) string {
	length := len(str)
	bstr := []byte(str)
	if length <= 3 {
		return str
	}
	return Comma(string(bstr[:length-3])) + "," + string(bstr[length-3:])
}

func CompareStr(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sha1 := sha256.Sum256(b1)
	sha2 := sha256.Sum256(b2)

	return sha1 == sha2

}
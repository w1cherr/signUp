package service

import (
	"crypto/md5"
	"encoding/hex"
)

func StrToMD5(s string) string{
	h := md5.New()
	h.Write([]byte(s))
	res := hex.EncodeToString(h.Sum(nil))
	return res
}

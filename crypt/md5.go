package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Check(data string, md5 string) bool {
	return strings.EqualFold(md5, Encode(data))
}

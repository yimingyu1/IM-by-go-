package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func EncodePwd(pwd, salt string) string {
	return Md5Encode(pwd + salt)
}

func ValidPassword(inputPwd, salt, truePwd string) bool {
	return EncodePwd(inputPwd, salt) == truePwd
}

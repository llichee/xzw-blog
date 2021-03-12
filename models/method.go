package models

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x\n", md5.Sum(data))
}

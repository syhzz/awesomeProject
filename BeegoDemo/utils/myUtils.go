package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func SwitchTimeStampToData(created int64) string {
	return time.Unix(created, 0).String()
}

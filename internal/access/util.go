package access

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(password string) string {
	b := md5.Sum([]byte(password))
	return hex.EncodeToString(b[:])
}

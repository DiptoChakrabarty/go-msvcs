package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateMD5(s string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

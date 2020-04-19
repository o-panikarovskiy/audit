package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512 returns hex of sha512
func SHA512(str string, salt string) string {
	h := sha512.New()
	h.Write([]byte(str + salt))
	return hex.EncodeToString(h.Sum(nil))
}

package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

// RandomString gen random string
func RandomString(size int) string {
	buff := make([]byte, size)

	_, err := rand.Read(buff)
	if err != nil {
		log.Fatal(err)
	}

	str := hex.EncodeToString(buff)
	return str[:size]
}

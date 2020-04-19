package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"math"
)

// RandomString gen random string
func RandomString(size int) string {
	buff := make([]byte, int(math.Round(float64(size)/2)))

	_, err := rand.Read(buff)
	if err != nil {
		log.Fatal(err)
	}

	str := hex.EncodeToString(buff)
	return str[:size]
}

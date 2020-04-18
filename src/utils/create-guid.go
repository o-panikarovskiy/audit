package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

// CreateGUID generates guid string
func CreateGUID() string {
	buff := make([]byte, 16)
	_, err := rand.Read(buff)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", buff[0:4], buff[4:6], buff[6:8], buff[8:10], buff[10:])
}

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

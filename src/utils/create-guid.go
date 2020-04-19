package utils

import (
	"crypto/rand"
	"fmt"
	"log"
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

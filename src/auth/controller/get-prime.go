package controller

import (
	"crypto/rand"
	"math/big"
)

// GetPrime returns random prime number
func GetPrime() (*big.Int, error) {
	p, err := rand.Prime(rand.Reader, 100)
	if err != nil {
		return nil, err
	}
	return p, nil
}

package utils

import (
	"crypto/rand"
	"math/big"
)

func GetRandWithIn1000() int64 {
	return GetRand(1000)
}

func GetRand(num int64) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(num))
	return result.Int64()
}

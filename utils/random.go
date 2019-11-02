package utils

import (
	cRand "crypto/rand"
	"math/big"
	mRand "math/rand"
	"time"
)

func MathRandNum(max int64) int64 {
	mRand.Seed(time.Now().Unix())
	return mRand.Int63n(max)
}

func RandNum(max int64) int64 {
	if max <= 0 {
		return 0
	}
	index, err := cRand.Int(cRand.Reader, big.NewInt(max))
	if err != nil {
		return MathRandNum(max)
	}
	return index.Int64()
}

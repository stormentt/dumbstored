package random

import (
  "crypto/rand"
  "math/big"
)

func Salt() string {
  return alphaNum(16)
}

func Int(max uint64) uint64 {
	bigMax := big.NewInt(0)
  bigMax.SetUint64(max)
	randInt, err := rand.Int(rand.Reader, bigMax)
	if err != nil {
		panic(err)
	}
	return randInt.Uint64()
}

func alphaNum(length int) string {
	alphanumeric := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	alphaLen := uint64(len(alphanumeric))

	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		letter := Int(alphaLen)
		ret[i] = alphanumeric[letter]
	}

	return string(ret)
}

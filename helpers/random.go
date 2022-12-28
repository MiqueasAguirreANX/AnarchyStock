package helpers

import (
	rand "crypto/rand"
	"fmt"
	"math/big"
)

func GenerateSecureRandomInteger(length uint64) (v uint64) {
	b := big.NewInt(int64(length))
	n, err := rand.Int(rand.Reader, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	v = n.Uint64() + 1
	return v
}

func GenerateSecureRandomString(length uint64) (v string) {
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	buf := make([]byte, length)
	var i uint64
	for i = 0; i < length; i++ {
		buf[i] = all[GenerateSecureRandomInteger(length)]
	}
	v = string(buf)
	return v
}

func GenerateSecureRandomToken(length uint64) (v string) {
	digits := "0123456789"
	buf := make([]byte, length)
	var i uint64
	for i = 0; i < length; i++ {
		buf[i] = digits[GenerateSecureRandomInteger(length)]
	}
	v = string(buf)
	return v
}

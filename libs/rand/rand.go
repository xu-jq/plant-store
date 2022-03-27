package rand

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allLetters   = lowerLetters + upperLetters

	numbers = "1234567890"
)

func CreateRandomString(len int) string {
	return createRandomString(len, allLetters+numbers)
}

func createRandomString(len int, str string) string {
	var container string
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

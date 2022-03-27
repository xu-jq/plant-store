package tools

import (
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"

	"bee_project/libs/rand"
)

const (
	iter   = 1000
	keyLen = 32
)

func EncryptPasswordWithSalt(password, salt string) string {
	pwdEn := encryptPasswordWithSalt([]byte(password), []byte(salt))
	return base64.StdEncoding.EncodeToString(pwdEn)
}

func encryptPasswordWithSalt(pwd, salt []byte) []byte {
	pwd = append(pwd, salt...)
	pwdEn := pbkdf2.Key(pwd, salt, iter, keyLen, sha256.New)
	return pwdEn
}

func GenerateSalt() string {
	return rand.CreateRandomString(16)
}

func IsEncryptPwdMatch(pwd, salt, encrypt string) bool {
	if len(encrypt) == 0 {
		return false
	}
	pwdEnBase64 := base64.StdEncoding.EncodeToString(encryptPasswordWithSalt([]byte(pwd), []byte(salt)))

	return encrypt == pwdEnBase64
}

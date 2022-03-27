package tools

import (
	"testing"
)

func TestEncrypt(t *testing.T) {

	truePass := "123456"
	salt := GenerateSalt()
	encrypt := EncryptPasswordWithSalt(truePass, salt)

	if !IsEncryptPwdMatch(truePass, salt, encrypt) {
		t.Errorf("the true pass not match the encrypt")
	}

	wrongPass := "1234567"
	if IsEncryptPwdMatch(wrongPass, salt, encrypt) {
		t.Errorf("the wrong pass match the encrypt")
	}
}

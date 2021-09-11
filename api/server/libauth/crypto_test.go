package auth

import (
	"fmt"
	"testing"
)

func TestCrypto(t *testing.T) {
	plainText := "Makabe mizuki is cute"
	key := "passw0rdpassw0rdpassw0rdpassw0rd"

	encryptedText, err := Encrypt([]byte(plainText), []byte(key))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("EncryptedText: %s\n", encryptedText)

	decryptedText, err := Decrypt(encryptedText, []byte(key))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("DecryptedText: %s\n", decryptedText)

	if plainText != string(decryptedText) {
		t.Error("Decrypted text is not equal to original text")
	}
}

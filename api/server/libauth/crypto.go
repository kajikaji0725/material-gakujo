package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func Encrypt(plainText, key []byte) (string, error) {
	encrypted := make([]byte, aes.BlockSize+len(plainText))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCTR(block, iv)
	mode.XORKeyStream(encrypted[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(encrypted), nil
}

func Decrypt(encryptedTextStr string, key []byte) ([]byte, error) {
	encryptedText, err := base64.URLEncoding.DecodeString(encryptedTextStr)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decryptedText := make([]byte, len(encryptedText[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(block, encryptedText[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptedText, encryptedText[aes.BlockSize:])

	return decryptedText, nil
}

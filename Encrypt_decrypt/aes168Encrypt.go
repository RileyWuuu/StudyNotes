package Encrypt_decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var (
	desKey = "ABCDEFG123456789" //must be 16 bytes
)

func EncryptMessage(plaintext []byte, key []byte) (string, error) {
	// block, err := aes.NewCipher(key)
	// if err != nil {
	// 	return "", err
	// }

	// // Pad the plaintext if its length is not a multiple of the block size (16 bytes).
	// padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	// paddedPlaintext := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// ciphertext := make([]byte, len(paddedPlaintext))

	// // Encrypt each block of plaintext with the AES block cipher.
	// for i := 0; i < len(paddedPlaintext); i += aes.BlockSize {
	// 	block.Encrypt(ciphertext[i:i+aes.BlockSize], paddedPlaintext[i:i+aes.BlockSize])
	// }
	// return base64.StdEncoding.EncodeToString(ciphertext), nil
	byteMsg := []byte(plaintext)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptMessage(key []byte, message string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

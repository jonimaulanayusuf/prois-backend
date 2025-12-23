package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"prois-backend/internal/config"
	"strconv"
)

func getAESKey(secret string) []byte {
	hash := sha256.Sum256([]byte(secret)) // hasilnya selalu 32 byte
	return hash[:]
}

func getSecret() []byte {
	config.LoadEnv()
	secret := config.GetEnv("ENC_SECRET", "")
	return getAESKey(secret)
}

func EncryptID(id uint) string {
	secret := getSecret()

	plaintext := []byte(strconv.Itoa(int(id)))

	block, _ := aes.NewCipher(secret)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	_, _ = io.ReadFull(rand.Reader, iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func DecryptID(enc string) *uint {
	secret := getSecret()

	ciphertext, err := base64.URLEncoding.DecodeString(enc)
	if err != nil || len(ciphertext) < aes.BlockSize {
		return nil
	}

	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	id64, err := strconv.ParseUint(string(ciphertext), 10, 64)
	if err != nil {
		return nil
	}

	id := uint(id64)
	return &id
}

package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateInvoiceNumber() (string, error) {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	randomPart := hex.EncodeToString(b)
	datePart := time.Now().Format("20060102")

	return fmt.Sprintf("INV-%s-%s", datePart, randomPart), nil
}

package otp

import (
	"crypto/rand"
	"encoding/base32"
)

func GenerateOTP(length int) string {
	randomBytes := make([]byte, 4)
	_, _ = rand.Read(randomBytes)
	otp := base32.StdEncoding.EncodeToString(randomBytes)
	return otp[:length]
}

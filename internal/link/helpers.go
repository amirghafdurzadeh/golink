package link

import (
	"crypto/rand"
	"math/big"
)

func buildCode(customCode string, length int) (string, error) {
	if customCode != "" {
		return customCode, nil
	}

	return generateShortCode(length)
}

func generateShortCode(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)

	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		b[i] = charset[n.Int64()]
	}

	return string(b), nil
}

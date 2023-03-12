package service

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/btcsuite/btcutil/base58"
	"fmt"
)

func GenerateRandomPublicKey() (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 256)
	if err != nil {
		return "", err
	}

	publicKeyBytes := privateKey.PublicKey.N.Bytes()
	publicKeyBase58 := base58.Encode(publicKeyBytes)

	return publicKeyBase58, nil
}

func CreateSolanaPayLink(recipient, reference string, amount float32) string {
	amountS := fmt.Sprintf("%.2f", amount)
	return fmt.Sprintf("solana:%s?amount=%s&reference=%s", recipient, amountS, reference)
}

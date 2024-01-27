package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func signTransaction(privateKey *ecdsa.PrivateKey, data []byte) (*big.Int, *big.Int, error) {
	hash := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, nil, err
	}
	return r, s, nil
}

func verifyTransaction(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	hash := sha256.Sum256(data)
	return ecdsa.Verify(publicKey, hash[:], r, s)
}
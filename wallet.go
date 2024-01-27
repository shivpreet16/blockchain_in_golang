package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
	"errors"
    // "fmt"
)

func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
    curve := elliptic.P256()
    private, err := ecdsa.GenerateKey(curve, rand.Reader)
    if err != nil {
        return nil, nil, err
    }
    public := &private.PublicKey

    return private, public, nil
}

func generateWallet() (*Wallet, error) {
    privateKey, publicKey, err := generateKeyPair()
    if err != nil {
        return nil, err
    }

    address := generateAddress(publicKey)

    return &Wallet{
        PrivateKey: privateKey,
        PublicKey:  publicKey,
        Address:    address,
    }, nil
}

func generateAddress(publicKey *ecdsa.PublicKey) string {
    publicKeyBytes := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
    hash := sha256.Sum256(publicKeyBytes)
    return hex.EncodeToString(hash[:])
}

func (wallet *Wallet) addEth(eth int) (error, int) {
    if eth <= 0 {
        return errors.New("Please provide a positive value"), 0
    }

    wallet.Eth += eth
    return nil, wallet.Eth
}

func (wallet *Wallet) debitEth(eth int) (error, int) {
    if eth <= 0 {
        return errors.New("Please provide a positive value"), 0
    } else if (wallet.Eth-eth) <= 0 {
        return errors.New("Not enough balance"), 0
    }

    wallet.Eth -= eth
    return nil, wallet.Eth
}

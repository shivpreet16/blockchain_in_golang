package main

import (
	"crypto/ecdsa"
)

type Transaction struct {
	From   *ecdsa.PublicKey
	To     *ecdsa.PublicKey
	Amount int
}

type Block struct{
	Timestamp int64
	PreviousHash []byte
	CurrHash []byte
	AllData []byte 
}

type Blockchain struct{
	Blocks []*Block
}

type Wallet struct {
    PrivateKey *ecdsa.PrivateKey
    PublicKey  *ecdsa.PublicKey
	Eth int
	Email string
    Address string
}

/*
type Transaction struct{
	...
}

type Block struct{
	Timestamp int64
	PreviousHash []byte
	CurrHash []byte
	Data []byte 
	Address string
	digitalSignature string
	transaction []Transaction
	Nonce int
}

Do txn -> data+nonce === hash || check verifyTXN()
*/
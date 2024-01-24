package main

import(
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

func (block *Block) SetHash(){
	timestamp:=[]byte(strconv.FormatInt(block.Timestamp, 10))

	headers:=bytes.Join([][]byte{timestamp, block.PreviousHash, block.AllData},[]byte{})

	hash := sha256.Sum256(headers)

	block.CurrHash=hash[:]
}

func NewBlock(data string, previousHash []byte) *Block{
	block:=&Block{time.Now().Unix(), previousHash,[]byte{}, []byte(data)}

	block.SetHash()

	return block
}

func NewGenesisBlock() *Block{
	return NewBlock("Genesis",[]byte{})
}
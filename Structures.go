package main

type Block struct{
	Timestamp int64
	PreviousHash []byte
	CurrHash []byte
	AllData []byte
}

type Blockchain struct{
	Blocks []*Block
}
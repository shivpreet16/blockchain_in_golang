package main

func (blockchain *Blockchain) AddBlock(data string){
	PreviosBlock:=blockchain.Blocks[len(blockchain.Blocks)-1]

	newBlock := NewBlock(data, PreviosBlock.CurrHash)

	blockchain.Blocks = append(blockchain.Blocks,newBlock)
}

func NewBlockchain() *Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
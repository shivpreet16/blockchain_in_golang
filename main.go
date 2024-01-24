package main

import(
	"fmt"
)

func main(){
	fmt.Println("First Blockchain!!")
	
	newblockchain:=NewBlockchain()

	newblockchain.AddBlock("1st blaock i mined goes to J")

	for _,block := range newblockchain.Blocks{
		fmt.Println("Hash of the block: ", block.CurrHash)

		fmt.Println("Hash of prev block: ", block.PreviousHash)

		fmt.Println("Data: ", string(block.AllData))
	}
}
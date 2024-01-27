package main

import(
	"fmt"
	"crypto/ecdsa"
	"encoding/hex"
)

func PublicKeyToString(publicKey *ecdsa.PublicKey) string {
    xStr := hex.EncodeToString(publicKey.X.Bytes())
    yStr := hex.EncodeToString(publicKey.Y.Bytes())
    return fmt.Sprintf("%s:%s", xStr, yStr)
} 	

func PrivateKeyToString(privateKey *ecdsa.PrivateKey) string {
	dStr := hex.EncodeToString(privateKey.D.Bytes())
	publicKeyStr := PublicKeyToString(&privateKey.PublicKey)
	return fmt.Sprintf("%s:%s", dStr, publicKeyStr)
}


func main() {
    fmt.Println("First Blockchain!!")

    newBlockchain := NewBlockchain()

    wallet, err := generateWallet()
    if err != nil {
        panic(err)
    }

    newBlockchain.AddBlock("1st block mined goes to J")

    fmt.Println("\nWallet Address:", wallet.Address)
    fmt.Println("Wallet Public:", *wallet.PublicKey,"\n")
    fmt.Println("Wallet Private:", *wallet.PrivateKey,"\n\n")

    wallet.addEth(10)
    err,eth:=wallet.addEth(10)
    if err!=nil{
        fmt.Println("Error:", err)
    } else{
        fmt.Println("New Eth Balance:", eth)
    }
    err,eth=wallet.debitEth(10)
    if err!=nil{
        fmt.Println("Error:", err)
    } else{
        fmt.Println("New Eth Balance:", eth)
    }

	// transactionData := []byte("Transaction data")

	// r,s,err:=signTransaction(wallet.PrivateKey,transactionData)

	// if err != nil {
	// 	panic(err)
	// }

	// isValid := verifyTransaction(wallet.PublicKey, transactionData, r, s)

	// if isValid {
	// 	fmt.Println("Transaction signature is valid.")
	// } else {
	// 	fmt.Println("Transaction signature is NOT valid.")
	// }
    // for _, block := range newBlockchain.Blocks {
    //     fmt.Println("Hash of the block:", block.CurrHash)
    //     fmt.Println("Hash of prev block:", block.PreviousHash)
    //     fmt.Println("Data:", string(block.AllData))
    // }
}

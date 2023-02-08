package main

import (
 "context"
 "fmt"
 "math/big"

 "github.com/ethereum/go-ethereum/accounts/abi/bind"
 "github.com/ethereum/go-ethereum/common"
 "github.com/ethereum/go-ethereum/core/types"
 "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
 // Assuming you already have an Ethereum client
 client, err := ethclient.Dial("https://mainnet.infura.io") // connect to an ethereum node
 if err != nil {
  panic(err)
 }

 // Get the private key
 privateKey := "yourprivatekey" // replace with real private key

 // Create an authorized transactor
 auth := bind.NewKeyedTransactor(common.HexToAddress(privateKey2:))

 // Create and sign a transaction
 nonce, err := client.PendingNonceAt(context.Background(), auth.From) // get the account nonce
 if err != nil {
  panic(err)
 }

 // Set the transaction values
 value := big.NewInt(1000000000000000000) // in wei (1 eth)
 gasPrice, err := client.SuggestGasPrice(context.Background())
 if err != nil {
  panic(err)
 }
 gasLimit := uint64(21000) // in units
 toAddress := common.HexToAddress("0xAddress") // address of the recipient

 // Create a new transaction
 tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

 // Sign the transaction
 signedTx, err := auth.Signer(types.HomesteadSigner{}, auth.From, tx)
 if err != nil {
  panic(err)
 }

 // Print the signed transaction
 fmt.Printf("%x\n", signedTx.Encode())
}

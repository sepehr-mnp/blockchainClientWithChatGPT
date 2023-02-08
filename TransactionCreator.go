
package main

import (
    "context"
    "fmt"
 "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/params"
    "github.com/ethereum/go-ethereum/rlp"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/core/types"
)

func main() {
    // Create an authorized transactor and signer
    client, err := ethclient.Dial("http://localhost:8545")
 if err != nil {
        log.Fatal(err)
    }

 fmt.Printf("Connected to Ethereum node\n")

    privateKey, err := crypto.HexToECDSA("yourprivatekey")
    if err != nil {
        log.Fatal(err)
    }

 fmt.Printf("Loaded private key\n")

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(ecdsa.PublicKey)
    if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(publicKeyECDSA)

    nonce,  := client.PendingNonceAt(context.Background(), fromAddress)

    gasPrice,  := client.SuggestGasPrice(context.Background())
    gasLimit := uint64(21000)

    // Build the transaction
    tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

 fmt.Printf("Created transaction\n")

    // Sign the transaction
    signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
    if err != nil {
        log.Fatal(err)
    }
 
 fmt.Printf("Signed transaction\n")

    // Print the transaction
    fmt.Printf("Raw Transaction: 0x%x\n", signedTx.RawSignatureValues())

    // Send the transaction 
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
 }
 
 fmt.Printf("Submitted transaction\n")
 fmt.Printf("Transaction Hash: 0x%x\n", signedTx.Hash().Hex())
}

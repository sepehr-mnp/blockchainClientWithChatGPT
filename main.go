package main

import (
 "context"
 "fmt"
 "log"

 "github.com/ethereum/go-ethereum/accounts/keystore"
 "github.com/ethereum/go-ethereum/common"
 "github.com/ethereum/go-ethereum/crypto"
)

func main() {
 ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
 account, err := ks.NewAccount("password")
 if err != nil {
  log.Fatal(err)
 }
 fmt.Printf("New Ethereum Account: 0x%s\n", account.Address.Hex())
 fmt.Printf("Public key (hex): %s\n", crypto.FromECDSAPub(&account.PrivateKey.PublicKey))
 fmt.Printf("Private key (hex): %s\n", crypto.FromECDSA(account.PrivateKey))
 ctx := context.Background()
 // check if the account was actually persisted / stored
 acc2, err := ks.Find(account.Address)
 if err != nil {
  log.Fatal(err)
 }
 if acc2.Address != account.Address {
  log.Fatalf("Address didn't match: %s != %s\n", acc2.Address.Hex(), account.Address.Hex())
 }
 // unlock the account to enable signing if desired
 err = ks.Unlock(acc2, "password")
 if err != nil {
  log.Fatal(err)
 }
 msg := []byte("This is a sign message, created from Go.")
 sig, err := ks.SignHash(acc2, crypto.Keccak256(msg))
 if err != nil {
  log.Fatal(err)
 }
 recoveredPub, err := crypto.SigToPub(crypto.Keccak256(msg), sig)
 if err != nil {
  log.Fatal(err)
 }
 recoveredAddr := crypto.PubkeyToAddress(*recoveredPub)
 if err != nil {
  log.Fatal(err)
 }
 if recoveredAddr != acc2.Address {
  log.Fatalf("Address didn't match: %s != %s\n", recoveredAddr.Hex(), acc2.Address.Hex())
 }
 fmt.Printf("Successfully recovered address from signature: 0x%s\n", recoveredAddr.Hex())
 // delete account from keystore
 ks.Delete(acc2, "password", &keystore.StandardScryptN, &keystore.StandardScryptP)
 if err != nil {
  log.Fatalf("Expected no error, got %v", err)
 }
 // check if we can find the account in the keystore
 _, err = ks.Find(acc2.Address)
 if err == nil {
  log.Fatalf("Expected error, got none")
 }
}

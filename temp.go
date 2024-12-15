package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// Match Address Block:  PrivateKey: 0000000000000000000000000000000000000000000000000000000000000000 P2PKHBitcoinAddressCompressed: 15wJjXvfQzo3SXqoWGbWZmNYND1Si4siqV, P2PKHBitcoinAddressUncompressed: 16QaFeudRUt8NYy2yzjm3BMvG4xBbAsBFM
func main() {
	spk := "0000000000000000000000000000000000000000000000000000000000000001"
	bspk, err := hex.DecodeString(spk)
	if err != nil {
		panic(err)
	}
	fmt.Println(bspk)
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), bspk)
	//privKey, err := btcec.NewPrivateKey(btcec.S256())
	wif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(wif.String())
}

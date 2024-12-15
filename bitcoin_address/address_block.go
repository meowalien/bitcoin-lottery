package bitcoin_address

import (
	"crypto/ecdsa"
	"encoding/binary"
	"sync"

	//"crypto/rand"
	crypto_rand "crypto/rand"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"math/rand"
)

type AddressBlockGenerator struct {
	prng     *rand.Rand
	prngOnce sync.Once
}

func (a *AddressBlockGenerator) NewAddressBlock() AddressBlock {
	key, err := ecdsa.GenerateKey(btcec.S256(), a.prng)
	if err != nil {
		panic(err)
	}
	privateKey := (*btcec.PrivateKey)(key)

	//privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		panic(err)
	}
	// get compressed public key
	compressedPubKey := privateKey.PubKey().SerializeCompressed()

	// get uncompressed public key
	uncompressedPubKey := privateKey.PubKey().SerializeUncompressed()
	//uncompressedPubKey, compressedPubKey := key.PrivateKeyToPublicKey(privateKey)

	p2pkhAddress := GenerateP2PKHBitcoinAddress(uncompressedPubKey)

	compressedpubkeyP2pkhaddress := GenerateP2PKHBitcoinAddress(compressedPubKey)
	// 17oscJf4PFTnBkJ9E4xjrT1tNEarLA9jiz

	return AddressBlock{
		PrivateKey:                      privateKey,
		P2PKHBitcoinAddressCompressed:   compressedpubkeyP2pkhaddress,
		P2PKHBitcoinAddressUncompressed: p2pkhAddress,
	}
}

func NewAddressBlockGenerator() *AddressBlockGenerator {
	var seed int64
	err := binary.Read(crypto_rand.Reader, binary.LittleEndian, &seed)
	if err != nil {
		panic("failed to seed PRNG: " + err.Error())
	}
	// 5750089927546709938
	fmt.Println("seed: ", seed)
	prng := rand.New(rand.NewSource(seed))
	return &AddressBlockGenerator{
		prng: prng,
	}
}

type AddressBlock struct {
	PrivateKey                      *btcec.PrivateKey
	P2PKHBitcoinAddressCompressed   []byte
	P2PKHBitcoinAddressUncompressed []byte
}

//var (
//	prng     *rand.Rand
//	prngOnce sync.Once
//)

//	func initializePRNG() {
//		var seed int64
//		err := binary.Read(crypto_rand.Reader, binary.LittleEndian, &seed)
//		if err != nil {
//			panic("failed to seed PRNG: " + err.Error())
//		}
//		// 5750089927546709938
//		fmt.Println("seed: ", seed)
//		prng = rand.New(rand.NewSource(seed))
//	}
func (b AddressBlock) String() string {
	//privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), b.PrivateKey)

	wif, err := btcutil.NewWIF(b.PrivateKey, &chaincfg.MainNetParams, true)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(
		"PrivateKey (WIF): %s P2PKHBitcoinAddressCompressed: %s, P2PKHBitcoinAddressUncompressed: %s",
		wif.String(),
		base58.Encode(b.P2PKHBitcoinAddressCompressed),
		base58.Encode(b.P2PKHBitcoinAddressUncompressed),
	)
	// hex.EncodeToString(b.P2PKHBitcoinAddressCompressed) and hex.EncodeToString(b.P2PKHBitcoinAddressUncompressed)
	//return fmt.Sprintf("PrivateKey: %s P2PKHBitcoinAddressCompressed: %s, P2PKHBitcoinAddressUncompressed: %s", hex.EncodeToString(b.PrivateKey), base58.Encode(b.P2PKHBitcoinAddressCompressed), base58.Encode(b.P2PKHBitcoinAddressUncompressed))
}

//func NewAddressBlock() AddressBlock {
//	//prngOnce.Do(initializePRNG)
//	//mockPrivateKey, err := hex.DecodeString("a2071cd615e38dd17d63ba9a26a5b37e966fd495d492388457ea68545b79740d")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//privateKey := mockPrivateKey
//	//privateKey := key.GeneratePrivateKey()
//	key, err := ecdsa.GenerateKey(btcec.S256(), prng)
//	if err != nil {
//		panic(err)
//	}
//	privateKey := (*btcec.PrivateKey)(key)
//
//	//privateKey, err := btcec.NewPrivateKey(btcec.S256())
//	if err != nil {
//		panic(err)
//	}
//	// get compressed public key
//	compressedPubKey := privateKey.PubKey().SerializeCompressed()
//
//	// get uncompressed public key
//	uncompressedPubKey := privateKey.PubKey().SerializeUncompressed()
//	//uncompressedPubKey, compressedPubKey := key.PrivateKeyToPublicKey(privateKey)
//
//	p2pkhAddress := GenerateP2PKHBitcoinAddress(uncompressedPubKey)
//
//	compressedpubkeyP2pkhaddress := GenerateP2PKHBitcoinAddress(compressedPubKey)
//	// 17oscJf4PFTnBkJ9E4xjrT1tNEarLA9jiz
//
//	return AddressBlock{
//		PrivateKey:                      privateKey,
//		P2PKHBitcoinAddressCompressed:   compressedpubkeyP2pkhaddress,
//		P2PKHBitcoinAddressUncompressed: p2pkhAddress,
//	}
//}

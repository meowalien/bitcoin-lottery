package bitcoin_address

import (
	"bitcoin-lottery/key"
	"encoding/hex"

	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

type AddressBlock struct {
	PrivateKey []byte
	//UncompressedPubKey              []byte
	//CompressedPubKey                []byte
	P2PKHBitcoinAddressCompressed   []byte
	P2PKHBitcoinAddressUncompressed []byte
}

func (b AddressBlock) String() string {
	// hex.EncodeToString(b.P2PKHBitcoinAddressCompressed) and hex.EncodeToString(b.P2PKHBitcoinAddressUncompressed)
	return fmt.Sprintf("PrivateKey: %s P2PKHBitcoinAddressCompressed: %s, P2PKHBitcoinAddressUncompressed: %s", hex.EncodeToString(b.PrivateKey), base58.Encode(b.P2PKHBitcoinAddressCompressed), base58.Encode(b.P2PKHBitcoinAddressUncompressed))
}

func NewAddressBlock() AddressBlock {
	//mockPrivateKey, err := hex.DecodeString("a2071cd615e38dd17d63ba9a26a5b37e966fd495d492388457ea68545b79740d")
	//if err != nil {
	//	panic(err)
	//}
	//privateKey := mockPrivateKey
	privateKey := key.GeneratePrivateKey()
	//fmt.Printf("Private Key: %s\n", hex.EncodeToString(privateKey))

	uncompressedPubKey, compressedPubKey := key.PrivateKeyToPublicKey(privateKey)

	p2pkhAddress := GenerateP2PKHBitcoinAddress(uncompressedPubKey)
	//fmt.Printf("P2PKH uncompressed Bitcoin Address: %s\n", p2pkhAddress)

	compressedpubkeyP2pkhaddress := GenerateP2PKHBitcoinAddress(compressedPubKey)
	// 17oscJf4PFTnBkJ9E4xjrT1tNEarLA9jiz
	//fmt.Printf("P2PKH compressed Bitcoin Address: %s\n", base58.Encode(compressedpubkeyP2pkhaddress))

	return AddressBlock{
		PrivateKey: privateKey,
		//UncompressedPubKey:              uncompressedPubKey,
		//CompressedPubKey:                compressedPubKey,
		P2PKHBitcoinAddressCompressed:   compressedpubkeyP2pkhaddress,
		P2PKHBitcoinAddressUncompressed: p2pkhAddress,
	}
}

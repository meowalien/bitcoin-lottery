package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
)

func privateKeyToPublicKey(privateKey []byte) []byte {
	// 加載 secp256k1 橢圓曲線的私鑰
	_, pubKey := btcec.PrivKeyFromBytes(privateKey)

	// 非壓縮格式公鑰（包含完整的 x, y 坐標）
	uncompressedPubKey := pubKey.SerializeUncompressed()

	//// 壓縮格式公鑰（僅包含 x 坐標，並附加 y 坐標的奇偶性）
	//compressedPubKey := pubKey.SerializeCompressed()

	fmt.Printf("Uncompressed Public Key: %s\n", hex.EncodeToString(uncompressedPubKey))
	//fmt.Printf("Compressed Public Key: %s\n", hex.EncodeToString(compressedPubKey))

	return uncompressedPubKey
}

func generatePrivateKey() []byte {
	privateKey := make([]byte, 32)
	_, err := rand.Read(privateKey)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func main() {
	privateKey := generatePrivateKey()
	fmt.Printf("Private Key: %s\n", hex.EncodeToString(privateKey))
	// 生成公鑰
	publicKey := privateKeyToPublicKey(privateKey)
	fmt.Printf("Public Key: %s\n", hex.EncodeToString(publicKey))
}

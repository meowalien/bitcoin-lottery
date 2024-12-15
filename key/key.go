package key

import (
	"crypto/rand"
	"github.com/btcsuite/btcd/btcec/v2"
)

// 生成私鑰
func GeneratePrivateKey() []byte {
	privateKey := make([]byte, 32)
	_, err := rand.Read(privateKey)
	if err != nil {
		panic(err)
	}
	return privateKey
}

// 生成公鑰
func PrivateKeyToPublicKey(privateKey []byte) (uncompressedPubKey []byte, compressedPubKey []byte) {
	_, pubKey := btcec.PrivKeyFromBytes(privateKey)
	uncompressedPubKey = pubKey.SerializeUncompressed()
	compressedPubKey = pubKey.SerializeCompressed()
	return
}

package bitcoin_address

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

// P2PKH 地址生成
func GenerateP2PKHBitcoinAddress(publicKey []byte) []byte {
	sha256Hasher := sha256.New()
	sha256Hasher.Write(publicKey)
	sha256Hash := sha256Hasher.Sum(nil)

	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(sha256Hash)
	publicKeyHash := ripemd160Hasher.Sum(nil)

	versionedPayload := append([]byte{0x00}, publicKeyHash...)

	sha256Hasher.Reset()
	sha256Hasher.Write(versionedPayload)
	firstSHA := sha256Hasher.Sum(nil)
	sha256Hasher.Reset()
	sha256Hasher.Write(firstSHA)
	checksum := sha256Hasher.Sum(nil)[:4]

	finalPayload := append(versionedPayload, checksum...)
	return finalPayload //base58.Encode(finalPayload)
}

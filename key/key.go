package key

import (
	crypto_rand "crypto/rand"
	"math/rand"
	"sync"

	//"encoding/binary"
	//"fmt"
	//"github.com/btcsuite/btcd/btcec/v2"
	"encoding/binary"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	//"math/rand"
	//"sync"
)

var (
	prng     *rand.Rand
	prngOnce sync.Once
)

// // initializePRNG initializes the PRNG with a cryptographically secure seed
func initializePRNG() {
	var seed int64
	err := binary.Read(crypto_rand.Reader, binary.LittleEndian, &seed)
	if err != nil {
		panic("failed to seed PRNG: " + err.Error())
	}
	// 5750089927546709938
	fmt.Println("seed: ", seed)
	prng = rand.New(rand.NewSource(seed))
}

// GeneratePrivateKey generates a 32-byte private key using a PRNG
func GeneratePrivateKey() *btcec.PrivateKey {
	//prngOnce.Do(initializePRNG) // Ensure the PRNG is initialized only once
	//privateKey := make([]byte, 32)
	//prng.Read(privateKey) // Generate the private key
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		panic(err)
	}

	return privKey
}

//import (
//	"crypto/rand"
//	"github.com/btcsuite/btcd/btcec/v2"
//)
//
//// 生成私鑰
//func GeneratePrivateKey() []byte {
//	privateKey := make([]byte, 32)
//	_, err := rand.Read(privateKey)
//	if err != nil {
//		panic(err)
//	}
//	return privateKey
//}

// 生成公鑰
func PrivateKeyToPublicKey(privateKey []byte) (uncompressedPubKey []byte, compressedPubKey []byte) {
	_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)
	uncompressedPubKey = pubKey.SerializeUncompressed()
	compressedPubKey = pubKey.SerializeCompressed()
	return
}

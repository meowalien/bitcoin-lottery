package validater

import (
	"bitcoin-lottery/binary_search"
	"bitcoin-lottery/bitcoin_address"
	"bitcoin-lottery/bloom_filter"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/btcsuite/btcutil/base58"
	"os"
)

type Validater struct {
	bloomFilter                      *bloom.BloomFilter
	addressAfterCleanOnlyAddressFile *os.File
}

func (v *Validater) ValidateAddressBlock(block bitcoin_address.AddressBlock) (match bool) {
	match = v.bloomFilter.Test(block.P2PKHBitcoinAddressCompressed)
	if match {
		match = v.SearchInFile(base58.Encode(block.P2PKHBitcoinAddressCompressed))
		if match {
			return
		}
	}
	match = v.bloomFilter.Test(block.P2PKHBitcoinAddressUncompressed)
	if match {
		match = v.SearchInFile(base58.Encode(block.P2PKHBitcoinAddressUncompressed))
		if match {
			return
		}
	}
	return
}

func (v *Validater) Close() {
	err := v.addressAfterCleanOnlyAddressFile.Close()
	if err != nil {
		panic(err)
	}
}

func (v *Validater) SearchInFile(address string) bool {
	fmt.Println("SearchInFile: ", address)
	match, err := binary_search.BinarySearchOnDisk(v.addressAfterCleanOnlyAddressFile, address)
	if err != nil {
		panic(err)
	}
	return match
}

func NewValidater(bloomFilterName string, addressAfterCleanOnlyAddressFileName string) *Validater {
	file, err := os.Open(addressAfterCleanOnlyAddressFileName)
	if err != nil {
		panic(err)
	}

	return &Validater{
		addressAfterCleanOnlyAddressFile: file,
		bloomFilter:                      bloom_filter.GetBloomFilter(bloomFilterName),
	}
}

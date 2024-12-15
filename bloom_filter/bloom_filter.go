package bloom_filter

import (
	"bufio"
	"github.com/bits-and-blooms/bloom/v3"
	"io"
	"os"
)

func GetBloomFilter(fileName string) *bloom.BloomFilter {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g := &bloom.BloomFilter{}
	r := bufio.NewReader(f)
	_, err = g.ReadFrom(r)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return g
}

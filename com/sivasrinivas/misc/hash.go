package main

import (
	"crypto/sha256"
	"fmt"
	"hash/crc32"
)

func main() {
	//Non-crypto hash function
	h := crc32.NewIEEE()
	h.Write([]byte("siva"))
	fmt.Println("non-crypto hash", h.Sum(nil))
	fmt.Println("non-crypto hash", h.Sum32())

	//Crypto hash function
	s := sha256.New()
	s.Write([]byte("siva"))
	fmt.Println("crypto hash", s.Sum(nil))
}

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	a := sha256.Sum256([]byte("X"))
	b := sha256.Sum256([]byte("x"))
	c := xor(a, b)
	fmt.Printf("%08b\n%08b\n%08b\n", a, b, c)
	fmt.Println(compareSha256(a, b))
}

func compareSha256(a [32]byte, b [32]byte) int {
	sum := 0
	for i := 0; i < 32; i++ {
		sum += sumDifferentBits(a[i], b[i])
	}

	return sum
}

func xor(a [32]byte, b [32]byte) [32]byte {
	var c [32]byte
	for i := 0; i < 32; i++ {
		c[i] = a[i] ^ b[i]
	}
	return c
}

func sumDifferentBits(a byte, b byte) int {
	xbits := a ^ b
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(((xbits & 1) ^ 0))
	}
	return sum
}

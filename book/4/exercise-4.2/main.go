package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {

	algorithm := flag.Int("sha", 256, "The digest you want to use. Supports 256, 384, and 512")
	flag.Parse()
	inputs := flag.Args()

	fmt.Println(inputs)
	for _, v := range inputs {
		var digest string
		if *algorithm == 256 {
			result := sha256.Sum256([]byte(v))
			digest = string(result[:])
		} else if *algorithm == 384 {
			result := sha512.Sum384([]byte(v))
			digest = string(result[:])
		} else if *algorithm == 512 {
			result := sha512.Sum512([]byte(v))
			digest = string(result[:])
		}

		fmt.Printf("%x\n", digest)
	}

}

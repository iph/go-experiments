package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCounts := make(map[string]int)

	in := bufio.NewScanner(bufio.NewReader(os.Stdin))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := in.Bytes()
		wordCounts[string(word)]++
	}

	fmt.Printf("word\tcount\n")
	for c, n := range wordCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}

}

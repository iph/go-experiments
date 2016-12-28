package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Commas: ")
	for _, arg := range os.Args[1:] {
		fmt.Printf("(%s: %s)\n", arg, comma(arg))
	}
}

// Writes what I think is the least allocating, most efficient (for all cases)
// comma separator.
func comma(s string) string {
	n := len(s)
	i := n % 3
	var buff bytes.Buffer

	if n <= 3 {
		return s
	}
	// need to advance token so we don't have an if condition in the for loop below.
	// but weird edge case for things like 100,000, where i == 0, but should be 3.
	if i == 0 {
		buff.WriteString(s[:i])
		i += 3
	}

	buff.WriteString(s[:i])
	for i+3 <= n {
		buff.WriteByte(',')
		buff.WriteString(s[i : i+3])
		i += 3
	}

	return buff.String()
}

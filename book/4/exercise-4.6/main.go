package main

import (
	"fmt"
	"unicode"
)

// Exercise: write a de-dupe of spaces using unicode standards instead of just ASCII standards.
func main() {
	text := "Hello     (◕ ‿ ◕)"
	text = string(uniqueSpaces([]byte(text)))
	fmt.Println(text)
}

func uniqueSpaces(text []byte) []byte {
	out := []byte{}
	for _, char := range text {
		if unicode.IsSpace(rune(char)) {
			lastOutputChar := out[len(out)-1]
			if unicode.IsSpace(rune(lastOutputChar)) {
				continue
			} else {
				out = append(out, char)
			}
		} else {
			out = append(out, char)
		}
	}

	return out
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	args := []string{"/lol/hello"}
	for _, arg := range args {
		fmt.Println(basename(arg))
	}
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

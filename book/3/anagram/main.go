package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(anagram("car", "rab"))
}

func anagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	splitA := strings.Split(a, "")
	splitB := strings.Split(b, "")

	sort.Strings(splitA)
	sort.Strings(splitB)

	for i := 0; i < len(splitA); i++ {
		if strings.Compare(splitA[i], splitB[i]) != 0 {
			return false
		}
	}

	return true
}

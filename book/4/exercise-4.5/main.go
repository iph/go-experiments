package main

import "fmt"

// exercise: delete duplicates from a string using in-place slices.
func main() {
	strings := []string{"a", "a", "b", "c"}
	strings = unique(strings)
	fmt.Println(strings)
}

func unique(strings []string) []string {
	i := 0
	for _, str := range strings {
		if str != strings[i] {
			i++
			strings[i] = str
		}
	}

	return strings[:i+1]
}

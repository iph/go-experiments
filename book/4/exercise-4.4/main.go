package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	rotate(arr)
	fmt.Println(arr)
}

func rotate(rot []int) {
	element := rot[0]

	for i, val := range rot[1:] {
		rot[i] = val
	}

	rot[len(rot)-1] = element
}

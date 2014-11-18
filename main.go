package main

import "fmt"

var UNASSIGNED int = 0

func main() {
	b := Init()
	b = b.SetRow(0, 1, 2)
	b = b.SetRow(1, 4, 5, 6)
	b = b.SetRow(2, 0, 8, 7, 9)
	b = b.SetRow(3, 3)
	possibles := b.PossibleCells(Position{2, 0})
	fmt.Println(possibles)
	b.Print()

}

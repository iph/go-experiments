package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"
)

var UNASSIGNED int = 0

func main() {
	b := BoardParser("test_board.txt")
	b.Print()
	fmt.Println("------")
	b = SudokuSolver(b)
	b.Print()

}

// Simple DFS Solver with Possibility Heuristics to cut down search time to less than a second.
func SudokuSolver(board Board) Board {
	// init
	stack := new(Stack)
	stack.Push(board)

	for stack.Len() > 0 {
		current_board := stack.Pop().(Board)
		if current_board.IsComplete() {
			return current_board
		} else {
			position := current_board.findUnassignedPosition()
			possibles := current_board.PossibleCells(position)
			for _, possible := range possibles {
				new_board := current_board.Assign(position, possible)
				if new_board.PossibleBoard() {
					stack.Push(new_board)
				}
			}
		}
	}

	return board
}

// Input a filename that has the contents like so:
//     8 - - - - 6 3 7 -
//     - - 3 - 4 5 - - 2
//     - - 2 3 - - - - -
//     - - - - - 7 5 - 1
//     - 6 7 - - - 9 3 -
//     9 - 8 4 - - - - -
//     - - - - - 2 1 - -
//     5 - - 8 9 - 7 - -
//     - 8 1 5 - - - - 6
// It disregards everything except numbers, dashes, and newlines. Isn't very fault tolerant...
// TODO: Make fault tolerant!
func BoardParser(filename string) (board Board) {
	board = Init()
	data, _ := ioutil.ReadFile(filename)
	counter := 0
	row := 0
	col := 0
	for counter < len(data) {
		if unicode.IsDigit(rune(data[counter])) {
			board.board[row][col], _ = strconv.Atoi(string(data[counter]))
			col++
		} else if data[counter] == '\n' {
			row++
			col = 0
		} else if data[counter] == '-' {
			board.board[row][col] = UNASSIGNED
			col++
		}
		counter++
	}
	return
}

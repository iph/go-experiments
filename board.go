package main

import "fmt"

type Board struct {
	board [9][9]int
}

type Position struct {
	row int
	col int
}

// initializes a normal sudoku board
func Init() Board {
	new_board := Board{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			new_board.board[i][j] = UNASSIGNED
		}
	}

	return new_board
}

func (b *Board) SetRow(pos int, rows ...int) (new_board Board) {
	new_board = b.Clone()
	for i, row := range rows {
		new_board.board[pos][i] = row
	}

	return
}

func (b *Board) Clone() Board {
	new_board := Board{}
	for i, _ := range b.board {
		for j, _ := range b.board[i] {
			new_board.board[i][j] = b.board[i][j]
		}
	}

	return new_board
}

func (b *Board) Assigned(pos Position, num int) Board {
	new_board := b.Clone()
	new_board.board[pos.row][pos.col] = num
	return new_board
}

func (b *Board) IsComplete() bool {
	for _, row := range b.board {
		for _, cell := range row {
			if cell == UNASSIGNED {
				return false
			}
		}
	}

	return true
}

func (b *Board) uniqueColumns(possible_num int, column int) bool {
	for _, row := range b.board {
		for col, cell := range row {
			if col == column && possible_num == cell {
				return false
			}
		}
	}
	return true
}

func (b *Board) uniqueRows(possible_num int, row int) bool {
	for _, cell := range b.board[row] {
		if possible_num == cell {
			return false
		}
	}
	return true
}

func (b *Board) uniqueBox(possible_num int, pos Position) bool {
	// check the box...
	starting_row := pos.row / 3
	starting_col := pos.col / 3
	ending_row := starting_row + 3
	ending_col := starting_col + 3

	for i := starting_row; i < ending_row; i++ {
		for j := starting_col; j < ending_col; j++ {
			if b.board[i][j] == possible_num {
				return false
			}
		}
	}
	return true
}

func (b *Board) PossibleCells(pos Position) (possibles []int) {
	possibles = []int{}
	for i := 1; i <= 9; i++ {
		if b.uniqueColumns(i, pos.col) && b.uniqueRows(i, pos.row) && b.uniqueBox(i, pos) {
			possibles = append(possibles, i)
		}
	}

	return possibles
}

func (b *Board) Print() {
	for _, row := range b.board {
		for _, cell := range row {
			if cell == UNASSIGNED {
				fmt.Print(" - ")
			} else {
				fmt.Print(" ", cell, " ")
			}

		}
		fmt.Println()
	}
}

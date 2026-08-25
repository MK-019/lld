package main

// Represents the Tic-Tac-Toe board
type Board struct {
	grid [3][3]string
}

// type Player struct {
// 	Player string
// }

// func NewPlayer()

// Initialize an empty Tic-Tac-Toe board
func NewBoard() *Board {
	board := &Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.grid[i][j] = " "
		}
	}
	return board
}

type BoardInterface interface {
	Display()
	IsGameOver()
	MoveInterface
}

type MoveInterface interface {
	MakeMove(row int, col int, player string) bool
}

package main

import (
	"fmt"
)

// Constants to represent players
const (
	PlayerX = "X"
	PlayerO = "O"
)

var (
	maxMoves = 9
)



// Display the Tic-Tac-Toe board
func (board *Board) Display() {
	fmt.Println("    1    2   3")
	fmt.Println("  -------------")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d |", i+1)
		for j := 0; j < 3; j++ {
			fmt.Printf(" %s |", board.grid[i][j])
		}
		fmt.Println("\n  -------------")
	}
}

// Make a move on the board
func (board *Board) MakeMove(row, col int, player string) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || board.grid[row][col] != " " {
		return false
	}
	board.grid[row][col] = player
	return true
}

// Check if the game is over
func (board *Board) IsGameOver() bool {
	// Check rows, columns, and diagonals for a win
	for i := 0; i < 3; i++ {
		if board.grid[i][0] != " " && board.grid[i][0] == board.grid[i][1] && board.grid[i][1] == board.grid[i][2] {
			// fmt.Println("from rows")
			return true
		}
		if board.grid[0][i] != " " && board.grid[0][i] == board.grid[1][i] && board.grid[1][i] == board.grid[2][i] {
			// fmt.Println("from col")
			return true
		}
	}
	if board.grid[0][0] != " " && board.grid[0][0] == board.grid[1][1] && board.grid[1][1] == board.grid[2][2] {
		// fmt.Println("from diagonal")
		return true
	}
	if board.grid[0][2] != " " && board.grid[0][2] == board.grid[1][1] && board.grid[1][1] == board.grid[2][0] {
		// fmt.Println("from reverse diagonal")
		return true
	}

	// Check for a draw
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.grid[i][j] == " " {
				return false
			}
		}
	}

	return false
}

func main() {
	board := NewBoard()
	currentPlayer := PlayerX

	board.Display()
	for {
		fmt.Printf("Player %s's turn. Enter row (1-3) and column (1-3): ", currentPlayer)
		var row, col int
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil || row < 1 || row > 3 || col < 1 || col > 3 {
			fmt.Println("Invalid input. Try again.")
			continue
		}
		row--
		col--

		if board.MakeMove(row, col, currentPlayer) {
			maxMoves--
			if maxMoves >= 0 {
				if board.IsGameOver() {
					fmt.Printf("****Player %s wins! ****\n", currentPlayer)
					board.Display()
					break
				}
				board.Display()
				if currentPlayer == PlayerX {
					currentPlayer = PlayerO
				} else {
					currentPlayer = PlayerX
				}
			} else {
				fmt.Println("Match Draw. Please Restart.")
				break
			}
		} else {
			fmt.Println("Invalid move. Try again.")

		}
		if maxMoves == 0 {
			fmt.Println("Match Draw. Please Restart.")
			break
		}
	}
}

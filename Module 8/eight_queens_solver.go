package sprint

/*
Eight Queens Solver

Learning Objectives:
*	Gain proficiency in implementing recursive functions to solve complex problems.
*	Practice string manipulation and output formatting.

Instructions
You're tasked with creating a function that returns the solutions to the eight queens puzzle. Use recursivity to solve this problem.
The function returns the solutions as a string, with each solution on a single line. The index of the placement of a queen starts at 1 and reads from left to right, with each digit representing the position for each column. The solutions should be in ascending order.

Allowed packages
For this task you're allowed to use:
*	strconv
*	strings
*/

import (
	"strconv"
	"strings"
)

const N = 8

func isSafe(board [N]int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || board[i]-i == row-col || board[i]+i == row+col {
			return false
		}
	}
	return true
}

func solveNQueens(board [N]int, col int, solutions *[]string) {
	if col == N {
		var solution strings.Builder
		for i := 0; i < N; i++ {
			solution.WriteString(strconv.Itoa(board[i] + 1))
		}
		*solutions = append(*solutions, solution.String())
		return
	}

	for row := 0; row < N; row++ {
		if isSafe(board, row, col) {
			board[col] = row
			solveNQueens(board, col+1, solutions)
		}
	}
}

func EightQueensSolver() string {
	var solutions []string
	var board [N]int
	solveNQueens(board, 0, &solutions)
	println(len(solutions))
	return strings.Join(solutions, "\n")
}

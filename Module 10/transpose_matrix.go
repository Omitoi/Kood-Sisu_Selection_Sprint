package sprint

/*
Transpose Matrix

Learning Objectives:
*	This task will provide practical experience in working with 2D arrays and transforming them by transposition.
*	Transposing a matrix involves using nested loops to iterate through rows and columns, which is a fundamental concept in programming.
*	Understanding how to manipulate data structures, such as a 2D array, is essential for various data processing tasks.
*	Transposing a matrix is a problem-solving task that requires an understanding of matrix operations and algorithmic thinking, skills valuable in software development.

Instructions
Write a Go function that takes a matrix represented as a 2D array and returns the transposed matrix. Transposing a matrix involves swapping its rows and columns.
The function should take the original matrix and return the resulting transposed 2D array.
*/

func TransposeMatrix(matrix [][]int) [][]int {
	// Get the number of rows and columns in the original matrix
	rows := len(matrix)
	if rows == 0 {
		return [][]int{} // Return an empty matrix if input is empty
	}
	cols := len(matrix[0])

	// Create a new matrix for the transposed version with swapped dimensions
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	// Fill the transposed matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

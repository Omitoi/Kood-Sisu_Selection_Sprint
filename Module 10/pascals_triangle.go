package sprint

/*
Pascal's Triangle

Learning Objectives:
*	This task will provide practical experience in working with 2D arrays.
*	You'll use loops and iteration to generate the rows and elements.
*	Implementing Pascal's triangle involves mathematical computations to calculate binomial coefficients.
*	Creating a 2D array involves working with nested arrays, which is a useful skill in handling structured data.
*	Generating Pascal's triangle is a problem-solving task that requires understanding mathematical patterns and algorithmic thinking, skills valuable in software development.

Instructions
Write a Go function that takes an integer as input and returns a 2D integer array representing Pascal's triangle up to the specified number of levels. 
Pascal's triangle is a triangular array of binomial coefficients where each number is the sum of the two directly above it. 
The function should generate the specified number of levels of Pascal's triangle.
*/

func PascalsTriangle(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

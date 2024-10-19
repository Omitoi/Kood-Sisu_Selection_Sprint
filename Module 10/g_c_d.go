package sprint

/*
Greatest Common Divisor

Learning Objectives:
*	You'll likely use conditional statements and loops to implement algorithms for finding the GCD, helping reinforce your understanding of conditional logic.
*	Finding the greatest common divisor is a problem-solving task that encourages algorithmic thinking and an understanding of numerical operations, which are valuable in software development.
*	You may explore and implement various algorithms to find the GCD, which can teach you about algorithmic efficiency and optimization.

Instructions
Write a Go function that takes two integers as input and returns their greatest common divisor (GCD). 
The GCD is the largest positive integer that divides both of the input integers without leaving a remainder.
*/

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b) 
}
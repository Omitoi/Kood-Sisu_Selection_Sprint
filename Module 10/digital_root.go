package sprint

/*
Digital Root

Learning Objectives:
*	The task includes performing basic arithmetic operations like addition and division, which are essential in programming and mathematics.
*	This task can be tackled very well both recursively and iteratively. Try solving it both ways.
*	Solving the digital root problem requires logical thinking and an understanding of numerical operations, skills that are valuable in software development.

Instructions
Write a Go function that solves the digital root problem. The function should take an integer as input and return an integer. The digital root is the recursive sum of the digits of a number until a single-digit result is achieved.
For example, the digital root of 9875 is 2 because 9 + 8 + 7 + 5 = 29, and 2 + 9 = 11, and finally 1 + 1 = 2.
*/

func DigitalRoot(n int) int {
	for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10 // Add the last digit to sum
			n /= 10       // Remove the last digit
		}
		n = sum // Update n to be the sum of its digits
	}
	return n // Return the single-digit result
}

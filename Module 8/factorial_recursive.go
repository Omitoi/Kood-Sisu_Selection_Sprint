package sprint

/*
Factorial Recursive

Learning Objectives:
*	Gain proficiency in writing recursive functions to solve mathematical problems.
*	Understand the concept of factorial and how to calculate it.

Instructions
You're tasked with creating a recursive function that calculates the factorial of an integer passed as a parameter. 
Make sure to handle errors, returning 0 for non-possible values or overflows. NB! You should achieve this without using a for loop.
*/

func FactorialRecursive(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

package sprint

/*
Factorial Iterative

Learning Objectives:
*	Gain proficiency in writing iterative functions to solve mathematical problems.
*	Understand the concept of factorial and how to calculate it.

Instructions
You're tasked with creating an iterative function that calculates the factorial of an integer passed as a parameter.
Make sure to handle errors, returning 0 for non-possible values or overflows.
*/

func FactorialIterative(n int) int {
	if n < 0 {
		return 0
	}
	for i := n - 1; i > 0; i-- {
		n *= i
	}
	return n
}

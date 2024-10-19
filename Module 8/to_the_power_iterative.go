package sprint

/*
To the Power Iterative

Learning Objectives:
*	Gain proficiency in writing iterative functions to solve mathematical problems.
*	Understand how to calculate the power of a number.

Instructions
You're tasked with creating an iterative function that calculates the power of an integer n to the given power. Handle negative powers by returning 0.
*/

func ToThePowerIterative(n int, power int) int {
	if power < 0 {
		return 0
	}
	res := n
	for i := 0; i < power-1; i++ {
		res *= n
	}
	return res
}

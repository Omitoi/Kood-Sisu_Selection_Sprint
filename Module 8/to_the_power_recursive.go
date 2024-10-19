package sprint

/*
To the Power Recursive

Learning Objectives:
Gain proficiency in writing recursive functions to solve mathematical problems.
Understand how to calculate the power of a number.

Instructions
You're tasked with creating an recursive function that calculates the power of an integer n to the given power. 
Handle negative powers by returning 0. NB! You should achieve this without using a for loop.
*/

func ToThePowerRecursive(n int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	}
	return n * ToThePowerRecursive(n, power-1)
}

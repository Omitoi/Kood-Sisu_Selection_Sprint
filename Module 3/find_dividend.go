package sprint

/*
Find Dividend

Learning Objectives:
*	This task will provide practical experience in using loops to iterate through a range of numbers, which is a fundamental programming concept.
*	You'll apply conditional logic within the loop to determine if a number is divisible by the given divisor.
*	Dealing with scenarios where there might be no divisible number and returning -1 demonstrates how to handle edge cases and provide appropriate outputs.

Instructions
Write a Go function that takes three integers as input: from, to, and divisor.
The function should loop through the numbers from from (inclusive) to to (exclusive) and return the first number in that range that is divisible by the divisor.
If there is no such number, the function should return -1.
*/

func FindDividend(from, to, divisor int) int {
	for i := from; i < to; i++ {
		if i%divisor == 0 {
			return i //this will end the for cycle and output the first i it finds in the steps
		}
	}
	return -1
}

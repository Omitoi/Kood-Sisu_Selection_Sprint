package sprint

/*
Count Divisible

Learning Objectives:
*	This task will teach you how to handle exceptions or edge cases where certain input conditions must be met.
*	Reinforces your ability to use loops and iteration to traverse a range of numbers, taking into account a step value.
*	You will apply conditional logic within the loop.

Instructions
Write a Go function that takes four integers as input: from, to, step, and divisor. The function should check for exceptions and return 0 if any of the following conditions are met:
*	step is negative or zero.
*	divisor is zero.
	Otherwise, it should loop through the range of numbers from from (inclusive) to to (exclusive), checking every step-th element, and return the count of elements among these that are divisible by the divisor.
	For the example below the numbers checked would be (divisible in bold): 5, 7, 9, 11, 13, 15.
*/

func CountDivisible(from, to, step, divisor int) int {
	if step <= 0 || divisor == 0 {
		return 0
	}
	counter := 0
	for i := from; i < to; i += step {
		if i%divisor == 0 {
			counter++
		}
	}
	return counter
}

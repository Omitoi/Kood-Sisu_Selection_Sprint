package sprint

/*
Accumulate

Learning Objectives:
*	Develop the ability to use loops, such as for loops, to iterate through sequences of data, perform calculations, and manage program flow efficiently.
*	Gain proficiency in using conditional statements, like if statements, to make decisions and handle different cases, ensuring that your code responds appropriately to varying input or conditions.

Instructions
Create a Go function that takes a non-negative integer n. If n is positive, return the sum of all the digits from 0 to n, including n itself. If n is negative, return 0.
For example, in case of 4, the sum would be 0 + 1 + 2 + 3 + 4 = 10.
*/

func Accumulate(n int) int {
	var res int
	if n <= 0 {
		return 0
	} else {
		for i := 0; i <= n; i++ {
			res += i
		}
		return res
	}
}

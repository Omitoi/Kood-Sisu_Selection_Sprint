package sprint

/*
Generate Range

Learning Objectives:
*	Practice using conditional statements to make decisions and handle different cases within your code, ensuring that it responds appropriately to varying input or conditions.
*	Gain proficiency in manipulating data to create sequences, ranges, or specific sets of values to meet the requirements of a task.
*	Learn to work with data structures like slices or arrays to represent and manage collections of data efficiently.

Instructions
Create a Go function that takes two integers, min and max, as input. The function should return a slice of integers containing all values between min (inclusive) and max (exclusive) using the make function to create the slice.
If min is greater than or equal to max, the function should return a nil slice.
*/
func GenerateRange(min, max int) []int {
	res := make([]int, 0)
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}

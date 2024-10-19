package sprint

/*
Bulk Atoi

Learning Objectives:
*	You'll use the StrToInt function you previously created as a part of a larger task, demonstrating how functions can be composed for more complex functionality.
*	Creating a function for this specific task promotes modular programming, a best practice for code organization and reusability.
*	Understanding how to manipulate data structures, such as arrays, is essential for various data processing tasks.

Instructions
Write a Go function that takes an array of strings, applies the StrToInt function you wrote earlier to every element in the array, and returns the resulting integer values as a new array.
Note that you cannot call the StrToInt function from the package in our current learning environment, you have to copypaste the function instead.

*/

func BulkAtoi(arr []string) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		res = append(res, StrToInt(arr[i]))
	}
	return res
}

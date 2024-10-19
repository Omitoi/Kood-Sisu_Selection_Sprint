package sprint

/*
Negative or Not?

Learning Objectives:
*	Develop the skill of handling and evaluating different data types, such as integers and boolean values, to make informed decisions in your code.

Instructions
Create a Go function that takes an integer as input and returns a boolean value (true or false) to indicate whether the input integer is negative or not.
*/

func IsNegative(n int) bool {
	if n > 0 {
		return false
	} else {
		return true
	}
}

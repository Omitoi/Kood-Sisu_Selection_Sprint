package sprint

/*
Do Operation

Learning Objectives:
*	Gain proficiency in creating functions to perform mathematical operations.
*	Learn to handle and validate user inputs, such as operators and values.
*	Understand error handling for division by 0 and invalid operations.

Instructions
You're tasked with creating a function called Doop. This function takes three parameters:
*	A value (integer).
*	An operator, which can be one of the following: +, -, /, *, %.
*	Another value (integer).
	In case of an invalid operator, invalid values, or incorrect number of arguments, the function returns 0. 
	The program must also handle modulo and division by 0.
*/

func Doop(a int, op string, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		if a == 0 || b == 0 { //Handling 0 in input, both cases give out 0 no matter the other number
			return 0
		}
		return a / b
	case "*":
		return a * b
	case "%":
		return a % b
	}
	return 0
}

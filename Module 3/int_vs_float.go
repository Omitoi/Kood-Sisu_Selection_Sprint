package sprint

/*
Int vs Float

Learning Objectives:
*	You'll learn how to convert data types.
*	This task reinforces your understanding of conditional logic (if-else) for decision-making.
*	Developing a function for comparing integers and floats will enhance your skills in numerical comparisons and handling different data types.

Instructions
Write a Go function that takes an integer and a float as input and returns a string. 
The function should compare the integer to the float and return one of the following: Integer if the integer is larger, Float if the float is larger, or Same if they are equal.
*/

func IntVsFloat(i int, f float32) string {
	conv := float32(i)
	if conv < f {
		return "Float"
	} else if conv > f {
		return "Integer"
	} else {
		return "Same"
	}
}

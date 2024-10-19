package sprint

/*
String Concatenation

Learning Objectives:
*	Gain proficiency in working with strings in Go, including techniques to concatenate and manipulate them effectively.
*	Develop the ability to design and create a Go function that accepts input parameters, performs string operations, and returns a meaningful result.

Instructions
Create a function that takes two strings and a delimiter. The function should combine the two strings, placing the delimiter between them, and return the combined result as a single string.
*/

func StrConcat(s1, s2, delim string) string {
	return s1 + delim + s2
}

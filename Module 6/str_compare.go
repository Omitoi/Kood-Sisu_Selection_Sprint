package sprint

/*
Str Compare

Learning Objectives:
*	Enhance your problem-solving skills and the ability to design algorithms for comparing and evaluating data efficiently.
*	Gain proficiency in working with data, specifically strings, and using comparison methods to understand their relationships.

Instructions
Create a Go function that mimics the behavior of the Compare function. It takes two strings, a and b, as input and returns an integer.
The function should compare the two strings and return:
* 0 if the strings are equal,
* -1 if a comes before b in lexicographic order,
* 1 if a comes after b in lexicographic order.
*/

func StrCompare(a, b string) int {
	lenA := len(a)
	lenB := len(b)

	for i := 0; i < lenA && i < lenB; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if lenA < lenB {
		return -1
	} else if lenA > lenB {
		return 1
	}
	return 0
}

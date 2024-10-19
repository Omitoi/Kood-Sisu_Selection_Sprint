package sprint

/*
Str Reverse

Learning Objectives:
*	Develop proficiency in manipulating strings, such as reversing their content, which is a common operation in text processing.

Instructions
Create a Go function that takes a string and returns the reversed version of that string.
*/

func StrReverse(s string) string {
	res := ""
	for _, r := range s {
		res = string(r) + res
	}
	return res
}

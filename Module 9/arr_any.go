package sprint

/*
Array Any

Learning Objectives:
*	Gain proficiency in working with higher-order functions in Go.
*	Understand how to apply a given function to each element of a slice.
*	Practice using the result of applying a function to make a decision.

Instructions
You're tasked with creating a function ArrAny that returns true for a string slice if at least one element returns true when passed through the f function.

Required functions
For this task you will need to have in the same file some functions you have previously created:
*	IsLower
*	IsNumeric
You will also need to implement two new functions:
*	IsUpper 		-Returns true if the string given contains only uppercase characters, and false otherwise.
*	IsAlphanumeric 	-Returns true if the string contains only alphanumeric characters, and false otherwise.
*/

func ArrAny(f func(string) bool, a []string) bool {
	for _, r := range a {
		if f(r) {
			return true
		}
	}
	return false
}

func IsUpper(s string) bool {
	isIt := true
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			isIt = false
		}
	}
	return isIt
}

func IsAlphanumeric(s string) bool {
	isIt := true
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			isIt = false
		}
	}
	return isIt
}

func IsNumeric(s string) bool {
	res := true
	for i := 0; i < len(s); i++ {
		if s[i] < 48 || s[i] > 57 {
			res = false
		}
	}
	return res
}

func IsLower(s string) bool {
	res := true
	for i := 0; i < len(s); i++ {
		if s[i] < 97 || s[i] > 122 {
			res = false
		}
	}
	return res
}

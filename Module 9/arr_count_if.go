package sprint

/*
Array Count If

Learning Objectives:
*	Gain proficiency in working with higher-order functions in Go.
*	Understand how to apply a given function to each element of a slice and count the results.
*	Practice using the result of applying a function to calculate a count.

Instructions
You're tasked with creating a function ArrCountIf that returns the number of elements in a string slice that return true when passed through the function f.

Required functions
For this task you will need to have in the same file some functions you have previously created:
*	IsLower
*	IsUpper
*	IsNumeric
*	IsAlphanumeric
*/

func ArrCountIf(f func(string) bool, a []string) int {
	counter := 0
	for _, r := range a {
		if f(r) {
			counter++
		}
	}
	return counter
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

func IsLower(s string) bool {
	res := true
	for i := 0; i < len(s); i++ {
		if s[i] < 97 || s[i] > 122 {
			res = false
		}
	}
	return res
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

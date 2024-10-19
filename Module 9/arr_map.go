package sprint

/*
Array Map
Learning Objectives:
*	Gain proficiency in working with higher-order functions in Go.
*	Understand how to apply a given function to each element of a slice.
*	Practice working with different data types and return values in functions.

Instructions
You're tasked with creating a function ArrMap that applies a function of type func(int) bool to each element of an int slice and returns a slice of the resulting return values.

Required functions
For this task you will need to have in the same file some functions you have previously created:
*	IsNegative
*	IsPrime
*/

func ArrMap(f func(int) bool, a []int) []bool {
	var res []bool
	for _, r := range a {
		if f(r) {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}

func IsNegative(n int) bool {
	if n > 0 {
		return false
	} else {
		return true
	}
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

package sprint

/*
Advanced Word Sorting

Learning Objectives:
*	Gain proficiency in creating functions to sort data structures with custom comparison functions.
*	Understand how to use custom comparison functions for sorting.
*	Practice working with different data types and return values in functions.

Instructions
You're tasked with creating a function called AdvancedSortWordArr that applies a sorting function f to a slice of strings a and returns the sorted slice.

Required functions
For this task you will need to have in the same file some functions you have previously created:
*	StrCompare
*/

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if f(a[j], a[j+1]) == 1 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

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

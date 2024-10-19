package sprint

/*
Sort Words

Learning Objectives:
*	Gain proficiency in creating functions to sort data structures.
*	Understand the concept of sorting based on ASCII values.
*	Practice working with different data types and return values in functions.

Instructions
You're tasked with creating a function called SortWordArr that sorts a string slice in ascending order based on ASCII values.
*/

func SortWordArr(a []string) []string {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

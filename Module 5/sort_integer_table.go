package sprint

/*
Sort Integers

Learning Objectives:
*	Develop proficiency in sorting data structures, in this case, a slice of integers.
*	Gain experience in modifying data structures in place to achieve the desired result.

Instructions
Write a function that sorts a slice of integers in ascending order.
*/
func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	return table
}

package sprint

/*
Remove Duplicates

Learning Objectives:
*	This task provides practical experience in working with arrays and performing operations to filter out duplicate elements.
*	You'll learn how to maintain the order of elements in an array while removing duplicates, a valuable skill for data manipulation.
*	You'll use iteration to scan the array, compare elements, and identify duplicates for removal.
*	Removing duplicates from an array while preserving order is a problem-solving task that encourages logical thinking and algorithm design, skills that are important in software development.

Instructions
Write a Go function that takes an array of integers as input and returns the same array with duplicate elements removed, preserving the order of the remaining elements.
For example, if the input array is [3, 5, 2, 3, 8, 5, 2], the function should return [3, 5, 2, 8].
*/

func RemoveDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	var res []int
	for _, r := range arr {
		if seen[r] {

		} else {
			res = append(res, r)
			seen[r] = true
		}
	}
	return res
}

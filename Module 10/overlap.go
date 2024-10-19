package sprint

/*
Overlap

Learning Objectives:
*	This task will provide practical experience in working with arrays.
*	You'll use a sorting algorithm to order the common elements, offering insight into how different sorting algorithms work.
*	You'll use iteration to scan both arrays and compare elements to identify common ones.
*	The task requires returning an array of common elements, possibly with multiple occurrences, demonstrating data extraction skills in programming.
*	Finding common elements in two arrays and presenting them in sorted order is a problem-solving task that encourages logical thinking and algorithm design, important skills in software development.

Instructions
Write a Go function that takes two arrays of integers and returns an array containing the elements that are common in both input arrays, sorted in ascending order. 
If an element occurs multiple times in both arrays, it should be included in the result array as many times as it occurs.
*/

func Overlap(arr1, arr2 []int) []int {
	// Maps to count occurrences
	count1 := make(map[int]int)
	count2 := make(map[int]int)

	// Count occurrences in arr1
	for _, num := range arr1 {
		count1[num]++
	}

	// Count occurrences in arr2
	for _, num := range arr2 {
		count2[num]++
	}

	// Result array to store common elements
	result := []int{}

	// Find common elements and add them to the result based on minimum occurrence
	for num, count := range count1 {
		if count2[num] > 0 {
			// Take the minimum occurrence from both arrays
			minCount := min(count, count2[num])
			for i := 0; i < minCount; i++ {
				result = append(result, num)
			}
		}
	}

	// Manually sort the result array using Bubble Sort
	bubbleSort(result)

	return result
}

// Helper function to return the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Bubble sort implementation to sort the result array
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

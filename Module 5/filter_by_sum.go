package sprint
/*
Filter by Sum

Learning Objectives:
*	You'll use conditional statements and loops to filter subarrays based on the sum of their elements.
*	Filtering out subarrays based on a specific condition is a valuable skill for data processing and manipulation.
*	Filtering subarrays based on their sum is a problem-solving task that encourages logical thinking and understanding of array manipulation, skills valuable in software development.

Instructions
Write a Go function that takes a 2D array of integers and an integer value. 
The function should filter out all subarrays from the array (2D) in which the sum of elements is lower than the given value. The resulting modified 2D array should be returned.

*/
func FilterBySum(arr [][]int, limit int) [][]int {
	var res [][]int
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := 0; j < len(arr[i]); j++ {
			sum += arr[i][j]
		}
		if sum >= limit {
			res = append(res, arr[i])
		}
	}
	if res == nil {
		return [][]int{}
	}
	return res
}

package sprint

/*
Longest Climb

Learning Objectives:
*	This task will provide practical experience in working with arrays and selecting contiguous subsequences based on certain conditions.
*	Designing an algorithm to find the longest contiguous increasing subarray in an array requires problem-solving skills and algorithmic thinking.
*	The task requires returning a subarray of elements that form the longest contiguous increasing subarray, which is useful for data extraction in real-world applications.
*	Identifying the longest contiguous increasing subarray within an array is a problem-solving task that encourages logical thinking and algorithm design, important skills in software development.

Instructions
Write a Go function that takes an array of integers and returns the longest contiguous subarray in the array in which the elements are increasing. 
The function should return the contiguous subarray.
*/

func LongestClimb(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	longest := []int{}
	current := []int{arr[0]}
	last := arr[0]
	for _, r := range arr[1:] {
		if last < r {
			current = append(current, r)
		} else {
			if len(current) > len(longest) {
				longest = current
			}
			current = []int{r}
		}
		last = r
	}
	if len(current) > len(longest) {
		longest = current
	}
	return longest
}

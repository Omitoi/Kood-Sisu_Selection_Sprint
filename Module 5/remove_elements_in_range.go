package sprint
/*
Remove Elements in Range

Learning Objectives:
*	You'll need to apply conditional statements to handle indices that can be negative or larger than the array length.
*	Handling indices that are out of bounds or negative is a valuable skill in addressing edge cases in programming.
*	Removing elements from an array within a specified range is a problem-solving task that encourages logical thinking and understanding of array manipulation, skills valuable in software development.

Instructions
Write a Go function that takes an array of float64, two indices, and removes elements that lie between these indices from the array. The lower index should be removed, and the upper index should be kept. The function should return the resulting modified array. 
The indices can be negative or larger than the length of the array, but the function should still remove the elements that fall within the specified range. The indices might also be in wrong order.
*/
func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	var res []float64
	if from > to {
		from, to = to, from
	}
	if from < 0 {
		from = 0
	}
	if to < 0 {
		to = 0
	}
	if to > len(arr) {
		to = len(arr)
	}
	if from > len(arr) {
		from = len(arr)
	}
	for i := 0; i < len(arr); i++ {
		if i < from || i >= to {
			res = append(res, arr[i])
		} else {

		}
	}
	if res == nil {
		return []float64{}
	}
	return res
}

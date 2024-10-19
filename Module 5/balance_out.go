package sprint

/*
Balance Out

Learning Objectives:
*	You'll need to apply conditional statements to calculate and add booleans to balance the counts.
*	Balancing the count of booleans in an array is a problem-solving task that encourages logical thinking and understanding of array manipulation, skills valuable in software development.
*	Working with boolean values and modifying an array of booleans demonstrates an understanding of boolean logic, a fundamental concept in programming.

Instructions
Write a Go function that takes an array of booleans and adds the minimum number of booleans necessary so that the count of true and false values in the array is equal.
The function should return the resulting modified array. The order of the elements must remain the same and new elements should be added at the end of the array.
*/

func BalanceOut(arr []bool) []bool {
	counter := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] {
			counter++
		}
		if !arr[i] {
			counter--
		}
	}
	println(counter)
	if counter < 0 {
		for i := counter; i != 0; i++ {
			arr = append(arr, true)
		}
	}
	if counter > 0 {
		for i := counter; i != 0; i-- {
			arr = append(arr, false)
		}
	}
	return arr
}

package sprint

import "fmt"

/*
Combinations

Learning Objectives:
*	Develop the ability to recognize and generate specific patterns in data, in this case, ascending triplets of digits.
*	Gain proficiency in using loops to generate and iterate through sequences of data efficiently.
*	Learn how to format strings by creating and concatenating unique triplets, understanding how to separate them as required.
*	Practice using conditional statements to filter out specific combinations or patterns that don't meet the task's criteria.

Instructions
Create a Go function that generates a series of unique triplets of digits. Each triplet should consist of different digits, and they should be in ascending order, from the smallest to the largest. The triplets should be separated by a comma and a space.
For example, the function should return triplets like 012, 013, 014, 015.... Avoid combinations with all the same digits, like 000 or 999, and exclude triplets in descending order, like 987.

Allowed packages
For this task you're allowed to use:
fmt
*/

func Combinations() string {
	var res string
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				if i == 7 && j == 8 && k == 9 {
					res += fmt.Sprintf("%d%d%d", i, j, k)
				} else {
					res += fmt.Sprintf("%d%d%d, ", i, j, k)
				}
			}
		}
	}
	return res
}

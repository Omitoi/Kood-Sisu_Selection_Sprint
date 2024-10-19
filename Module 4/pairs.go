package sprint

/*
Pairs

Learning Objectives:
*	Develop skills in working with strings, including formatting and concatenating strings to create structured outputs.
*	Gain expertise in using loops to generate pairs or sequences of data and efficiently iterate through them.
*	Enhance problem-solving skills by designing a solution to generate all possible pairs, demonstrating the ability to break down complex tasks into manageable steps.
*	Learn how to represent and format data effectively, including considerations for padding numbers or applying other transformations as required.

Instructions
Create a Go function that finds all possible combinations of two two-digit numbers. The pairs should be in ascending order, and each number should be padded with a leading zero if it's less than 10. The pairs should be separated by a comma and a space. Each number within a pair should be separated by a space. The function should return a string containing these pairs.

Allowed packages
For this task you're allowed to use:
fmt
*/
import "fmt"

func Pairs() string {
	var res string
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if i != 98 || j != 99 {
				if i != j {
					res += fmt.Sprintf("%02d %02d", i, j)
				}
				res += ", "
			} else {
				res += fmt.Sprintf("%02d %02d", i, j)
				return res
			}
		}
	}
	return ""
}

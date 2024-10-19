package sprint

/*
Substring Index

Learning Objectives:
*	Enhance your capacity to think algorithmically, enabling you to devise step-by-step procedures to solve a wide range of computational problems.
*	Gain proficiency in manipulating data, allowing you to perform operations like searching, extracting, and indexing data elements in various data structures.
*	Practice using conditional statements to make decisions in your code, ensuring that it responds appropriately to different cases and conditions.

Instructions
Create a Go function that behaves like the Index function. It takes two strings as input: s and toFind. 
The function should find the index of the first occurrence of toFind in s and return that index as an integer.
If toFind is not present in s, the function should return -1.
*/

func SubstrIndex(s string, toFind string) int {
	toFindlen := len(toFind)
	for i := 0; i < len(s); {
		if i <= len(s)-toFindlen && s[i:i+toFindlen] == toFind {
			return i
		} else {
			i++
		}
	}
	if toFind == "" {
		return 0
	}
	return -1
}

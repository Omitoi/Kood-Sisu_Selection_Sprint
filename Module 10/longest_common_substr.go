package sprint

/*
Longest Common Substring

Learning Objectives:
*	This task will provide practical experience in working with strings and manipulating them to find common substrings.
*	Designing an efficient algorithm to find the longest common substring in two strings requires problem-solving skills and algorithmic thinking.
*	You will use iteration to compare substrings and identify the longest common substring while considering edge cases.
*	Addressing the case where there are multiple substrings of the same length and determining which one to return is part of the problem-solving process.
*	Identifying and extracting the longest common substring from two strings is a problem-solving task that encourages logical thinking and algorithm design, important skills in software development.

Instructions
Write a Go function that takes two strings as input and finds the longest common substring that occurs in both strings.
The function should return the substring that occurs earlier in the first string passed if there are multiple substrings of the same length.
*/

func LongestCommonSubstr(str1, str2 string) string {
	// Get the lengths of both strings
	m, n := len(str1), len(str2)

	// Edge case: If either string is empty, return an empty result
	if m == 0 || n == 0 {
		return ""
	}

	// Initialize a 2D slice (matrix) for dynamic programming
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	maxLen := 0
	endIndex := 0

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// If characters match, extend the length of the common substring
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				// Update max length and end position if we find a longer substring
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					endIndex = i // Track the end position of the substring in str1
				}
			}
		}
	}

	if maxLen == 0 {
		return ""
	}

	// Return the substring from str1 that ends at endIndex and has length maxLen
	return str1[endIndex-maxLen : endIndex]
}

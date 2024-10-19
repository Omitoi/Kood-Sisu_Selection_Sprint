package sprint

/*
Payout

Learning Objectives:
*	This task involves working with arrays and performing operations like sorting and selecting elements.
*	Designing an algorithm to make payments using denominations involves problem-solving skills and algorithmic thinking.
*	Creating a function for this specific task promotes modular programming, a best practice for code organization and reusability.

Instructions
Write a Go function that takes an integer amount and a slice of integers denominations. 
The function should use the values in denominations to pay exactly the amount. The denominations can be used any number of times, but higher denominations should be preferred. The function should return the resulting denominations as an array of integers in descending order. 
If the payout cannot be made, return an empty array.
*/

func Payout(amount int, denominations []int) (payout []int) {
	for i := 0; i < len(denominations)-1; i++ {
		for j := 0; j < len(denominations)-i-1; j++ {
			if denominations[j] < denominations[j+1] {
				denominations[j], denominations[j+1] = denominations[j+1], denominations[j]
			}
		}
	}

	for _, denom := range denominations {
		for amount >= denom {
			payout = append(payout, denom)
			amount -= denom
		}
	}
	if amount != 0 {
		return []int{}
	}
	return payout
}

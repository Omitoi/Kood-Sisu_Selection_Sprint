package sprint

/*
Combination of N

Learning Objectives:
*	Practice working with loops and combinatorics to generate number combinations.
*	Gain experience in optimizing code for efficiency and performance.

Instructions
Dive into the world of number combinations! Your goal is to create a function that reveals all possible combinations of ascending digits of length n (see example).

Allowed packages
For this task you're allowed to use:
fmt

*/
import "fmt"

func CombN(n int) []string {
	var res []string
	if n > 0 {
		for i := 0; i < 10; i++ {
			if n > 1 {
				for j := i + 1; j < 10; j++ {
					if n > 2 {
						for k := j + 1; k < 10; k++ {
							if n > 3 {
								for l := k + 1; l < 10; l++ {
									if n > 4 {
										for m := l + 1; m < 10; m++ {
											if n > 5 {
												for no := m + 1; no < 10; no++ {
													if n > 6 {
														for o := no + 1; o < 10; o++ {
															if n > 7 {
																for p := o + 1; p < 10; p++ {
																	if n > 8 {
																		for q := p + 1; q < 10; q++ {
																			if n > 9 {
																				for r := q + 1; r < 10; r++ {
																					res = append(res, fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d", i, j, k, l, m, no, o, p, q, r))
																				}
																			} else {
																				res = append(res, fmt.Sprintf("%d%d%d%d%d%d%d%d%d", i, j, k, l, m, no, o, p, q))
																			}
																		}
																	} else {
																		res = append(res, fmt.Sprintf("%d%d%d%d%d%d%d%d", i, j, k, l, m, no, o, p))
																	}
																}
															} else {
																res = append(res, fmt.Sprintf("%d%d%d%d%d%d%d", i, j, k, l, m, no, o))
															}
														}
													} else {
														res = append(res, fmt.Sprintf("%d%d%d%d%d%d", i, j, k, l, m, no))
													}
												}
											} else {
												res = append(res, fmt.Sprintf("%d%d%d%d%d", i, j, k, l, m))
											}
										}
									} else {
										res = append(res, fmt.Sprintf("%d%d%d%d", i, j, k, l))
									}
								}
							} else {
								res = append(res, fmt.Sprintf("%d%d%d", i, j, k))
							}
						}
					} else {
						res = append(res, fmt.Sprintf("%d%d", i, j))
					}
				}
			} else {
				res = append(res, fmt.Sprintf("%d", i))
			}
		}
	} else {
		return []string{}
	}
	return res
}

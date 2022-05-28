package dynamicprog

import "log"

/*
Paint Fence

1. You are given a number n and a number k in separate lines, representing the number of fences and number of colors.
2. You are required to calculate and print the number of ways in which the fences could be painted so that not more than two consecutive  fences have same colors.

Sample Input
8
3

Sample Output
3672
*/

func TestPaintFence() {
	log.Print(paintFence(8, 3))
}

func paintFence(n int, k int) int {
	iiCount := k
	ijCount := k * (k - 1)

	for i := 3; i <= n; i++ {
		ii_new_count := ijCount
		ij_new_count := (iiCount + ijCount) * 2

		iiCount = ii_new_count
		ijCount = ij_new_count
	}

	return iiCount + ijCount
}

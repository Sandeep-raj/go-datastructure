package dynamicprog

import "log"

/*
Tiling With M * 1 Tiles

1. You are given a number n and a number m separated by line-break representing the length and breadth of a m * n floor.
2. You've an infinite supply of m * 1 tiles.
3. You are required to calculate and print the number of ways floor can be tiled using tiles.

Sample Input
39
16

Sample Output
61
*/

func TestTileMx1() {
	log.Print(mx1Tile(39, 16))
}

func mx1Tile(n int, m int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if i < m {
			dp[i] = 1
		} else if i == m {
			dp[i] = 2
		} else {
			dp[i] = dp[i-1] + dp[i-m]
		}
	}

	return dp[n]
}

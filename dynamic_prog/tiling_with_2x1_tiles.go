package dynamicprog

import "log"

/*
Tiling With 2 * 1 Tiles

1. You are given a number n representing the length of a floor space which is 2m wide. It's a 2 * n board.
2. You've an infinite supply of 2 * 1 tiles.
3. You are required to calculate and print the number of ways floor can be tiled using tiles.

Sample Input
8
Sample Output
34
*/

func TestTile2x1() {
	log.Print(tile2x1(8))
}

func tile2x1(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

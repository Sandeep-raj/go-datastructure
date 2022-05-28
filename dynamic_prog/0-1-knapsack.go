package dynamicprog

import "log"

/*
Zero One Knapsack

1. You are given a number n, representing the count of items.
2. You are given n numbers, representing the values of n items.
3. You are given n numbers, representing the weights of n items.
3. You are given a number "cap", which is the capacity of a bag you've.
4. You are required to calculate and print the maximum value that can be created in the bag without
     overflowing it's capacity.

Note -> Each item can be taken 0 or 1 number of times. You are not allowed to put the same item
               again and again.

Sample Input
5
15 14 10 45 30
2 5 1 3 4
7

Sample Output
75
*/

func Test01Knapsack() {
	val := []int{15, 14, 10, 45, 30}
	wt := []int{2, 5, 1, 3, 4}
	W := 7
	log.Print(knapsack01(W, wt, val))
}

func knapsack01(W int, wt []int, val []int) int {
	n := len(wt)

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= W; j++ {
			if wt[i-1] <= j {
				dp[i][j] = maxVal(val[i-1]+dp[i-1][j-wt[i-1]], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][W]
}

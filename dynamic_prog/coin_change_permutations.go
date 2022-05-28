package dynamicprog

import "log"

/*
Coin Change Permutations

1. You are given a number n, representing the count of coins.
2. You are given n numbers, representing the denominations of n coins.
3. You are given a number "amt".
4. You are required to calculate and print the number of permutations of the n coins using which the
     amount "amt" can be paid.

Note1 -> You have an infinite supply of each coin denomination i.e. same coin denomination can be
                  used for many installments in payment of "amt"
Note2 -> You are required to find the count of permutations and not combinations i.e.
                  2 + 2 + 3 = 7 and 2 + 3 + 2 = 7 and 3 + 2 + 2 = 7 are different permutations of same
                  combination. You should treat them as 3 and not 1.

Sample Input
4
2
3
5
6
7

Sample Output
5
*/

func TestCoinChangePermutation() {
	coins := []int{2, 3, 5, 6}
	sum := 7

	log.Print(ccp(coins, sum))
}

func ccp(coins []int, sum int) int {
	dp := make([]int, sum+1)
	dp[0] = 1

	for i := 1; i <= sum; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] += dp[i-coins[j]]
			}
		}
	}

	return dp[sum]
}

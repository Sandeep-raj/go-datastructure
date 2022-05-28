package dynamicprog

import "log"

/*
Coin Change Combination

1. You are given a number n, representing the count of coins.
2. You are given n numbers, representing the denominations of n coins.
3. You are given a number "amt".
4. You are required to calculate and print the number of combinations of the n coins using which the
     amount "amt" can be paid.

Note1 -> You have an infinite supply of each coin denomination i.e. same coin denomination can be
                  used for many installments in payment of "amt"
Note2 -> You are required to find the count of combinations and not permutations i.e.
                  2 + 2 + 3 = 7 and 2 + 3 + 2 = 7 and 3 + 2 + 2 = 7 are different permutations of same
                  combination. You should treat them as 1 and not 3.

Sample Input
4
2
3
5
6
7

Sample Output
2
*/

func TestCoinChangeComb() {
	//n := 4
	inp := []int{2, 3, 5, 6}
	sum := 7

	log.Print(cccTopDown(inp, sum))

}

func cccTopDown(inp []int, sum int) int {
	n := len(inp)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			if j == 0 {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if inp[i-1] <= j {
				// val := maxVal(dp[i][j-inp[i-1]], dp[i-1][j])
				// if val > 0 {
				// 	val += 1
				// }
				dp[i][j] = dp[i][j-inp[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		log.Printf("%+v", dp[i])
	}
	return dp[n][sum]
}

func maxVal(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}

	return val2
}

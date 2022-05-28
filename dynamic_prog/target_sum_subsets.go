package dynamicprog

import (
	"fmt"
	"log"
)

/*
Target Sum Subsets - Dp

1. You are given a number n, representing the count of elements.
2. You are given n numbers.
3. You are given a number "tar".
4. You are required to calculate and print true or false, if there is a subset the elements of which add
     up to "tar" or not.

Sample Input
5
4
2
7
1
3
10

Sample Output
true
*/

func TestTargetSumSubset() {
	inp := []int{4, 2, 7, 1, 3}
	//dp := make(map[string]string)
	sum := 9
	//log.Print(tssRecursive(inp, sum, 0, &dp))
	log.Print(tssTopDown(inp, sum))
}

func tssRecursive(inp []int, sum int, curridx int, dp *map[string]string) bool {
	if curridx == len(inp) {
		return false
	}

	if sum < 0 {
		return false
	}

	if sum == 0 {
		return true
	}

	if (*dp)[fmt.Sprintf("%d_%d", sum, curridx)] != "" {
		return (*dp)[fmt.Sprintf("%d_%d", sum, curridx)] == "true"
	}

	log.Print(curridx, sum)

	exist1 := tssRecursive(inp, sum, curridx+1, dp)
	exist2 := tssRecursive(inp, sum-inp[curridx], curridx+1, dp)

	val := exist1 || exist2

	if val {
		(*dp)[fmt.Sprintf("%d_%d", sum, curridx)] = "true"
	} else {
		(*dp)[fmt.Sprintf("%d_%d", sum, curridx)] = "false"
	}

	return val
}

func tssTopDown(inp []int, sum int) bool {
	n := len(inp)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
		for j := 0; j <= sum; j++ {
			if j == 0 {
				dp[i][j] = true
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if inp[i-1] <= j {
				dp[i][j] = (dp[i-1][j-inp[i-1]] || dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}

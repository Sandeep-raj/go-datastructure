package dynamicprog

import "log"

/*
Partition Into Subsets

1. You are given a number n, representing the number of elements.
2. You are given a number k, representing the number of subsets.
3. You are required to print the number of ways in which these elements can be partitioned in k non-empty subsets.
E.g.
For n = 4 and k = 3 total ways is 6
12-3-4
1-23-4
13-2-4
14-2-3
1-24-3
1-2-34

Sample Input
4
3

Sample Output
6
*/

func TestPartitionSubset() {
	log.Print(partitionSubset(4, 3))
}

func partitionSubset(people int, team int) int {
	dp := make([][]int, team+1)
	for i := 0; i <= team; i++ {
		dp[i] = make([]int, people+1)
	}

	for i := 1; i <= team; i++ {
		for j := 1; j <= people; j++ {
			if i == j {
				dp[i][j] = 1
			} else if j < i {
				dp[i][j] = 0
			} else {
				dp[i][j] = (j-1)*dp[i][j-1] + dp[i-1][j-1]
			}
		}
	}

	return dp[team][people]
}

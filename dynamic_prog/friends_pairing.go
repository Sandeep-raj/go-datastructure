package dynamicprog

import (
	"fmt"
	"log"
)

/*
Friends Pairing

1. You are given a number n, representing the number of friends.
2. Each friend can stay single or pair up with any of it's friends.
3. You are required to print the number of ways in which these friends can stay single or pair up.
E.g.
1 person can stay single or pair up in 1 way.
2 people can stay singles or pair up in 2 ways. 12 => 1-2, 12.
3 people (123) can stay singles or pair up in 4 ways. 123 => 1-2-3, 12-3, 13-2, 23-1.

Sample Input
4

Sample Output
10
*/

func TestFriendsPairing() {
	//used := make([]bool, 3)
	//frdPairRec(0, 3, used, "")
	log.Print(frdPairTopdown(3))
}

func frdPairRec(idx int, n int, used []bool, res string) {
	if idx == n {
		log.Print(res)
		return
	}

	if used[idx] {
		frdPairRec(idx+1, n, used, res)
	} else {
		used[idx] = true
		frdPairRec(idx+1, n, used, fmt.Sprintf("%s (%d) ", res, idx+1))

		for j := idx + 1; j < n; j++ {
			if !used[j] {
				used[j] = true
				frdPairRec(idx+1, n, used, fmt.Sprintf("%s (%d-%d) ", res, idx+1, j+1))
				used[j] = false
			}
		}

		used[idx] = false
	}
}

func frdPairTopdown(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + (n-1)*dp[i-2]
	}

	return dp[n-1]
}

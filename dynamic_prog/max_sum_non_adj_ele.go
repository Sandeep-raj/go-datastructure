package dynamicprog

import "log"

/*
Maximum Sum Non Adjacent Elements

1. You are given a number n, representing the count of elements.
2. You are given n numbers, representing n elements.
3. You are required to find the maximum sum of a subsequence with no adjacent elements.

Sample Input
6
5
10
10
100
5
6

Sample Output
116
*/

func TestMaxSumNonAdjEle() {
	elements := []int{5, 10, 10, 100, 5}
	//log.Print(maxSumNonAdjEle(elements, 0))
	log.Print(maxSumNonAdjEleTopDown(elements))
}

func maxSumNonAdjEle(eles []int, n int) int {
	if n >= len(eles) {
		return 0
	}

	val1 := eles[n] + maxSumNonAdjEle(eles, n+2)
	val2 := maxSumNonAdjEle(eles, n+1)

	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}

func maxSumNonAdjEleTopDown(eles []int) int {
	dp := make([][]int, len(eles))

	for i := 0; i < len(eles); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = eles[0]

	for i := 1; i < len(eles); i++ {
		dp[i][0] = maxVal(dp[i-1][0], dp[i-1][1])
		dp[i][1] = eles[i] + dp[i-1][0]
	}

	return maxVal(dp[len(eles)-1][0], dp[len(eles)-1][1])
}

package stacksandqueues

import "log"

/*
1. Given an integer array nums and a positive integer k;
2. return the lexicographically smallest subsequence of size k;

Sample Input
8
2
4
3
3
5
4
9
6
4

Sample Output
4
2
3
3
4
*/

func TestLexSmallestSubstring() {
	inpArr := []int{2, 4, 3, 3, 5, 4, 9, 6}
	k := 3

	lexSmallestSubstring(inpArr, k)
}

func lexSmallestSubstring(inpArr []int, k int) {
	eleRemoved := len(inpArr) - k
	stack := make([]int, 0)

	for i := 0; i < len(inpArr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > inpArr[i] && eleRemoved > 0 {
			stack = stack[:len(stack)-1]
			eleRemoved--
		}

		stack = append(stack, inpArr[i])
	}

	log.Printf("%+v", stack[:len(stack)-eleRemoved])
}

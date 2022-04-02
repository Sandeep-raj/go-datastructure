package recursion

import "log"

/*
 Count ways to reach the n’th stair
 There are n stairs, a person standing at the bottom wants to reach the top. The person can climb either 1 stair or 2 stairs at a time. Count the number of ways, the person can reach the top.

Input: n = 1
Output: 1
There is only one way to climb 1 stair

Input: n = 2
Output: 2
There are two ways: (1, 1) and (2)

Input: n = 4
Output: 5
(1, 1, 1, 1), (1, 1, 2), (2, 1, 1), (1, 2, 1), (2, 2)
*/

func getStairs(n int, currStr string, res *[]string) {
	if n == 0 {
		*res = append(*res, currStr)
	}

	if n >= 1 {
		getStairs(n-1, currStr+"1", res)
	}

	if n >= 2 {
		getStairs(n-2, currStr+"2", res)
	}
}

func TestGetStairs() {
	res := make([]string, 0)
	getStairs(5, "", &res)
	for _, str := range res {
		log.Print(str)
	}
}

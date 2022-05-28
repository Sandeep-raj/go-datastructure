package dynamicprog

import "log"

/*
Climb Stairs

1. You are given a number n, representing the number of stairs in a staircase.
2. You are on the 0th step and are required to climb to the top.
3. In one move, you are allowed to climb 1, 2 or 3 stairs.
4. You are required to print the number of different paths via which you can climb to the top.

Sample Input
4

Sample Output
7
*/

func TestClimbStairs() {
	dp := make(map[int]int)
	res := climbStairs(10, &dp)
	log.Print(res)
}

func climbStairs(n int, dp *map[int]int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	if (*dp)[n] != 0 {
		return (*dp)[n]
	}

	val := climbStairs(n-1, dp) + climbStairs(n-2, dp) + climbStairs(n-3, dp)

	(*dp)[n] = val

	return val
}

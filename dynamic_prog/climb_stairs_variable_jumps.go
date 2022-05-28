package dynamicprog

import (
	"log"
)

/*
Climb Stairs With Variable Jumps

1. You are given a number n, representing the number of stairs in a staircase.
2. You are on the 0th step and are required to climb to the top.
3. You are given n numbers, where ith element's value represents - till how far from the step you
     could jump to in a single move.
     You can of course jump fewer number of steps in the move.
4. You are required to print the number of different paths via which you can climb to the top.

Sample Input
10
3
3
0
2
1
2
4
2
0
0

Sample Output
5
*/

func TestClimbStairsVarJumps() {
	steps := []int{3, 3, 0, 2, 1, 2, 4, 2, 0, 0}

	dp := make(map[int]int)

	res := climbStairsVarJumps(10, 0, steps, &dp)
	log.Print(res)
}

func climbStairsVarJumps(n int, currStep int, steps []int, dp *map[int]int) int {
	if n-currStep < 0 {
		return 0
	}

	if n-currStep == 0 {
		return 1
	}

	if (*dp)[currStep] != 0 {
		return (*dp)[currStep]
	}

	var totPos int
	for i := 1; i <= steps[currStep]; i++ {
		log.Print(currStep, i)
		totPos += climbStairsVarJumps(n, currStep+i, steps, dp)
	}

	(*dp)[currStep] = totPos

	return totPos
}

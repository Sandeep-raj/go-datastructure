package dynamicprog

import (
	"log"
	"math"
)

/*
Climb Stairs With Minimum Moves

1. You are given a number n, representing the number of stairs in a staircase.
2. You are on the 0th step and are required to climb to the top.
3. You are given n numbers, where ith element's value represents - till how far from the step you
     could jump to in a single move.  You can of-course fewer number of steps in the move.
4. You are required to print the number of minimum moves in which you can reach the top of
     staircase.

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
4
*/

func TestClimbStairsMinJumps() {
	steps := []int{3, 3, 0, 2, 1, 2, 4, 2, 0, 0}

	dp := make(map[int]int)

	minJumps := climbStairsMinJumps(10, 0, steps, &dp)
	log.Print(*minJumps)
}

func climbStairsMinJumps(n int, currStep int, steps []int, dp *map[int]int) *int {
	if n-currStep < 0 {
		return nil
	}

	if n-currStep == 0 {
		res := 0
		return &res
	}

	if (*dp)[currStep] != 0 {
		res := (*dp)[currStep]
		return &res
	}

	minJumps := math.MaxInt
	var currRes *int

	log.Print(currStep)

	for i := steps[currStep]; i > 0; i-- {
		currRes = climbStairsMinJumps(n, currStep+i, steps, dp)

		if currRes != nil && *currRes < minJumps {
			minJumps = *currRes
		}
	}

	if minJumps != math.MaxInt {
		minJumps += 1
		(*dp)[currStep] = minJumps
		return &minJumps
	} else {
		return nil
	}

}

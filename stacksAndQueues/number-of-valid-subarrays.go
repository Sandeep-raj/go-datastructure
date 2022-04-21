package stacksandqueues

import "log"

/*
Number Of Valid Subarrays

Given an array nums of integers, return the number of non-empty continuous subarrays that satisfy the following condition:
The leftmost element of the subarray is not larger than other elements in the subarray.

Sample Input
5
1
4
2
5
3

Sample Output
11
*/

func TestNumValidSubarrays() {
	inpArr := []int{1, 4, 2, 5, 3}
	numValidSubArr(inpArr)
}

func numValidSubArr(inpArr []int) {
	totCount := 0
	stack := make([]int, 0)

	for i := len(inpArr) - 1; i >= 0; i-- {
		for len(stack) > 0 && inpArr[stack[len(stack)-1]] > inpArr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			totCount = totCount + (len(inpArr) - i)
		} else {
			totCount = totCount + (stack[len(stack)-1] - i)
		}

		stack = append(stack, i)
	}

	log.Print(totCount)
}

package stacksandqueues

import "log"

/*
132 Pattern

Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].
Return true if there is a 132 pattern in nums, otherwise, return false.

Sample Input
4
3
1
4
2

Sample Output
true
*/

func Test132Pattern() {
	//n := 4
	inpArr := []int{3, 1, 4, 2}

	log.Print(pattern132(inpArr))

}

func pattern132(inpArr []int) bool {
	// lowest Arr
	minArr := make([]int, len(inpArr))
	min := 1000

	for i := 0; i < len(inpArr); i++ {
		if inpArr[i] < min {
			min = inpArr[i]
		}
		minArr[i] = min
	}

	stack := make([]int, 0)

	for j := len(inpArr) - 1; j >= 1; j-- {
		for len(stack) > 0 && stack[len(stack)-1] < minArr[j] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && stack[len(stack)-1] < inpArr[j] {
			return true
		} else {
			stack = append(stack, inpArr[j])
		}
	}

	return false
}

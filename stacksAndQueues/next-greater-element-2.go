package stacksandqueues

import "log"

/*
Next Greater Element Ii

Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]),
return the next greater number for every element in nums.

Sample Input
3 8 4 1 2

Sample Output
8 -1 8 2 3
*/

func TestNGE2() {
	inpArr := []int{3, 8, 4, 1, 2}
	nge2(inpArr)
}

func nge2(inpArr []int) {
	rstack := make([]int, 0)
	res := make([]int, len(inpArr))

	for i := len(inpArr) - 1; i >= 0; i-- {
		if inpArr[i] > inpArr[len(inpArr)-1] {
			rstack = append(rstack, inpArr[i])
		}
	}

	for i := len(inpArr) - 1; i >= 0; i-- {
		for len(rstack) > 0 && inpArr[i] >= rstack[len(rstack)-1] {
			rstack = rstack[:len(rstack)-1]
		}

		if len(rstack) == 0 {
			res[i] = -1
		} else {
			res[i] = rstack[len(rstack)-1]
		}

		rstack = append(rstack, inpArr[i])
	}

	log.Printf("%+v", res)
}

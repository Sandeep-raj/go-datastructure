package stacksandqueues

import "log"

/*
Maximal Rectangle

1. Given a rows x cols binary matrix filled with 0's and 1's.
2. Find the largest rectangle containing only 1's and return its area.

Sample Input
4
4
0110
1111
1111
1100

Sample Output
8
*/

func TestMaxRectangle() {
	inpArr := make([][]int, 4)
	inpArr[0] = []int{0, 1, 1, 0}
	inpArr[1] = []int{1, 1, 1, 1}
	inpArr[2] = []int{1, 1, 1, 1}
	inpArr[3] = []int{1, 1, 0, 0}

	currHistogram := make([]int, len(inpArr[0]))
	maxArea := 0

	for i := 0; i < len(inpArr); i++ {
		for j := 0; j < len(inpArr[i]); j++ {
			currHistogram[j] = inpArr[i][j] * (currHistogram[j] + inpArr[i][j])
		}
		currArea := MAH1(currHistogram)
		if currArea > maxArea {
			maxArea = currArea
		}
	}

	log.Print(maxArea)
}

func MAH1(inpArr []int) int {
	lstack := make([]int, 0)
	rstack := make([]int, 0)

	ls := make([]int, len(inpArr))
	rs := make([]int, len(inpArr))

	for i := 0; i < len(inpArr); i++ {
		for len(lstack) > 0 && inpArr[lstack[len(lstack)-1]] >= inpArr[i] {
			lstack = lstack[:len(lstack)-1]
		}

		for len(rstack) > 0 && inpArr[rstack[len(rstack)-1]] >= inpArr[len(inpArr)-1-i] {
			rstack = rstack[:len(rstack)-1]
		}

		if len(lstack) == 0 {
			ls[i] = -1
		} else {
			ls[i] = lstack[len(lstack)-1]
		}

		if len(rstack) == 0 {
			rs[len(inpArr)-1-i] = len(inpArr)
		} else {
			rs[len(inpArr)-1-i] = rstack[len(rstack)-1]
		}

		lstack = append(lstack, i)
		rstack = append(rstack, len(inpArr)-1-i)
	}

	maxArea := 0

	for i := 0; i < len(inpArr); i++ {
		currArea := (inpArr[i]) * (rs[i] - ls[i] - 1)

		if currArea > maxArea {
			maxArea = currArea
		}
	}

	return maxArea
}

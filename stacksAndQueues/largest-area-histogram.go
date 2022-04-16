package stacksandqueues

import "log"

/*
Largest Area Histogram

1. You are given a number n, representing the size of array a.
2. You are given n numbers, representing the height of bars in a bar chart.
3. You are required to find and print the area of largest rectangle in the histogram.

e.g. for the array [6 2 5 4 5 1 6] -> 12

Sample Input
7
6
2
5
4
5
1
6

Sample Output
12
*/

func TestLAH() {
	inpArr := []int{6, 2, 5, 4, 5, 1, 6}

	lah(inpArr)
}

func lah(inpArr []int) {
	indexArr := make([]int, 0)
	maxArea := 0

	for i := 0; i < len(inpArr); i++ {
		for len(indexArr) > 0 && inpArr[i] < inpArr[indexArr[len(indexArr)-1]] {
			currIndex := indexArr[len(indexArr)-1]
			indexArr = indexArr[:len(indexArr)-1]

			currMax := 0

			if len(indexArr) == 0 {
				currMax = i * inpArr[currIndex]
			} else {
				currMax = (i - indexArr[len(indexArr)-1] - 1) * inpArr[currIndex]
			}

			if currMax > maxArea {
				maxArea = currMax
			}
		}

		indexArr = append(indexArr, i)
	}

	for len(indexArr) > 0 {
		currIndex := indexArr[len(indexArr)-1]
		indexArr = indexArr[:len(indexArr)-1]

		currMax := 0

		if len(indexArr) == 0 {
			currMax = len(inpArr) * inpArr[currIndex]
		} else {
			currMax = (len(inpArr) - indexArr[len(indexArr)-1] - 1) * inpArr[currIndex]
		}

		if currMax > maxArea {
			maxArea = currMax
		}
	}

	log.Print(maxArea)
}

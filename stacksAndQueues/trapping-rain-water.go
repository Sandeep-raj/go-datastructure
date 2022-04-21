package stacksandqueues

import (
	"log"
	"math"
)

/*
Trapping Rain Water

1. Given n non-negative integers representing an elevation map where the width of each bar is 1.
2. Compute how much water it can trap after raining.

Sample Input
12
0
1
0
2
1
0
1
3
2
1
2
1

Sample Output
6
*/

func TestRainWaterTrap() {
	inpArr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	rainWaterTrap(inpArr)
}

func rainWaterTrap(inpArr []int) {

	leftMaxima := make([]int, 0)
	rightMaxima := make([]int, 0)

	ngel := make([]int, len(inpArr))
	nger := make([]int, len(inpArr))

	for i := 0; i < len(inpArr); i++ {
		for len(leftMaxima) > 0 && leftMaxima[len(leftMaxima)-1] < inpArr[i] {
			leftMaxima = leftMaxima[:len(leftMaxima)-1]
		}

		for len(rightMaxima) > 0 && rightMaxima[len(rightMaxima)-1] < inpArr[len(inpArr)-1-i] {
			rightMaxima = rightMaxima[:len(rightMaxima)-1]
		}

		if len(leftMaxima) == 0 {
			ngel[i] = -1
		} else {
			ngel[i] = leftMaxima[len(leftMaxima)-1]
		}

		if len(rightMaxima) == 0 {
			nger[len(inpArr)-1-i] = -1
		} else {
			nger[len(inpArr)-1-i] = rightMaxima[len(rightMaxima)-1]
		}

		if len(leftMaxima) == 0 || (len(leftMaxima) > 0 && inpArr[i] > leftMaxima[len(leftMaxima)-1]) {
			leftMaxima = append(leftMaxima, inpArr[i])
		}

		if len(rightMaxima) == 0 || (len(rightMaxima) > 0 && inpArr[len(inpArr)-1-i] > rightMaxima[len(rightMaxima)-1]) {
			rightMaxima = append(rightMaxima, inpArr[len(inpArr)-1-i])
		}
	}

	trappedWater := 0

	for i := 0; i < len(inpArr); i++ {
		if ngel[i] == -1 || nger[i] == -1 {
			continue
		} else {
			trappedWater = trappedWater + int(math.Min(float64(ngel[i]), float64(nger[i]))) - inpArr[i]
		}
	}

	log.Print(trappedWater)
}

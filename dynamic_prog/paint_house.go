package dynamicprog

import (
	"log"
	"math"
)

/*
Paint House - Many Colors

1. You are given a number n and a number k separated by a space, representing the number of houses and number of colors.
2. In the next n rows, you are given k space separated numbers representing the cost of painting nth house with one of the k colors.
3. You are required to calculate and print the minimum cost of painting all houses without painting any consecutive house with same color.

Sample Input
4 3
1 5 7
5 8 4
3 2 9
1 2 4

Sample Output
8
*/

func TestPaintHouse() {
	houseColorCost := make([][]int, 4)
	houseColorCost[0] = []int{1, 5, 7}
	houseColorCost[1] = []int{5, 8, 4}
	houseColorCost[2] = []int{3, 2, 9}
	houseColorCost[3] = []int{1, 2, 4}

	// log.Print(paintHouse(houseColorCost, 0, -1))

	log.Print(paintHouseTopDown(houseColorCost))
}

func paintHouseTopDown(houseColorCost [][]int) int {
	dp := make([][]int, len(houseColorCost))
	for i := 0; i < len(houseColorCost); i++ {
		dp[i] = make([]int, len(houseColorCost[0]))
		if i == 0 {
			for j := 0; j < len(houseColorCost[0]); j++ {
				dp[i][j] = houseColorCost[i][j]
			}
		}
	}

	for i := 1; i < len(houseColorCost); i++ {
		for j := 0; j < len(houseColorCost[0]); j++ {
			currCost := DeepCopy(dp[i-1])
			currCost[j] = math.MaxInt
			dp[i][j] = houseColorCost[i][j] + minVal(currCost)
		}
	}

	return minVal(dp[len(houseColorCost)-1])
}

func paintHouse(houseColorCost [][]int, n int, color int) int {
	if n == len(houseColorCost) {
		return 0
	}

	currVals := make([]int, len(houseColorCost[0]))

	for i := 0; i < len(houseColorCost[0]); i++ {
		if color == i {
			currVals[i] = math.MaxInt
		} else {
			currVals[i] = houseColorCost[n][i] + paintHouse(houseColorCost, n+1, i)
		}
	}

	return minVal(currVals)
}

func minVal(valArr []int) int {
	minRes := math.MaxInt

	for _, v := range valArr {
		if v < minRes {
			minRes = v
		}
	}

	return minRes
}

func DeepCopy(src []int) []int {
	res := make([]int, len(src))

	for i := 0; i < len(src); i++ {
		res[i] = src[i]
	}

	return res
}

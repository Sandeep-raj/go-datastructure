package dynamicprog

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Fractional Knapsack

1. You are given a number n, representing the count of items.
2. You are given n numbers, representing the values of n items.
3. You are given n numbers, representing the weights of n items.
3. You are given a number "cap", which is the capacity of a bag you've.
4. You are required to calculate and print the maximum value that can be created in the bag without overflowing it's capacity.
Note1 -> Items can be added to the bag even partially. But you are not allowed to put same items again and again to the bag.

Sample Input
10
33 14 50 9 8 11 6 40 2 15
7 2 5 9 3 2 1 10 3 3
5

Sample Output
50.0
*/

func TestFractionalKnapsack() {
	val := []int{33, 14, 50, 9, 8, 11, 6, 40, 2, 15}
	wt := []int{7, 2, 5, 9, 3, 2, 1, 10, 3, 3}

	W := 5

	log.Print(fractionalKnapsack(wt, val, W))
}

func fractionalKnapsack(wt []int, val []int, W int) float64 {
	qs := utils.QS{
		List:  []utils.QSObject{},
		Count: 0,
	}

	for i := 0; i < len(wt); i++ {
		qs.List = append(qs.List, utils.QSObject{
			Key:     float32(val[i]) / float32(wt[i]),
			Payload: []int{wt[i], val[i]},
		})
		qs.Count++
	}

	qs.QSSort()

	currIndex := qs.Count - 1
	maxVal := 0.0
	for W > 0 {
		vwratio := qs.List[currIndex]
		if vwratio.Payload.([]int)[0] < W {
			W = W - vwratio.Payload.([]int)[0]
			maxVal = maxVal + float64(vwratio.Payload.([]int)[1])
		} else {
			maxVal = maxVal + float64((float32(W) * vwratio.Key))
			W = 0
		}
		currIndex--
	}

	return maxVal
}

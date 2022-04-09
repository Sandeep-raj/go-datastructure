package recursion

import (
	"fmt"
	"log"
	"math"
)

/*
Largest Number Possible After At Most K Swaps

1. You are given a string which represents digits of a number.
2. You have to create the maximum number by performing at-most k swap operations on its digits.

Sample Input
1234567
4

Sample Output
7654321
*/

var maxVal int64

func TestLargestNumber() {
	inp := 1234567
	inpArr := make([]int, len(fmt.Sprintf("%d", inp)))
	for i := len(fmt.Sprintf("%d", inp)) - 1; i >= 0; i-- {
		inpArr[i] = inp % 10
		inp = inp / 10
	}

	largestNumber(0, 4, inpArr)
	log.Print(maxVal)

}

func largestNumber(idx int, k int, numArr []int) {

	currVal := getInt(numArr)
	if maxVal < currVal {
		maxVal = currVal
	}

	if k == 0 {
		return
	}

	for i := idx + 1; i < len(numArr); i++ {
		if numArr[idx] < numArr[i] {
			swapArr(numArr, idx, i)
			largestNumber(idx+1, k-1, numArr)
			swapArr(numArr, i, idx)
		}
	}
}

func swapArr(numArr []int, i int, j int) {
	temp := numArr[i]
	numArr[i] = numArr[j]
	numArr[j] = temp
}

func getInt(numArr []int) int64 {
	var res int64
	for i := 0; i < len(numArr); i++ {
		res = res + int64(int(math.Pow10(i))*numArr[len(numArr)-1-i])
	}
	return res
}

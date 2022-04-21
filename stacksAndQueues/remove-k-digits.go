package stacksandqueues

import (
	"fmt"
	"strconv"
)

/*
Remove K Digits

Given string num representing a non-negative integer num, and an integer k, print the smallest possible integer after removing k digits from num.

Sample Input
1432219
3

Sample Output
1219
*/

func TestRemKDigits() {
	inpStr := "1432219"
	k := 3

	remkdigits(inpStr, k)
}

func remkdigits(inpArr string, k int) {
	stack := make([]int, 0)

	for i := 0; i < len(inpArr); i++ {
		currVal, _ := strconv.Atoi(string(inpArr[i]))

		for len(stack) > 0 && stack[len(stack)-1] > currVal && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, currVal)
	}

	for i := 0; i < len(stack); i++ {
		fmt.Print(stack[i])
	}

}

package stacksandqueues

import (
	"fmt"
	"log"
)

/*
Smallest Number Following Pattern

1. You are given a pattern of upto 8 length containing characters 'i' and 'd'.
2. 'd' stands for decreasing and 'i' stands for increasing
3. You have to print the smallest number, using the digits 1 to 9 only without repetition, such that the digit decreases following a d and increases follwing an i.
e.g. d -> 21 i -> 12 ddd -> 4321 iii -> 1234 dddiddd -> 43218765 iiddd -> 126543

Sample Input
ddddiiii

Sample Output
543216789
*/

func TestSNFP() {
	inpStr := "iiddd"
	snfp(inpStr)
}

func snfp(inpStr string) {
	stack := make([]int, 0)
	currVal := 1
	res := ""

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if currChar == "d" {
			stack = append(stack, currVal)
			currVal++
		} else {
			res = fmt.Sprintf("%s%d", res, currVal)
			for len(stack) > 0 {
				res = fmt.Sprintf("%s%d", res, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			currVal++
		}
	}

	res = fmt.Sprintf("%s%d", res, currVal)
	if len(stack) > 0 {
		for len(stack) > 0 {
			res = fmt.Sprintf("%s%d", res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	log.Print(res)
}

package stacksandqueues

import "log"

/*
Validate Stack Sequences

1: Given two sequences pushed and popped with distinct values.
2: You have to return true if and only if this could have been the result of a sequence of push and pop operations on an initially empty stack

Sample Input

1 2 3 4 5
4 5 3 2 1

Sample Output
true
*/

func TestValidateStackSeq() {
	pushedArr := []int{1, 2, 3, 4, 5}
	popedArr := []int{4, 5, 3, 2, 1}

	log.Print(validateStackSeq(pushedArr, popedArr))
}

func validateStackSeq(pushedArr []int, popedArr []int) bool {
	stack := make([]int, 0)

	currIndex := 0

	for i := 0; i < len(pushedArr); i++ {
		stack = append(stack, pushedArr[i])

		for len(stack) > 0 && currIndex < len(popedArr) && stack[len(stack)-1] == popedArr[currIndex] {
			stack = stack[:len(stack)-1]
			currIndex++
		}
	}

	return currIndex == len(popedArr)
}

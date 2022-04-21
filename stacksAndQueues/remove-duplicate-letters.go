package stacksandqueues

import "fmt"

/*
Remove Duplicate Letters

1. Given a string s, remove duplicate letters so that every letter appears once and only once.
2. You must make sure your result is the first in dictionary order among all possible results.

Sample Input
bcabc

Sample Output
abc
*/

func TestRemDupLetter() {
	inpStr := "bcabc"
	remDupLetters(inpStr)
}

func remDupLetters(inpStr string) {
	stack := make([]byte, 0)
	tempStack := make([]byte, 0)

	for i := 0; i < len(inpStr); i++ {

		for len(stack) > 0 && stack[len(stack)-1] > inpStr[i] {
			tempStack = append(tempStack, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 || len(stack) > 0 && stack[len(stack)-1] != inpStr[i] {
			stack = append(stack, inpStr[i])
		}

		for len(tempStack) > 0 {
			stack = append(stack, tempStack[len(tempStack)-1])
			tempStack = tempStack[:len(tempStack)-1]
		}
	}

	for i := 0; i < len(stack); i++ {
		fmt.Print(string(stack[i]))
	}
}

package stacksandqueues

import (
	"log"
	"strings"
)

/*
Reverse Substrings Between Each Pair Of Parentheses

1: You are given a string s that consists of lower case English letters and brackets.
2: Reverse the strings in each pair of matching parentheses, starting from the innermost one.
3: Your result should not contain any brackets.

Example
(abcd) -> dcba
(u(love)i) -> iloveu
(gni(pc(do))ep) -> pepcoding
*/

func TestRevStringBwParenthesis() {
	inpStr := "(gni(pc(do))ep)"
	revStringParenthesis(inpStr)

}

func revStringParenthesis(inpStr string) {
	stack := make([]string, 0)

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if currChar == ")" {
			tempQ := make([]string, 0)
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				tempQ = append(tempQ, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for len(tempQ) > 0 {
				stack = append(stack, tempQ[0])
				tempQ = tempQ[1:]
			}
		} else {
			stack = append(stack, currChar)
		}
	}

	log.Print(strings.Join(stack, ""))
}

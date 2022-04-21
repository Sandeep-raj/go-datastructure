package stacksandqueues

import (
	"fmt"
	"log"
	"strconv"
)

/*
Score Of Parentheses

Given a balanced parentheses string S, compute the score of the string based on the following rule:
    () has score 1
    AB has score A + B, where A and B are balanced parentheses strings.
    (A) has score 2 * A, where A is a balanced parentheses string.

Score of ()()() string is 3 => 1 + 1 + 1
Score of (()) string is 2 => 2 * 1

Sample Input
(()(()))

Sample Output
6
*/

func TestScoreParenthesis() {
	inpStr := "(()(()))"

	scoreParenthesis(inpStr)
}

func scoreParenthesis(inpStr string) {
	stack := make([]string, 0)

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if currChar == ")" {
			tempSum := 0
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				val, _ := strconv.Atoi(stack[len(stack)-1])
				tempSum = tempSum + val
				stack = stack[:len(stack)-1]
			}

			stack = stack[:len(stack)-1]
			if tempSum == 0 {
				stack = append(stack, fmt.Sprintf("%d", 1))
			} else {
				stack = append(stack, fmt.Sprintf("%d", (2*tempSum)))
			}
		} else {
			stack = append(stack, currChar)
		}
	}

	log.Print(stack[0])
}

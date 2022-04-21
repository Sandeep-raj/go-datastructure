package stacksandqueues

import "log"

/*
Minimum Add To Make Parentheses Valid

1: Given a string S of '(' and ')' parentheses.
2: You need to find count of minimum number of parentheses '(' or ')' when added in any positions so that the resulting parentheses string is valid.

Sample Input
()))((

Sample Output
4
*/

func TestMinParanthesis() {
	inpStr := "()))(("
	minParanthesis(inpStr)
}

func minParanthesis(inpStr string) {
	stack := make([]string, 0)

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if currChar == ")" && len(stack) > 0 && stack[len(stack)-1] == "(" {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, currChar)
		}
	}

	log.Print(len(stack))
}

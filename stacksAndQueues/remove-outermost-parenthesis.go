package stacksandqueues

import (
	"fmt"
	"log"
	"strings"
)

/*
Remove Outermost Parentheses

1. You are given a valid parentheses string in form of A+B+C... where A, B and C are valid primitive strings.
2. A primitive string is a valid parentheses string which cant is split in s = x+y, with x and y nonempty valid parentheses strings.
3. You have to remove the outermost parentheses from all primitive strings.

Example "(()())(())" = "(()())" + "(())".
removing outermost parentheses from "(()())" and "(())" will result in ()()().

Sample Input
(()())(())

Sample Output
()()()
*/

func TestRemOutermostParenthesis() {
	inpStr := "(()())(())()"
	remOutermostParenthesis(inpStr)

}

func remOutermostParenthesis(inpStr string) {
	stack := make([]string, 0)

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if currChar == ")" {
			tempStr := ""
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				tempStr = fmt.Sprintf("%s%s", tempStr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			if len(tempStr) > 0 {
				stack = stack[:len(stack)-1]
				stack = append(stack, tempStr)
			} else {
				stack = stack[:len(stack)-1]
				stack = append(stack, "()")
			}
		} else {
			stack = append(stack, currChar)
		}
	}

	log.Print(strings.Join(stack, ""))
}

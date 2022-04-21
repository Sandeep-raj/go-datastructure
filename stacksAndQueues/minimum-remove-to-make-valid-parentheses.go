package stacksandqueues

import (
	"fmt"
	"log"
)

/*
Minimum Remove To Make Valid Parentheses

1: Given a string s of '(' , ')' and lowercase English characters
2: Your task is to remove the minimum number of parentheses ( '(' or ')') so that the resulting parentheses string is valid and return it.
3: In case of multiple valid strings give precedence in keeping innermost parenthesis.

example
(a(b(c)d) this string has (a(bc)d), (ab(c)d) and a(b(c)d) 3 valid strings.
Among all 3 valid strings a(b(c)d) has the innermost parentheses.

Sample Input
a)b(c)d

Sample Output
ab(c)d
*/

func TestMinRemParenthesis() {
	inpStr := "a)b(c)d"
	minRemParenthesis(inpStr)
}

func minRemParenthesis(inpStr string) {
	stack := make([]string, 0)

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if currChar == ")" {
			tempStr := ""

			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				tempStr = fmt.Sprintf("%s%s", stack[len(stack)-1], tempStr)
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				stack = append(stack, tempStr)
			} else {
				stack = append(stack, fmt.Sprintf("(%s)", tempStr))
			}
		} else {
			stack = append(stack, currChar)
		}
	}

	res := ""
	for i := 0; i < len(stack); i++ {
		if stack[i] == "(" {
			continue
		} else {
			res = fmt.Sprintf("%s%s", res, stack[i])
		}
	}

	log.Print(res)
}

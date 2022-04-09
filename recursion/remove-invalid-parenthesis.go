package recursion

import (
	"fmt"
	"log"
)

/*
Remove Invalid Parenthesis

1. You are given a string, which represents an expression having only opening and closing parenthesis.
2. You have to remove minimum number of parenthesis to make the given expression valid.
3. If there are multiple answers, you have to print all of them.

Sample Input
()())()

Sample Output
(())()
()()()
*/

func TestInvalidParenthesis() {
	inpStr := "()())()"
	finalRes := make([]string, 0)
	removeInvalidParanthesis(0, inpStr, 0, "", &finalRes)
	log.Printf("%+v", finalRes)
}

func removeInvalidParanthesis(idx int, inpStr string, blocks int, resStr string, finalRes *[]string) {
	if idx == len(inpStr) {
		if blocks == 0 {
			if len(*finalRes) > 0 {
				if len((*finalRes)[0]) < len(resStr) {
					*finalRes = append((*finalRes)[0:0], resStr)
				} else if len((*finalRes)[0]) == len(resStr) {
					*finalRes = append((*finalRes), resStr)
				}
			} else {
				*finalRes = append((*finalRes), resStr)
			}
		}
		return
	}

	curChar := string(inpStr[idx])
	if curChar == "(" {
		// either you add
		removeInvalidParanthesis(idx+1, inpStr, blocks+1, fmt.Sprintf("%s%s", resStr, curChar), finalRes)
		// go ahead without adding
		removeInvalidParanthesis(idx+1, inpStr, blocks, resStr, finalRes)
	} else {
		if blocks-1 >= 0 {
			removeInvalidParanthesis(idx+1, inpStr, blocks-1, fmt.Sprintf("%s%s", resStr, curChar), finalRes)
		}
		removeInvalidParanthesis(idx+1, inpStr, blocks, resStr, finalRes)
	}
}

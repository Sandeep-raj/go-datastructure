package stacksandqueues

import "log"

/*
Balanced Brackets

1. You are given a string exp representing an expression.
2. You are required to check if the expression is balanced i.e. closing brackets and opening brackets match up well.

e.g.
[(a + b) + {(c + d) * (e / f)}] -> true
[(a + b) + {(c + d) * (e / f)]} -> false
[(a + b) + {(c + d) * (e / f)} -> false
([(a + b) + {(c + d) * (e / f)}] -> false

Sample Input
[(a + b) + {(c + d) * (e / f)}]

Sample Output
true
*/

func TestBalancedBrackets() {
	log.Print(balancedBrackets("[(a + b) + {(c + d) * (e / f)]}"))
}

func balancedBrackets(inpStr string) bool {
	stack := make([]string, 0)

	openBracketMap := make(map[string]string)
	openBracketMap["("] = ")"
	openBracketMap["{"] = "}"
	openBracketMap["["] = "]"

	bracketMap := make(map[string]string)
	bracketMap["]"] = "["
	bracketMap[")"] = "("
	bracketMap["}"] = "{"
	for i := 0; i < len(inpStr); i++ {
		if string(inpStr[i]) == " " {
			continue
		} else if bracketMap[string(inpStr[i])] == "" {
			stack = append(stack, string(inpStr[i]))
		} else {
			limitBracket := bracketMap[string(inpStr[i])]
			for len(stack) > 0 && string(stack[len(stack)-1]) != limitBracket {
				if openBracketMap[stack[len(stack)-1]] != "" {
					return false
				}
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

package stacksandqueues

import (
	"fmt"
	"log"
)

/*
Infix Conversions

Sample Input
a*(b-c+d)/e

Sample Output
abc-d+*e/
/*a+-bcde
*/

func TestInfixConv() {
	inpStr := "a*(b-c+d)/e"
	infix2Postfix(inpStr)
	infix2Prefix(inpStr)
}

func infix2Postfix(inpStr string) {
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3

	opStack := make([]string, 0)
	res := ""

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if currChar == " " {
			continue
		} else if opPrecedence[currChar] == 0 {
			res = fmt.Sprintf("%s%s", res, currChar)
		} else {
			if currChar == ")" {
				for opStack[len(opStack)-1] != "(" {
					res = fmt.Sprintf("%s%s", res, opStack[len(opStack)-1])
					opStack = opStack[:len(opStack)-1]
				}
				opStack = opStack[:len(opStack)-1]
			} else {
				for len(opStack) > 0 &&
					opPrecedence[opStack[len(opStack)-1]] >= opPrecedence[currChar] && opStack[len(opStack)-1] != "(" {
					res = fmt.Sprintf("%s%s", res, opStack[len(opStack)-1])
					opStack = opStack[:len(opStack)-1]
				}
				opStack = append(opStack, currChar)
			}
		}
	}

	for len(opStack) > 0 {
		res = fmt.Sprintf("%s%s", res, opStack[len(opStack)-1])
		opStack = opStack[:len(opStack)-1]
	}

	log.Print(res)
}

func infix2Prefix(inpStr string) {
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3
	opStack := make([]string, 0)
	res := ""

	for i := len(inpStr) - 1; i >= 0; i-- {
		currChar := string(inpStr[i])

		if currChar == " " {
			continue
		} else if opPrecedence[currChar] == 0 {
			res = fmt.Sprintf("%s%s", currChar, res)
		} else {
			if currChar == "(" {
				for len(opStack) > 0 && opStack[len(opStack)-1] != ")" {
					res = fmt.Sprintf("%s%s", opStack[len(opStack)-1], res)
					opStack = opStack[:len(opStack)-1]
				}
				opStack = opStack[:len(opStack)-1]
			} else {
				for len(opStack) > 0 &&
					opPrecedence[opStack[len(opStack)-1]] > opPrecedence[currChar] &&
					opStack[len(opStack)-1] != ")" {
					res = fmt.Sprintf("%s%s", opStack[len(opStack)-1], res)
					opStack = opStack[:len(opStack)-1]
				}

				opStack = append(opStack, currChar)
			}
		}
	}

	for len(opStack) > 0 {
		res = fmt.Sprintf("%s%s", opStack[len(opStack)-1], res)
		opStack = opStack[:len(opStack)-1]
	}

	log.Print(res)
}

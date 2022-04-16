package stacksandqueues

import (
	"fmt"
	"log"
	"strconv"
)

/*
Constraints
1. Expression is a valid postfix expression
2. The only operators used are +, -, *, /
3. All operands are single digit numbers.

Sample Input
264*8/+3-

Sample Output
2
((2+((6*4)/8))-3)
-+2/*6483
*/

func TestPostfixConv() {
	inpStr := "264*8/+3-"

	postfixEval(inpStr)
	postfix2Infix(inpStr)
	postfix2Prefix(inpStr)
}

func postfixEval(inpStr string) {
	operandStack := make([]int, 0)
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if opPrecedence[currChar] == 0 {
			val, _ := strconv.Atoi(currChar)
			operandStack = append(operandStack, val)
		} else {
			val1 := operandStack[len(operandStack)-2]
			val2 := operandStack[len(operandStack)-1]

			operandStack = operandStack[:len(operandStack)-2]

			operandStack = append(operandStack, CalculateExp(currChar, val1, val2))
		}
	}

	log.Print(operandStack[0])
}

func postfix2Infix(inpStr string) {
	operandStack := make([]string, 0)
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if opPrecedence[currChar] == 0 {
			operandStack = append(operandStack, currChar)
		} else {
			val1 := operandStack[len(operandStack)-2]
			val2 := operandStack[len(operandStack)-1]

			operandStack = operandStack[:len(operandStack)-2]

			operandStack = append(operandStack, fmt.Sprintf("(%s%s%s)", val1, currChar, val2))
		}
	}

	log.Print(operandStack[0])
}

func postfix2Prefix(inpStr string) {
	operandStack := make([]string, 0)
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3

	for i := 0; i < len(inpStr); i++ {
		currChar := string(inpStr[i])

		if opPrecedence[currChar] == 0 {
			operandStack = append(operandStack, currChar)
		} else {
			val1 := operandStack[len(operandStack)-2]
			val2 := operandStack[len(operandStack)-1]

			operandStack = operandStack[:len(operandStack)-2]

			operandStack = append(operandStack, fmt.Sprintf("%s%s%s", currChar, val1, val2))
		}
	}

	log.Print(operandStack[0])
}

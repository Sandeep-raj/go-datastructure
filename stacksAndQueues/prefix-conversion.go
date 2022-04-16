package stacksandqueues

import (
	"fmt"
	"log"
	"strconv"
)

/*
Prefix Evaluation And Conversions

1. You are given a prefix expression.
2. You are required to evaluate it and print it's value.
3. You are required to convert it to infix and print it.
4. You are required to convert it to postfix and print it.

Constraints
1. Expression is a valid prefix expression
2. The only operators used are +, -, *, /
3. All operands are single digit numbers.

Sample Input
-+2/*6483

Sample Output
2
((2+((6*4)/8))-3)
264*8/+3-
*/

func TestPrefixConv() {
	inpStr := "-+2/*6483"
	PrefixEval(inpStr)
	Prefix2Infix(inpStr)
	Prefix2Postfix(inpStr)
}

func PrefixEval(inpStr string) {
	operandStack := make([]int, 0)
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3

	for i := len(inpStr) - 1; i >= 0; i-- {
		currChar := string(inpStr[i])

		if opPrecedence[currChar] == 0 {
			val, _ := strconv.Atoi(currChar)
			operandStack = append(operandStack, val)
		} else {
			val1 := operandStack[len(operandStack)-1]
			val2 := operandStack[len(operandStack)-2]

			operandStack = operandStack[:len(operandStack)-2]

			operandStack = append(operandStack, CalculateExp(currChar, val1, val2))
		}
	}

	log.Print(operandStack[0])
}

func Prefix2Infix(inpStr string) {
	operandStack := make([]string, 0)
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3

	for i := len(inpStr) - 1; i >= 0; i-- {
		currChar := string(inpStr[i])

		if opPrecedence[currChar] == 0 {
			operandStack = append(operandStack, currChar)
		} else {
			val1 := operandStack[len(operandStack)-1]
			val2 := operandStack[len(operandStack)-2]

			operandStack = operandStack[:len(operandStack)-2]

			operandStack = append(operandStack, fmt.Sprintf("(%s%s%s)", val1, currChar, val2))
		}
	}

	log.Print(operandStack[0])
}

func Prefix2Postfix(inpStr string) {
	operandStack := make([]string, 0)
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3

	for i := len(inpStr) - 1; i >= 0; i-- {
		currChar := string(inpStr[i])

		if opPrecedence[currChar] == 0 {
			operandStack = append(operandStack, currChar)
		} else {
			val1 := operandStack[len(operandStack)-1]
			val2 := operandStack[len(operandStack)-2]

			operandStack = operandStack[:len(operandStack)-2]

			operandStack = append(operandStack, fmt.Sprintf("%s%s%s", val1, val2, currChar))
		}
	}

	log.Print(operandStack[0])
}

package stacksandqueues

import (
	"log"
	"strconv"
)

/*
Infix Evaluation

1. You are given an infix expression.
2. You are required to evaluate and print it's value.

Constraints
1. Expression is balanced
2. The only operators used are +, -, *, /
3. Opening and closing brackets - () - are used to impact precedence of operations
4. + and - have equal precedence which is less than * and /. * and / also have equal precedence.
5. In two operators of equal precedence give preference to the one on left.
6. All operands are single digit numbers.

Sample Input
2 + 6 * 4 / 8 - 3

Sample Output
2
*/

func TestInfixExp() {
	// exp := "2 + 6 * 4 / 8 - 3"
	exp := "(8 - 3) * (2 - 8)"
	log.Print(infixExp(exp))
}

func infixExp(exp string) int {
	opPrecedence := make(map[string]int)
	opPrecedence["+"] = 1
	opPrecedence["-"] = 1
	opPrecedence["*"] = 2
	opPrecedence["/"] = 2
	opPrecedence["("] = 3
	opPrecedence[")"] = 3

	opStack := make([]string, 0)
	operandStack := make([]int, 0)

	for i := 0; i < len(exp); i++ {
		curChar := string(exp[i])

		if curChar == " " {
			continue
		} else if opPrecedence[curChar] != 0 {
			if (len(opStack) == 0) ||
				(len(opStack) > 0 && opPrecedence[opStack[len(opStack)-1]] < opPrecedence[curChar]) {
				if curChar == ")" {
					for opStack[len(opStack)-1] != "(" {
						EvaluateExp(&opStack, &operandStack)
					}
					opStack = opStack[:len(opStack)-1]
				} else {
					opStack = append(opStack, curChar)
				}
			} else {
				for len(opStack) > 0 &&
					opPrecedence[opStack[len(opStack)-1]] >= opPrecedence[curChar] && opStack[len(opStack)-1] != "(" {
					EvaluateExp(&opStack, &operandStack)
				}

				opStack = append(opStack, curChar)
			}
		} else {
			val, _ := strconv.Atoi(curChar)
			operandStack = append(operandStack, val)
		}
	}

	for len(opStack) > 0 {
		EvaluateExp(&opStack, &operandStack)
	}

	return operandStack[0]
}

func EvaluateExp(opStack *[]string, operandStack *[]int) {
	tempOp := (*opStack)[len(*opStack)-1]
	*opStack = (*opStack)[:len(*opStack)-1]

	tempVal2 := (*operandStack)[len(*operandStack)-1]
	tempVal1 := (*operandStack)[len(*operandStack)-2]
	*operandStack = (*operandStack)[:len(*operandStack)-2]

	*operandStack = append(*operandStack, CalculateExp(tempOp, tempVal1, tempVal2))
}

func CalculateExp(op string, val1 int, val2 int) int {
	switch op {
	case "+":
		return val1 + val2
	case "-":
		return val1 - val2
	case "*":
		return val1 * val2
	case "/":
		return val1 / val2
	default:
		return 0
	}
}

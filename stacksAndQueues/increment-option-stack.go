package stacksandqueues

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Design A Stack With Increment Operation

Design a stack which supports the following operations.

Implement the CustomStack class:

    1: void push(int x) Adds x to the top of the stack if the stack hasn't reached the maxSize.
    2: int pop() Pops and returns the top of stack or -1 if the stack is empty.
    3: void inc(int k, int val) Increments the bottom k elements of the stack by val. If there are less than k elements in the stack, just increment all the elements in the stack.


Sample Input
3
push 1
push 2
pop
push 2
push 3
push 4
increment 5 100
increment 2 100
pop
pop
pop
pop
exit

Sample Output
2
103
202
201
-1
exit
*/

type IncStack struct {
	ElemStack *utils.Stack
	IncStack  *utils.Stack
	Count     int
}

func initIncStack() *IncStack {
	return &IncStack{
		ElemStack: utils.InitStack(),
		IncStack:  utils.InitStack(),
		Count:     0,
	}
}

func (is *IncStack) push(val int) {

}

func TestIncrementStack() {

}

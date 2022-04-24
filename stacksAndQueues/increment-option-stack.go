package stacksandqueues

import (
	"errors"
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

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
	IncStack  []int
	Count     int
}

func initIncStack() *IncStack {
	return &IncStack{
		ElemStack: utils.InitStack(),
		IncStack:  []int{},
		Count:     0,
	}
}

func (is *IncStack) push(val int) {
	is.ElemStack.Push(val)
	is.IncStack = append(is.IncStack, 0)
	is.Count++
}

func (is *IncStack) pop() (int, error) {
	if is.Count == 0 {
		return -1, errors.New("stack is empty")
	}

	tempVal, _ := is.ElemStack.Pop()
	currVal := tempVal.(int) + is.IncStack[is.Count-1]
	if is.Count > 1 {
		is.IncStack[is.Count-2] = is.IncStack[is.Count-2] + is.IncStack[is.Count-1]
	}
	is.IncStack = is.IncStack[:is.Count-1]
	is.Count--

	return currVal, nil
}

func (is *IncStack) inc(k int, val int) {
	if k > is.Count {
		is.IncStack[is.Count-1] = val
		return
	}

	is.IncStack[k-1] = val
}

func TestIncrementStack() {
	incStack := initIncStack()
	incStack.push(1)
	incStack.push(2)

	val, _ := incStack.pop()
	log.Print(val)

	incStack.push(2)
	incStack.push(3)
	incStack.push(4)

	incStack.inc(5, 100)
	incStack.inc(2, 100)

	val, _ = incStack.pop()
	log.Print(val)
	val, _ = incStack.pop()
	log.Print(val)
	val, _ = incStack.pop()
	log.Print(val)
	val, _ = incStack.pop()
	log.Print(val)

	val, _ = incStack.pop()
	log.Print(val)

}

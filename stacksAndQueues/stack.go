package stacksandqueues

import (
	"fmt"
	"log"
)

type stack struct {
	Count    int
	Elements []int
}

func StackInit() *stack {
	return &stack{}
}

func (stack *stack) Push(val int) {
	stack.Count++
	stack.Elements = append(stack.Elements, val)
}

func (stack *stack) Pop() int {
	if stack.Count == 0 {
		log.Print("Stack Underflow")
		return -1
	}

	curVal := stack.Elements[stack.Count-1]
	stack.Elements = stack.Elements[:stack.Count-1]
	stack.Count--

	return curVal
}

func (stack *stack) Size() int {
	return stack.Count
}

func (stack *stack) GetTop() int {
	if stack.Count == 0 {
		log.Print("Stack Underflow")
		return -1
	}

	return stack.Elements[stack.Count-1]
}

func (stack *stack) Display() {
	for i := stack.Count - 1; i >= 0; i-- {
		fmt.Print(stack.Elements[i], "	")
	}
}

func TestStack() {
	s := StackInit()
	s.Push(10)
	s.Display()
	s.Push(20)
	s.Display()
	s.Push(30)
	s.Display()
	log.Print(s.GetTop())
	log.Print(s.Pop())
	log.Print(s.Pop())
	log.Print(s.Pop())
	log.Print(s.Pop())
}

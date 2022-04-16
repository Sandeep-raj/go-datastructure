package stacksandqueues

import "log"

/*
Two Stacks In An Array

1. You are required to complete the code of our TwoStack class. The class should implement LIFO behaviours for two stacks in the same array.
2. Here is the list of functions that you are supposed to complete
    2.1. push1, push2 -> Should accept new data for appropriate stack if there is
    space available in the underlying array or print "Stack overflow" otherwise.
    2.2. pop1, po2 -> Should remove and return data from appropriate stack if
    available or print "Stack underflow" otherwise and return -1.
    2.3. top1, top2 -> Should return data from appropriate stack if available or print
    "Stack underflow" otherwise and return -1.
    2.4. size1, size2 -> Should return the number of elements available in appropriate
     stack.
3. Input and Output is managed for you.

*/

type multiStacks struct {
	size int

	count1 int
	count2 int

	stack []int
}

func initmultistacks(size int) *multiStacks {
	return &multiStacks{
		size:   size,
		count1: 0,
		count2: 0,
		stack:  make([]int, size),
	}
}

func (mstack *multiStacks) push1(val int) {
	if mstack.count1+mstack.count2 == mstack.size {
		log.Print("stack overflow")
		return
	}

	mstack.count1++
	mstack.stack[(mstack.count1 - 1)] = val
}

func (mstack *multiStacks) push2(val int) {
	if mstack.count1+mstack.count2 == mstack.size {
		log.Print("stack overflow")
		return
	}

	mstack.count2++
	mstack.stack[(mstack.size - mstack.count2)] = val
}

func (mstack *multiStacks) pop1() int {
	if mstack.count1 == 0 {
		log.Print("stack underflow")
		return -1
	}
	currval := mstack.stack[mstack.count1-1]
	mstack.count1--
	return currval
}

func (mstack *multiStacks) pop2() int {
	if mstack.count2 == 0 {
		log.Print("stack underflow")
		return -1
	}
	currval := mstack.stack[mstack.size-mstack.count2]
	mstack.count2--
	return currval
}

func TestMultiStack() {
	ms := initmultistacks(5)
	ms.push1(10)
	ms.push1(20)
	ms.push2(30)
	ms.push2(40)
	ms.push1(50)

	ms.push2(60)
	log.Print(ms.pop1())
	ms.push1(60)

	log.Print(ms.pop1())
	log.Print(ms.pop1())
	log.Print(ms.pop1())
	log.Print(ms.pop1())

	log.Print(ms.pop2())
	log.Print(ms.pop2())
	log.Print(ms.pop2())
	log.Print(ms.pop2())
}

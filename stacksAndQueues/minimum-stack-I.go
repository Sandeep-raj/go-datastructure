package stacksandqueues

import "log"

/*
Minimum Stack - I

1. You are required to complete the code of our MinStack class.
2. As data members you've two stacks available in the class - one for data values and another for minimum values. (This is cryptic - take hint from video)
2. Here is the list of functions that you are supposed to complete
2.1. push -> Should accept new data in LIFO manner
2.2. pop -> Should remove and return data in LIFO manner. If not available, print "Stack underflow" and return -1.
2.3. top -> Should return data in LIFO manner. If not available, print "Stack underflow" and return -1.
2.4. size -> Should return the number of elements available in the stack
2.5. min -> Should return the smallest element available in the stack. If not available, print "Stack underflow" and return -1.
3. Input and Output is managed for you.
*/

func TestMinStack() {
	ms := stackInit()
	ms.push(10)
	ms.push(8)
	ms.push(15)
	ms.push(7)
	ms.push(3)
	ms.push(9)
	ms.push(1)

	log.Print(ms.getmin())
	log.Print(ms.pop())
	log.Print(ms.getmin())
	log.Print(ms.pop())
	log.Print(ms.getmin())
	log.Print(ms.pop())
	log.Print(ms.getmin())
	log.Print(ms.pop())
	log.Print(ms.getmin())
	log.Print(ms.pop())
	log.Print(ms.getmin())
	log.Print(ms.pop())
	log.Print(ms.getmin())
	log.Print(ms.pop())
}

type minstack struct {
	count    int
	elements []int
	min      int
}

func stackInit() *minstack {
	return &minstack{}
}

func (ms *minstack) push(val int) {
	if ms.count == 0 {
		ms.elements = append(ms.elements, val)
		ms.count++
		ms.min = val
	} else {
		if val > ms.min {
			ms.elements = append(ms.elements, val)
			ms.count++
		} else {
			ms.elements = append(ms.elements, 2*val-ms.min)
			ms.min = val
			ms.count++
		}
	}
}

func (ms *minstack) pop() int {
	currval := -1
	if ms.count > 0 {
		if ms.elements[ms.count-1] < ms.min {
			currval = ms.min
			ms.min = 2*ms.min - ms.elements[ms.count-1]
			ms.count--
		} else {
			currval = ms.elements[ms.count-1]
			ms.elements = ms.elements[:ms.count-1]
			ms.count--
		}
	} else {
		log.Print("Stack Underflow")
	}

	return currval
}

func (ms *minstack) getmin() int {
	return ms.min
}

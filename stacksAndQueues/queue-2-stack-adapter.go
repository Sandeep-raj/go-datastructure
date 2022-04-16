package stacksandqueues

import "log"

/*
Queue To Stack Adapter - Push Efficient

1. You are required to complete the code of our QueueToStackAdapter class.
2. As data members you've two queues available - mainQ and helperQ. mainQ is to contain data and helperQ is to assist in operations. (This is cryptic - take hint from video)
3. Here is the list of functions that you are supposed to complete
     3.1. push -> Should accept new data in LIFO manner.
     3.2. pop -> Should remove and return data in LIFO manner. If not available, print
      Stack underflow" and return -1.
     3.3. top -> Should return data in LIFO manner. If not available, print "Stack
     underflow" and return -1.
     3.4. size -> Should return the number of elements available in the stack.
4. Input and Output is managed for you.
*/

type q2sAdapter struct {
	count   int
	mainq   *queue
	helperq *queue
}

func InitQ2SAdapter() *q2sAdapter {
	return &q2sAdapter{
		count:   0,
		mainq:   InitQueue(),
		helperq: InitQueue(),
	}
}

func (stack *q2sAdapter) PushPE(val int) {
	stack.mainq.Add(val)
	stack.count++
}

func (stack *q2sAdapter) PopPE() int {
	if stack.count == 0 {
		log.Print("stack underflow")
		return -1
	}

	for stack.mainq.count > 1 {
		stack.helperq.Add(stack.mainq.Remove())
	}

	currVal := stack.mainq.Remove()
	stack.count--

	temp := stack.mainq
	stack.mainq = stack.helperq
	stack.helperq = temp

	return currVal
}

func (stack *q2sAdapter) PushPoE(val int) {
	stack.helperq.Add(val)

	for stack.mainq.count > 0 {
		stack.helperq.Add(stack.mainq.Remove())
	}

	temp := stack.mainq
	stack.mainq = stack.helperq
	stack.helperq = temp

	stack.count++
}

func (stack *q2sAdapter) PopPoE() int {
	if stack.count == 0 {
		log.Print("stack underflow")
		return -1
	}

	stack.count--
	return stack.mainq.Remove()
}

func TestQ2SAdapter() {
	s := InitQ2SAdapter()
	s.PushPoE(10)
	s.PushPoE(20)
	s.PushPoE(30)

	log.Print(s.PopPoE())
	log.Print(s.PopPoE())
	log.Print(s.PopPoE())
	s.PushPE(40)
	log.Print(s.PopPoE())
}

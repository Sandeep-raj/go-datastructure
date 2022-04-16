package stacksandqueues

import "log"

/*
Stack To Queue Adapter - Add Efficient

1. You are required to complete the code of our StackToQueueAdapter class. The class should mimic the behaviour of a Queue and implement FIFO semantic.
2. Here is the list of functions that are written in the class
    2.1. add -> Accepts new data if there is space available in the underlying array or
    print "Queue overflow" otherwise.
    2.2. remove -> Removes and returns value according to FIFO, if available or print
    "Queue underflow" otherwise and return -1.
    2.3. peek -> Returns value according to FIFO, if available or print "Queue
    underflow" otherwise and return -1.
    2.4. size -> Returns the number of elements available in the queue.
3. Input and Output is managed for you.

Stack To Queue Adapter - Remove Efficient


1. You are required to complete the code of our StackToQueueAdapter class. The class should mimic the behaviour of a Queue and implement FIFO semantic.
2. Here is the list of functions that are written in the class
    2.1. add -> Accepts new data if there is space available in the underlying array or
    print "Queue overflow" otherwise.
    2.2. remove -> Removes and returns value according to FIFO, if available or print
    "Queue underflow" otherwise and return -1.
    2.3. peek -> Returns value according to FIFO, if available or print "Queue
    underflow" otherwise and return -1.
    2.4. size -> Returns the number of elements available in the queue.
3. Input and Output is managed for you.
*/

type s2Queue struct {
	count   int
	mains   *stack
	helpers *stack
}

func InitS2Queue() *s2Queue {
	return &s2Queue{
		count:   0,
		mains:   StackInit(),
		helpers: StackInit(),
	}
}

func (queue *s2Queue) AddAE(val int) {
	queue.count++
	queue.mains.Push(val)
}

func (queue *s2Queue) RemoveAE() int {
	if queue.count == 0 {
		log.Print("queue empty")
		return -1
	}

	for queue.mains.Size() > 1 {
		queue.helpers.Push(queue.mains.Pop())
	}

	currval := queue.mains.Pop()
	queue.count--

	temp := queue.mains
	queue.mains = queue.helpers
	queue.helpers = temp

	return currval
}

func (queue *s2Queue) AddRE(val int) {
	for queue.mains.Size() > 0 {
		queue.helpers.Push(queue.mains.Pop())
	}

	queue.mains.Push(val)

	for queue.helpers.Size() > 0 {
		queue.mains.Push(queue.helpers.Pop())
	}

	queue.count++
}

func (queue *s2Queue) RemoveRE() int {
	currval := queue.mains.Pop()
	queue.count--
	return currval
}

func TestS2Queue() {
	q := InitS2Queue()
	q.AddRE(10)
	q.AddRE(20)
	q.AddRE(30)

	log.Print(q.RemoveRE())
	log.Print(q.RemoveRE())
	log.Print(q.RemoveRE())
	log.Print(q.RemoveRE())
}

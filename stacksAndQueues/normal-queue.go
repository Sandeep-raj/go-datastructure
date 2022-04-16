package stacksandqueues

import (
	"fmt"
	"log"
)

/*
Normal Queue

1. You are required to complete the code of our CustomQueue class. The class should mimic the behaviour of a Queue and implement FIFO semantic.
2. Here is the list of functions that you are supposed to complete
     2.1. add -> Should accept new data if there is space available in the underlying
     array or print "Queue overflow" otherwise.
     2.2. remove -> Should remove and return value according to FIFO, if available or
     print "Queue underflow" otherwise and return -1.
     2.3. peek -> Should return value according to FIFO, if available or print "Queue
     underflow" otherwise and return -1.
     2.4. size -> Should return the number of elements available in the queue.
     2.5. display -> Should print the elements of queue in FIFO manner (space-
     separated) ending with a line-break.
3. Input and Output is managed for you.
*/

type queue struct {
	count    int
	elements []int
}

func InitQueue() *queue {
	return &queue{}
}

func (q *queue) Add(val int) {
	q.elements = append(q.elements, val)
	q.count++
}

func (q *queue) Remove() int {
	if q.count == 0 {
		log.Print("queue underflow")
		return -1
	}

	currVal := q.elements[0]
	q.elements = q.elements[1:]
	q.count--

	return currVal
}

func (q *queue) Peek() int {
	if q.count == 0 {
		log.Print("queue underflow")
		return -1
	}

	return q.elements[0]
}

func (q *queue) Size() int {
	return q.count
}

func (q *queue) Display() {
	if q.count == 0 {
		log.Print("queue empty")
	}

	for i := 0; i < q.count; i++ {
		fmt.Print(q.elements[i], "	")
	}
	log.Println()
}

func TestQueue() {
	q := InitQueue()
	q.Add(10)
	q.Add(20)
	q.Add(30)
	log.Print(q.Peek())
	q.Remove()
	q.Display()
	q.Remove()
	q.Add(40)
	log.Print(q.Peek())
	q.Display()
	q.Remove()
	q.Remove()
}

package stacksandqueues

import (
	"errors"
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Design Circular Deque

Design your implementation of the circular double-ended queue (deque).

Your implementation should support following operations:

1. MyCircularDeque(): Initializes the deque object.
2. insertFront(): Adds an item at the front of Deque.
3. insertLast(): Adds an item at the rear of Deque.
4. deleteFront(): Deletes the front item from the Deque and return it's value. If the deque is empty, return null.
5. deleteLast(): Deletes the last item from Deque and return it's value. If the deque is empty, return null.
6. getFront(): Gets the front item from the Deque. If the deque is empty, return null.
7. getRear(): Gets the last item from Deque. If the deque is empty, return null.
8. isEmpty(): Checks whether Deque is empty or not.

Sample Input
insertLast 1
insertLast 2
insertFront 3
insertFront 4
getRear
deleteLast
getRear
deleteLast
insertFront 4
getFront

Sample Output
2
2
1
1
4
*/

type deque struct {
	count int
	llist utils.LinkedList
}

func myCircularDeque() deque {
	return deque{
		count: 0,
		llist: *utils.InitLL(),
	}
}

func (dq *deque) insertFront(val int) {
	dq.count++
	node := utils.LLNode{
		Key:  val,
		Data: val,
	}
	dq.llist.AddFront(node)
}

func (dq *deque) insertLast(val int) {
	dq.count++
	node := utils.LLNode{
		Key:  val,
		Data: val,
	}
	dq.llist.AddRear(node)
}

func (dq *deque) deleteFront() (int, error) {
	if dq.count == 0 {
		return -1, errors.New("empty queue")
	}
	val, _ := dq.llist.RemoveFront()
	dq.count--
	return val.Data.(int), nil
}

func (dq *deque) deleteLast() (int, error) {
	if dq.count == 0 {
		return -1, errors.New("empty queue")
	}
	val, _ := dq.llist.RemoveRear()
	dq.count--
	return val.Data.(int), nil
}

func (dq *deque) getFront() (int, error) {
	if dq.count == 0 {
		return -1, errors.New("empty queue")
	}
	val, _ := dq.llist.PeekFront()
	return val.Data.(int), nil
}

func (dq *deque) getRear() (int, error) {
	if dq.count == 0 {
		return -1, errors.New("empty queue")
	}
	val, _ := dq.llist.PeekRear()
	return val.Data.(int), nil
}

func (dq *deque) isEmpty() bool {
	return dq.llist.GetSize() == 0
}

func TestDeque() {
	dq := myCircularDeque()

	dq.insertFront(2)
	dq.insertLast(1)
	dq.insertFront(3)
	dq.insertLast(4)
	dq.insertFront(5)
	dq.insertLast(6)
	dq.insertFront(7)
	dq.insertLast(8)

	dq.llist.PrintLL()

	for !dq.isEmpty() {
		val, _ := dq.getFront()
		log.Printf("%d deleted from the front", val)
		dq.deleteFront()

		val, _ = dq.getRear()
		log.Printf("%d deleted from the rear", val)
		dq.deleteLast()
	}
}

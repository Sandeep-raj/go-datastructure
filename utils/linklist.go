package utils

import (
	"errors"
	"log"
)

type LLNode struct {
	Key  int
	Data interface{}
	Next *LLNode
}

type LinkedList struct {
	Count int
	Head  *LLNode
	Rear  *LLNode
}

func InitLL() *LinkedList {
	return &LinkedList{
		Count: 0,
		Head:  nil,
		Rear:  nil,
	}
}

func InitLLNode(key int, data interface{}) LLNode {
	return LLNode{
		Key:  key,
		Data: data,
		Next: nil,
	}
}

func (ll *LinkedList) AddFront(data LLNode) int {
	if ll.Count == 0 {
		ll.Head = &data
		ll.Rear = &data
		ll.Count++
		return ll.Count
	} else {
		data.Next = ll.Head
		ll.Head = &data
		ll.Count++
		return ll.Count
	}
}

func (ll *LinkedList) AddRear(data LLNode) int {
	if ll.Count == 0 {
		ll.Head = &data
		ll.Rear = &data
		ll.Count++
		return ll.Count
	} else {
		ll.Rear.Next = &data
		ll.Rear = ll.Rear.Next
		ll.Count++
		return ll.Count
	}
}

func (ll *LinkedList) RemoveFront() (*LLNode, error) {
	if ll.Count == 0 {
		return nil, errors.New("empty linklist")
	} else if ll.Count == 1 {
		currVal := ll.Head
		ll.Head = nil
		ll.Rear = nil
		ll.Count--
		return currVal, nil
	} else {
		currVal := ll.Head
		ll.Head = ll.Head.Next
		ll.Count--

		currVal.Next = nil
		return currVal, nil
	}
}

func (ll *LinkedList) RemoveRear() (*LLNode, error) {
	if ll.Count == 0 {
		return nil, errors.New("empty linklist")
	} else if ll.Count == 1 {
		currVal := ll.Head
		ll.Head = nil
		ll.Rear = nil
		ll.Count--
		return currVal, nil
	} else {
		currVal := ll.Rear
		temp := ll.Head
		for temp.Next != currVal {
			temp = temp.Next
		}
		temp.Next = nil
		ll.Rear = temp
		ll.Count--

		currVal.Next = nil
		return currVal, nil
	}
}

func (ll *LinkedList) GetSize() int {
	return ll.Count
}

func (ll *LinkedList) GetElement(key int) (*LLNode, error) {
	temp := ll.Head
	var res *LLNode
	if temp.Key == key {
		ll.Head = temp.Next
		temp.Next = nil
		ll.Count--
		res = temp
	} else {
		for temp.Next != nil && temp.Next.Key != key {
			temp = temp.Next
		}

		if temp.Next != nil && temp.Next.Key == key {
			res = temp.Next
			temp.Next = temp.Next.Next
			res.Next = nil
			ll.Count--
		} else {
			return nil, errors.New("element not found")
		}
	}

	return res, nil
}

func (ll *LinkedList) ElementExist(key int) bool {
	if ll.Count == 0 {
		return false
	}
	temp := ll.Head
	for temp != nil {
		if temp.Key == key {
			return true
		}
		temp = temp.Next
	}

	return false
}

func (ll *LinkedList) PeekFront() (*LLNode, error) {
	if ll.Count == 0 {
		return nil, errors.New("empty linked list")
	} else {
		return ll.Head, nil
	}
}

func (ll *LinkedList) PeekRear() (*LLNode, error) {
	if ll.Count == 0 {
		return nil, errors.New("empty linked list")
	} else {
		return ll.Rear, nil
	}
}

func (ll *LinkedList) PrintLL() {
	temp := ll.Head
	for temp != nil {
		log.Print(temp.Key)
		temp = temp.Next
	}
}

package utils

import (
	"errors"
	"log"
)

type Dll_node struct {
	Key  int
	Data interface{}
	Next *Dll_node
	Prev *Dll_node
}

type Dll struct {
	Count int
	Front *Dll_node
	Rear  *Dll_node
}

func InitDllNode(key int, data interface{}) *Dll_node {
	return &Dll_node{
		Key:  key,
		Data: data,
		Next: nil,
		Prev: nil,
	}
}

func InitDll() *Dll {
	return &Dll{
		Count: 0,
		Front: nil,
		Rear:  nil,
	}
}

// Add first DLL
func (dll *Dll) DllInsertFront(node *Dll_node) {
	if dll.Count == 0 {
		dll.Front = node
		dll.Rear = node
	} else {
		node.Next = dll.Front
		dll.Front.Prev = node
		dll.Front = node
	}

	dll.Count++
}

// ADD last DLL
func (dll *Dll) DllInsertRear(node *Dll_node) {
	if dll.Count == 0 {
		dll.Front = node
		dll.Rear = node
	} else {
		node.Next = nil
		node.Prev = nil

		dll.Rear.Next = node
		node.Prev = dll.Rear
		dll.Rear = node
	}

	dll.Count++
}

// Remove first DLL
func (dll *Dll) DllRemoveFront() (*Dll_node, error) {
	if dll.Count == 0 {
		return nil, errors.New("doubly linked list empty")
	} else if dll.Count == 1 {
		temp := dll.Front
		dll.Count = 0
		dll.Front = nil
		dll.Rear = nil

		return temp, nil
	} else {
		temp := dll.Front
		dll.Front = temp.Next
		dll.Front.Prev = nil

		dll.Count--
		temp.Next = nil
		temp.Prev = nil

		return temp, nil
	}
}

// Remove last DLL
func (dll *Dll) DllRemoveRear() (*Dll_node, error) {
	if dll.Count == 0 {
		return nil, errors.New("doubly linked list empty")
	} else if dll.Count == 1 {
		temp := dll.Rear
		dll.Count = 0
		dll.Front = nil
		dll.Rear = nil

		return temp, nil
	} else {
		temp := dll.Rear
		dll.Rear = temp.Prev
		dll.Rear.Next = nil

		dll.Count--
		temp.Next = nil
		temp.Prev = nil

		return temp, nil
	}
}

//Remove node DLL
func (dll *Dll) DllRemoveNode(node *Dll_node) {
	prevNode := node.Prev
	nextNode := node.Next

	if prevNode == nil && nextNode == nil {

		dll.Front = nil
		dll.Rear = nil
	} else if prevNode == nil {
		dll.Front = nextNode
		nextNode.Prev = nil

	} else if nextNode == nil {
		dll.Rear = prevNode
		prevNode.Next = nil

	} else {
		prevNode.Next = nextNode
		nextNode.Prev = prevNode

	}

	dll.Count--
}

func (dll *Dll) GetSize() int {
	return dll.Count
}

// Get At DLL
func (dll *Dll) GetElement(key int) (*Dll_node, error) {
	if dll.Count == 0 {
		return nil, errors.New("doubly linklist empty")
	} else if dll.Count == 1 && dll.Front.Key == key {
		temp := dll.Front
		dll.Front = nil
		dll.Rear = nil
		dll.Count--

		return temp, nil
	} else {
		temp := dll.Front

		for temp != nil {
			if temp.Key == key {
				break
			}
			temp = temp.Next
		}

		if temp == nil {
			return nil, errors.New("element not found")
		}

		if temp == dll.Front {
			dll.Front = dll.Front.Next
			dll.Front.Prev = nil
		} else if temp == dll.Rear {
			dll.Rear = dll.Rear.Prev
			dll.Rear.Next = nil
		} else {
			prevEle := temp.Prev
			nextEle := temp.Next

			prevEle.Next = nextEle
			nextEle.Prev = prevEle
		}

		dll.Count--
		return temp, nil
	}
}

func (dll *Dll) DllElementExist(key int) bool {
	temp := dll.Front

	for temp != nil {
		if temp.Key == key {
			return true
		}
		temp = temp.Next
	}

	return false
}

// Get first
func (dll *Dll) DllPeakFront() *Dll_node {
	if dll.Count == 0 {
		return nil
	}

	return dll.Front
}

// Get last
func (dll *Dll) DllPeakRear() *Dll_node {
	if dll.Count == 0 {
		return nil
	}

	return dll.Rear
}

// Print DLL
func (dll *Dll) PrintDll() {
	temp := dll.Front

	for temp != nil {
		log.Print(temp.Key)
		temp = temp.Next
	}
}

// Add At DLL
// Remove At DLL
// Add before DLL
// Add after DLL
// Remove after DLL
// Remove before DLL

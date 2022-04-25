package stacksandqueues

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Max-stack

Design a max stack data structure that supports the stack operations and supports finding the stack's maximum element.

Implement the MaxStack class:
1. MaxStack() Initializes the stack object.
2. void push(int x) Pushes element x onto the stack.
3. int pop() Removes the element on top of the stack and returns it.
4. int top() Gets the element on the top of the stack without removing it.
5. int peekMax() Retrieves the maximum element in the stack without removing it.
6. int popMax() Retrieves the maximum element in the stack and removes it. If there is more than one maximum element, only remove the top-most one.

Sample Input
push 5
push 1
push 5
top
popMax
top
peekMax
pop
top

Sample Output
5
5
1
5
1
5
*/

func TestMaxStack() {
	ms := initmaxstack()

	ms.push(5)
	ms.push(1)
	ms.push(5)

	log.Print(ms.top())
	ms.push(3)
	log.Print(ms.popmax())
	log.Print(ms.top())

	log.Print(ms.peekMax())
	log.Print(ms.popmax())
	log.Print(ms.peekMax())
	log.Print(ms.pop())
	log.Print(ms.pop())
}

type maxstack struct {
	count    int
	stack    *utils.Dll
	stackmap *utils.BST
}

func initmaxstack() *maxstack {
	return &maxstack{
		count:    0,
		stack:    utils.InitDll(),
		stackmap: utils.InitBST(),
	}
}

func (ms *maxstack) push(val int) {
	data := utils.InitDllNode(val, val)
	ms.stack.DllInsertRear(data)
	currNode := ms.stackmap.BSTSearch(val)
	if currNode == nil {
		ms.stackmap.BSTInsert(*utils.InitBSTNode(val, []*utils.Dll_node{data}))
	} else {
		currNode.Data = append(currNode.Data.([]*utils.Dll_node), data)
	}

	ms.count++
}

func (ms *maxstack) pop() int {
	if ms.count > 0 {
		data, _ := ms.stack.DllRemoveRear()

		currnode := ms.stackmap.BSTSearch(data.Key)
		if len(currnode.Data.([]*utils.Dll_node)) == 1 {
			ms.stackmap.BSTDelete(currnode.Key)
		} else {
			currnode.Data = currnode.Data.([]*utils.Dll_node)[:len(currnode.Data.([]*utils.Dll_node))-1]
		}
		ms.count--

		return data.Key
	}

	return -1
}

func (ms *maxstack) top() int {
	if ms.count > 0 {
		return ms.stack.DllPeakRear().Key
	}

	return -1
}

func (ms *maxstack) peekMax() int {
	if ms.count == 0 {
		return -1
	}

	currNode := ms.stackmap.BSTMax()
	return currNode.Data.([]*utils.Dll_node)[len(currNode.Data.([]*utils.Dll_node))-1].Key
}

func (ms *maxstack) popmax() int {
	if ms.count == 0 {
		return -1
	}

	currNode := ms.stackmap.BSTMax()

	dllnode := currNode.Data.([]*utils.Dll_node)[len(currNode.Data.([]*utils.Dll_node))-1]

	if len(currNode.Data.([]*utils.Dll_node)) == 1 {
		ms.stackmap.BSTDelete(currNode.Key)
	} else {
		currNode.Data = currNode.Data.([]*utils.Dll_node)[:len(currNode.Data.([]*utils.Dll_node))-1]
	}

	ms.stack.DllRemoveNode(dllnode)
	ms.count--

	return dllnode.Key

}

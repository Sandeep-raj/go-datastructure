package linkedlist

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Merge Sort A Linked List

1. You are given a partially written LinkedList class.
2. You are required to complete the body of mergeSort function. The function is static and is passed the head and tail of an unsorted list. The function is expected to return a new sorted list. The original list must not change.
3. Input and Output is managed for you.

Sample Input
6
10 2 19 22 3 7

Sample Output
2 3 7 10 19 22
10 2 19 22 3 7
*/

func TestMergeSortLL() {
	ll := utils.InitLL()

	ll.AddRear(utils.InitLLNode(10, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(19, ""))
	ll.AddRear(utils.InitLLNode(22, ""))
	ll.AddRear(utils.InitLLNode(3, ""))
	ll.AddRear(utils.InitLLNode(7, ""))

	ll.PrintLL()

	resll := mergeSortLL(ll.Head, ll.Rear)

	log.Print("----------------------------------")
	resll.PrintLL()
}

func mergeSortLL(startPtr *utils.LLNode, endPtr *utils.LLNode) *utils.LinkedList {
	if startPtr == endPtr {
		resll := utils.InitLL()
		resll.AddRear(utils.InitLLNode(startPtr.Key, startPtr.Data))
		return resll
	}

	mid1 := getMidNode(startPtr, endPtr)
	ll1 := mergeSortLL(startPtr, mid1)
	ll2 := mergeSortLL(mid1.Next, endPtr)

	mergeLL(ll1, ll2)
	return ll1
}

func getMidNode(startPtr *utils.LLNode, endPtr *utils.LLNode) *utils.LLNode {
	slowPtr := startPtr
	fastPtr := startPtr

	for fastPtr.Next != nil && fastPtr != endPtr && fastPtr.Next != endPtr && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	return slowPtr
}

package linkedlist

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Mid Of Linked List

You are required to complete the body of mid function. The function should be an iterative function and should return the mid of linked list. Also, make sure to not use size data member directly or indirectly (by calculating size via making a traversal). In linked list of odd size, mid is unambigous. In linked list of even size, consider end of first half as mid.
*/

func TestMidLL() {
	ll := utils.InitLL()
	ll.AddFront(utils.InitLLNode(10, "test"))
	ll.AddFront(utils.InitLLNode(20, "test"))
	ll.AddRear(utils.InitLLNode(30, "test"))
	ll.AddRear(utils.InitLLNode(40, "test"))
	ll.AddRear(utils.InitLLNode(50, "test"))
	ll.AddFront(utils.InitLLNode(60, "test"))
	ll.AddFront(utils.InitLLNode(100, "test"))
	ll.AddFront(utils.InitLLNode(200, "test"))
	ll.AddRear(utils.InitLLNode(300, "test"))
	ll.AddRear(utils.InitLLNode(400, "test"))

	midLL(ll)

}

func midLL(ll *utils.LinkedList) {
	fastPtr := ll.Head
	slowPtr := ll.Head

	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	log.Print(slowPtr.Key)
}

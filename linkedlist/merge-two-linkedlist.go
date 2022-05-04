package linkedlist

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Merge Two Sorted Linked Lists

1. You are given a partially written LinkedList class.
2. You are required to complete the body of mergeTwoSortedLists function. The function is static and is passed two lists which are sorted. The function is expected to return a new sorted list containing elements of both lists. Original lists must stay as they were.
3. Input and Output is managed for you.

Sample Input
5
10 20 30 40 50
10
7 9 12 15 37 43 44 48 52 56

Sample Output
7 9 10 12 15 20 30 37 40 43 44 48 50 52 56
10 20 30 40 50
7 9 12 15 37 43 44 48 52 56
*/

func TestMergeLL() {
	ll1 := utils.InitLL()
	ll1.AddRear(utils.InitLLNode(10, ""))
	ll1.AddRear(utils.InitLLNode(20, ""))
	ll1.AddRear(utils.InitLLNode(30, ""))
	ll1.AddRear(utils.InitLLNode(40, ""))
	ll1.AddRear(utils.InitLLNode(50, ""))

	ll2 := utils.InitLL()
	ll2.AddRear(utils.InitLLNode(7, ""))
	ll2.AddRear(utils.InitLLNode(9, ""))
	ll2.AddRear(utils.InitLLNode(12, ""))
	ll2.AddRear(utils.InitLLNode(15, ""))
	ll2.AddRear(utils.InitLLNode(37, ""))
	ll2.AddRear(utils.InitLLNode(43, ""))
	ll2.AddRear(utils.InitLLNode(44, ""))
	ll2.AddRear(utils.InitLLNode(48, ""))
	ll2.AddRear(utils.InitLLNode(52, ""))
	ll2.AddRear(utils.InitLLNode(56, ""))

	mergeLL(ll1, ll2)

	ll1.PrintLL()
	log.Print("--------------------------------------")
	ll2.PrintLL()
}

func mergeLL(ll1 *utils.LinkedList, ll2 *utils.LinkedList) {
	ptr1 := ll1.Head
	ptr2 := ll2.Head

	if ptr1.Key > ptr2.Key {
		temp := utils.InitLLNode(ptr2.Key, ptr2.Data)
		temp.Next = ptr1
		ll1.Head = &temp

		ptr1 = ll1.Head
		ptr2 = ptr2.Next
		ll1.Count++
	}

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Key <= ptr2.Key {
			ptr1 = ptr1.Next
		} else {
			temp := utils.InitLLNode(ptr1.Key, ptr1.Data)
			temp.Next = ptr1.Next

			ptr1.Key = ptr2.Key
			ptr1.Data = ptr2.Data

			ptr1.Next = &temp

			ptr1 = ptr1.Next
			ptr2 = ptr2.Next

			ll1.Count++
		}

		if ptr1 != nil {
			ll1.Rear = ptr1
		}
	}

	if ptr2 != nil {
		for ptr2 != nil {
			ll1.AddRear(utils.InitLLNode(ptr2.Key, ptr2.Data))
			ptr2 = ptr2.Next
		}
	}
}

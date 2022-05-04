package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Odd Even Linked List

1. You are given a partially written LinkedList class.
2. You are required to complete the body of oddEven function. The function is expected to tweak the list such that all odd values are followed by all even values. The relative order of elements should not change. Also, take care of the cases when there are no odd or no even elements. Make sure to properly set head, tail and size as the function will be tested by calling addFirst and addLast.
3. Input and Output is managed for you.

Sample Input
7
2 8 9 1 5 4 3
10
100

Sample Output
2 8 9 1 5 4 3
9 1 5 3 2 8 4
10 9 1 5 3 2 8 4 100
*/

func TestOddEvenLL() {
	ll := utils.InitLL()

	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(8, ""))
	ll.AddRear(utils.InitLLNode(9, ""))
	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(4, ""))
	ll.AddRear(utils.InitLLNode(3, ""))

	oddEvenLL(ll)

	ll.AddFront(utils.InitLLNode(10, ""))
	ll.AddRear(utils.InitLLNode(100, ""))

	ll.PrintLL()
}

func oddEvenLL(ll *utils.LinkedList) {
	var oddHead *utils.LLNode
	var oddRear *utils.LLNode
	var evenHead *utils.LLNode
	var evenRear *utils.LLNode

	temp := ll.Head

	for temp != nil {
		if temp.Key%2 == 0 {
			if evenHead == nil {
				evenHead = temp
				evenRear = temp
			} else {
				evenRear.Next = temp
				evenRear = evenRear.Next
			}
			temp = temp.Next
			evenRear.Next = nil
		} else {
			if oddHead == nil {
				oddHead = temp
				oddRear = temp
			} else {
				oddRear.Next = temp
				oddRear = oddRear.Next
			}
			temp = temp.Next
			oddRear.Next = nil
		}
	}

	if oddHead == nil {
		ll.Head = evenHead
	} else {
		ll.Head = oddHead
	}

	if evenRear == nil {
		ll.Rear = oddRear
	} else {
		oddRear.Next = evenHead
		ll.Rear = evenRear
	}
}

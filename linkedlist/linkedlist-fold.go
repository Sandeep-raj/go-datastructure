package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Fold A Linked List

1. You are given a partially written LinkedList class.
2. You are required to complete the body of fold function. The function is expected to place last element after 1st element, 2nd last element after 2nd element and so on. For more insight check the example

Example 1
1->2->3->4->5
will fold as
1->5->2->4->3

Example 2
1->2->3->4->5->6
1->6->2->5->3->4

Sample Input
5
1 2 3 4 5
10
20

Sample Output
1 2 3 4 5
1 5 2 4 3
10 1 5 2 4 3 20
*/

func Testllfold() {
	ll := utils.InitLL()

	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(3, ""))
	ll.AddRear(utils.InitLLNode(4, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(6, ""))

	llfold(ll)

	ll.AddFront(utils.InitLLNode(15, ""))
	ll.AddRear(utils.InitLLNode(25, ""))
	ll.PrintLL()
}

func llfold(ll *utils.LinkedList) {
	lback, _ := llfoldhelper(ll.Head, ll.Head)

	ll.Rear = lback
}

func llfoldhelper(lfirst *utils.LLNode, lback *utils.LLNode) (*utils.LLNode, bool) {
	if lback == nil {
		return lfirst, false
	}

	currfirst, isfold := llfoldhelper(lfirst, lback.Next)
	if isfold {
		return currfirst, true
	}
	if currfirst.Next == lback || currfirst == lback {
		lback.Next = nil
		return lback, true
	} else {
		tempfirst := currfirst.Next
		currfirst.Next = lback
		lback.Next = tempfirst

		return tempfirst, false
	}
}

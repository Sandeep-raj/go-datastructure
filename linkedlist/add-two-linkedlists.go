package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Add Two Linked Lists

1. You are given a partially written LinkedList class.
2. You are required to complete the body of addLinkedLists function. The function is passed two linked lists which represent two numbers - the first element is the most significant digit and the last element is the least significant digit. The function is expected to add the two linked list and return a new linked list.

The following approaches are not allowed :
    1. Don't reverse the linked lists in order to access them from least significant digit
     to most significant.
     2. Don't use arrays or explicit extra memory.
     3. Don't convert linked lists to number, add them up and convert the result back
     to a linked list.

Hint - You are expected to take help of recursion to access digits from least significant to most significant. You have to tackle the challenge of unequal size of lists and manage carry where required.

Sample Input
1
1
3
9 9 9
10
20

Sample Output
1
9 9 9
1 0 0 0
10 1 0 0 0 20
*/

func TestAddTwoLL() {
	ll1 := utils.InitLL()
	ll1.AddRear(utils.InitLLNode(1, ""))
	ll1.AddRear(utils.InitLLNode(1, ""))
	ll1.AddRear(utils.InitLLNode(1, ""))

	ll2 := utils.InitLL()
	ll2.AddRear(utils.InitLLNode(9, ""))
	ll2.AddRear(utils.InitLLNode(9, ""))
	ll2.AddRear(utils.InitLLNode(9, ""))

	resll := add2ll(ll1, ll2)

	resll.AddFront(utils.InitLLNode(25, ""))
	resll.AddRear(utils.InitLLNode(15, ""))

	resll.PrintLL()
}

func add2ll(ll1 *utils.LinkedList, ll2 *utils.LinkedList) *utils.LinkedList {
	resll := utils.InitLL()
	if ll1.Count < ll2.Count {
		temp := ll1
		ll1 = ll2
		ll2 = temp
	}

	extra := add2llhelper(ll1.Head, ll2.Head, 0, ll1.Count-ll2.Count, resll)

	if extra > 0 {
		resll.AddFront(utils.InitLLNode(extra, ""))
	}

	return resll
}

func add2llhelper(node1 *utils.LLNode, node2 *utils.LLNode, currk int, k int, resll *utils.LinkedList) int {
	if node1 == nil && node2 == nil {
		return 0
	}

	var extra int

	if currk < k {
		extra = add2llhelper(node1.Next, node2, currk+1, k, resll)
		sum := node1.Key + extra
		tempex := sum / 10
		tempval := sum % 10

		resll.AddFront(utils.InitLLNode(tempval, ""))
		return tempex
	} else {
		extra = add2llhelper(node1.Next, node2.Next, currk+1, k, resll)

		sum := node1.Key + node2.Key + extra
		tempex := sum / 10
		tempval := sum % 10

		resll.AddFront(utils.InitLLNode(tempval, ""))
		return tempex
	}
}

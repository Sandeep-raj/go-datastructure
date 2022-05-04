package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Segregate Node Of Linkedlist Over Pivot Index

1. Given a singly linklist, Segregate Node of LinkedList over pivot index and return starting node of linkedlist.
2. pivot will be any random index in range of 0 to length Of Linkedlist
3. After segregation pivot Element should have to be present at correct position as in sorted linkedlist.

Input
1->5->2->9->5->14->11->1->10->10->1->3->null
11

Output
1->2->1->1->3->5->9->5->14->11->10->10->null

Example
Sample Input
12
1 5 2 9 5 14 11 1 10 10 1 3
7

Sample Output
1 1 1 5 2 9 5 14 11 10 10 3
*/

func TestSegNodePivot() {
	ll := utils.InitLL()

	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(9, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(14, ""))
	ll.AddRear(utils.InitLLNode(11, ""))
	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(10, ""))
	ll.AddRear(utils.InitLLNode(10, ""))
	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(3, ""))

	pivotIndex := 0
	temp := ll.Head

	for i := 0; i < pivotIndex; i++ {
		temp = temp.Next
	}

	ll.Head, ll.Rear = segPivotLL(ll.Head, temp)
	ll.PrintLL()
}

func segPivotLL(startNode *utils.LLNode, pivotNode *utils.LLNode) (*utils.LLNode, *utils.LLNode) {
	temp := startNode

	var largePtr *utils.LLNode
	var smallPtr *utils.LLNode

	var tempLarge *utils.LLNode
	var tempSmall *utils.LLNode

	for temp != nil {
		if temp == pivotNode {
			temp = temp.Next
			continue
		}

		currTemp := temp.Next
		temp.Next = nil

		if temp.Key > pivotNode.Key {
			if largePtr == nil {
				largePtr = temp
				tempLarge = temp
			} else {
				tempLarge.Next = temp
				tempLarge = temp
			}
		} else {
			if smallPtr == nil {
				smallPtr = temp
				tempSmall = temp
			} else {
				tempSmall.Next = temp
				tempSmall = temp
			}
		}

		temp = currTemp
	}

	if smallPtr != nil && largePtr != nil {
		tempSmall.Next = pivotNode
		pivotNode.Next = largePtr
		return smallPtr, tempLarge
	} else if smallPtr != nil && largePtr == nil {
		tempSmall.Next = pivotNode
		pivotNode.Next = nil
		return smallPtr, pivotNode
	} else if smallPtr == nil && largePtr != nil {
		pivotNode.Next = largePtr
		tempLarge.Next = nil
		return pivotNode, tempLarge
	} else {
		return pivotNode, pivotNode
	}
}

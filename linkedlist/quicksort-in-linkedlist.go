package linkedlist

import (
	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Quicksort In Linkedlist

1. Given the head of a linked list, return the list after sorting it in increasing order.
2. You must apply quick sort.
3. Time Complexity : O(nlogn)
4. Space Complexity : constant space

Input
1->7->2->6->3->5->4->null

Output
1->2->3->4->5->6->7->null

Example
Sample Input
4
0
6
7
5

Sample Output
0 5 6 7
*/

func TestQuickSortLL() {
	ll := utils.InitLL()

	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(7, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(6, ""))
	ll.AddRear(utils.InitLLNode(3, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(4, ""))

	quicksortLL(ll)
	ll.PrintLL()
}

func quicksortLL(ll *utils.LinkedList) {
	resNode := quickSortLLhelper(ll.Head)
	ll.Head = resNode[0]
	ll.Rear = resNode[1]
}

func quickSortLLhelper(startnode *utils.LLNode) []*utils.LLNode {
	if startnode == nil {
		return nil
	}

	resNode := segOnPivot(startnode)

	leftArr := quickSortLLhelper(resNode[0])
	rightArr := quickSortLLhelper(resNode[2])

	if leftArr != nil && rightArr != nil {
		leftArr[1].Next = resNode[1]
		resNode[1].Next = rightArr[0]
		return []*utils.LLNode{leftArr[0], rightArr[1]}
	} else if leftArr != nil && rightArr == nil {
		leftArr[1].Next = resNode[1]
		return []*utils.LLNode{leftArr[0], resNode[1]}
	} else if leftArr == nil && rightArr != nil {
		resNode[1].Next = rightArr[0]
		return []*utils.LLNode{resNode[1], rightArr[1]}
	} else {
		return []*utils.LLNode{resNode[1], resNode[1]}
	}
}

func segOnPivot(startNode *utils.LLNode) []*utils.LLNode {
	if startNode == nil {
		return nil
	}

	var smallHead *utils.LLNode
	var smallTail *utils.LLNode
	var largeHead *utils.LLNode
	var largeTail *utils.LLNode

	pivotNode := startNode
	temp := startNode.Next
	pivotNode.Next = nil

	for temp != nil {
		nextTemp := temp.Next
		temp.Next = nil

		if temp.Key > pivotNode.Key {
			if largeHead == nil {
				largeHead = temp
				largeTail = temp
			} else {
				largeTail.Next = temp
				largeTail = temp
			}
		} else {
			if smallHead == nil {
				smallHead = temp
				smallTail = temp
			} else {
				smallTail.Next = temp
				smallTail = temp
			}
		}
		temp = nextTemp
	}

	return []*utils.LLNode{smallHead, pivotNode, largeHead}
}

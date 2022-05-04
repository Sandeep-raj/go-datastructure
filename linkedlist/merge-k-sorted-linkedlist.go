package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Merge K Sorted LinkedList

You are given an array of k linked-lists, each linked-list is sorted in increasing order.
Merge all the linkedlists into one sorted linkedlist and return it.

Input
3 sorted linkedlist :
{
    0->0->0->null,
    0->0->1->1->1->2->2->4->null,
    0->0->0->0->5->5->6->null
}

Output
after merging them :
0->0->0->0->0->0->0->0->0->1->1->1->2->2->4->5->5->6->null

Example
Sample Input
3
3
0 0 0
8
0 0 1 1 1 2 2 4
7
0 0 0 0 5 5 6

Sample Output
0 0 0 0 0 0 0 0 0 1 1 1 2 2 4 5 5 6
*/

func TestMergeKLL() {
	k := 3
	llArr := make([]*utils.LinkedList, k)

	ll1 := utils.InitLL()
	ll1.AddRear(utils.InitLLNode(0, ""))
	ll1.AddRear(utils.InitLLNode(0, ""))
	ll1.AddRear(utils.InitLLNode(0, ""))

	ll2 := utils.InitLL()
	ll2.AddRear(utils.InitLLNode(0, ""))
	ll2.AddRear(utils.InitLLNode(0, ""))
	ll2.AddRear(utils.InitLLNode(1, ""))
	ll2.AddRear(utils.InitLLNode(1, ""))
	ll2.AddRear(utils.InitLLNode(1, ""))
	ll2.AddRear(utils.InitLLNode(2, ""))
	ll2.AddRear(utils.InitLLNode(2, ""))
	ll2.AddRear(utils.InitLLNode(4, ""))

	ll3 := utils.InitLL()
	ll3.AddRear(utils.InitLLNode(0, ""))
	ll3.AddRear(utils.InitLLNode(0, ""))
	ll3.AddRear(utils.InitLLNode(0, ""))
	ll3.AddRear(utils.InitLLNode(0, ""))
	ll3.AddRear(utils.InitLLNode(5, ""))
	ll3.AddRear(utils.InitLLNode(5, ""))
	ll3.AddRear(utils.InitLLNode(6, ""))

	llArr[0] = ll1
	llArr[1] = ll2
	llArr[2] = ll3

	resLL := mergeKSortedLL([]*utils.LLNode{ll1.Head, ll2.Head, ll3.Head})
	resLL.PrintLL()
}

func mergeKSortedLL(llArr []*utils.LLNode) *utils.LinkedList {
	resLL := utils.InitLL()

	for {
		minIndex := -1
		for i := 0; i < len(llArr); i++ {
			if llArr[i] != nil {
				if minIndex == -1 {
					minIndex = i
				} else if llArr[minIndex].Key > llArr[i].Key {
					minIndex = i
				}
			}
		}

		if minIndex == -1 {
			break
		}

		if resLL.Count == 0 {
			resLL.Head = llArr[minIndex]
			resLL.Rear = llArr[minIndex]
			resLL.Count++
			llArr[minIndex] = llArr[minIndex].Next
		} else {
			resLL.Rear.Next = llArr[minIndex]
			resLL.Rear = llArr[minIndex]
			resLL.Count++
			llArr[minIndex] = llArr[minIndex].Next
		}
	}

	return resLL
}

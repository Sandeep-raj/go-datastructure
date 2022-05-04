package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Reverse In Range

1. Given a singly linklist, Reverse a linkedlist from position starting position to end position.
2. Do it in one-pass. without using any extra space.
3. Indexing start from numeric 1.

Input
8->8->14->1->10->12->null
3
5

Output
8->8->10->1->14->12->null

Example
Sample Input
6
8 8 14 1 10 12
3
5

Sample Output
8 8 10 1 14 12
*/

func TestRevRange() {
	ll := utils.InitLL()

	ll.AddRear(utils.InitLLNode(8, ""))
	ll.AddRear(utils.InitLLNode(8, ""))
	ll.AddRear(utils.InitLLNode(14, ""))
	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(10, ""))
	ll.AddRear(utils.InitLLNode(12, ""))
	n := 1
	m := 6

	revRange(ll, n, m)
	ll.PrintLL()
}

func revRange(ll *utils.LinkedList, n int, m int) {
	k := 1
	temp := ll.Head
	var starttemp *utils.LLNode
	var midstart *utils.LLNode
	var midend *utils.LLNode
	var endtemp *utils.LLNode

	for temp != nil {
		if k < n {
			starttemp = temp
			temp = temp.Next
		} else if k >= n && k <= m {
			currtemp := temp.Next
			temp.Next = midstart
			midstart = temp

			if midend == nil {
				midend = midstart
			}
			temp = currtemp
		} else {
			endtemp = temp
			break
		}
		k++
	}

	if starttemp == nil {
		ll.Head = midstart
	} else {
		starttemp.Next = midstart
	}

	if endtemp == nil {
		ll.Rear = midend
	} else {
		midend.Next = endtemp
	}

}

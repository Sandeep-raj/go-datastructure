package linkedlist

import (
	"math/rand"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Segregate 012 Node Of Linkedlist By Swapping Data

1. Given a singly linklist, Segregate 012 Node of LinkedList and return pivot node of linkedlist.
2. After segregation zero nodes should come first and then ones node followed by two's nodes.
3. You are only allowed to swap data not swap nodes.

Input
1->0->1->0->0->1->2->1->1->1->2->1->1->null

Output
0->0->0->1->1->1->1->1->1->1->1->2->2->null

Example
Sample Input
17
2 2 0 2 1 0 0 2 2 1 2 1 2 0 1 0 0

Sample Output
0 0 0 0 0 0 1 1 1 1 2 2 2 2 2 2 2
*/

func TestSeg012LL() {
	ll := utils.InitLL()
	for i := 0; i < 15; i++ {
		val := rand.Intn(3)
		ll.AddRear(utils.InitLLNode(val, ""))
	}

	seg012LL(ll)
	ll.PrintLL()
}

func seg012LL(ll *utils.LinkedList) {

	heads := make([]*utils.LLNode, 3)
	tails := make([]*utils.LLNode, 3)

	temp := ll.Head

	for temp != nil {
		currtemp := temp.Next
		temp.Next = nil

		switch temp.Key {
		case 0:
			if heads[0] == nil {
				heads[0] = temp
				tails[0] = temp
			} else {
				tails[0].Next = temp
				tails[0] = temp
			}
		case 1:
			if heads[1] == nil {
				heads[1] = temp
				tails[1] = temp
			} else {
				tails[1].Next = temp
				tails[1] = temp
			}
		case 2:
			if heads[2] == nil {
				heads[2] = temp
				tails[2] = temp
			} else {
				tails[2].Next = temp
				tails[2] = temp
			}
		}
		temp = currtemp
	}

	ll.Head = nil
	ll.Rear = nil

	var prev *utils.LLNode
	for i := 0; i < 3; i++ {
		if heads[i] != nil && ll.Head == nil {
			ll.Head = heads[i]
		}

		if tails[i] != nil {
			ll.Rear = tails[i]
		}

		if prev != nil && heads[i] != nil {
			prev.Next = heads[i]
		}

		if tails[i] != nil {
			prev = tails[i]
		}
	}

}

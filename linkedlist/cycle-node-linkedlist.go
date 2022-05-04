package linkedlist

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Cycle Node In Linkedlist

1. Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
2. There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.

Notice that you should not modify the linked list.

Sample Input

8
1
18
1
8
-1
138
31
84
3

Sample Output
8
*/

func TestCycleNode() {
	inpVal := []int{8, 1, 18, 1, 8, -1, 138, 31, 84, 3}
	ll := utils.InitLL()

	for _, val := range inpVal {
		curNode := utils.InitLLNode(val, "")
		ll.AddRear(curNode)
	}

	temp := ll.Head
	for i := 0; i < 4; i++ {
		temp = temp.Next
	}

	ll.Rear.Next = temp

	log.Printf("%+v", cycleNodeLL(ll))

}

func cycleNodeLL(ll *utils.LinkedList) *utils.LLNode {
	slowPtr := ll.Head
	fastPtr := ll.Head

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if slowPtr == fastPtr {
			slowPtr = ll.Head
			for slowPtr != fastPtr {
				slowPtr = slowPtr.Next
				fastPtr = fastPtr.Next
			}

			return slowPtr
		}
	}

	return nil

}

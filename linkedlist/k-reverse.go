package linkedlist

import (
	"math"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
K Reverse In Linked List

1. You are given a partially written LinkedList class.
2. You are required to complete the body of kReverse function. The function is expected to tweak the list such that all groups of k elements in the list get reversed and linked. If the last set has less than k elements, leave it as it is (don't reverse).
3. Input and Output is managed for you.

Sample Input
11
1 2 3 4 5 6 7 8 9 10 11
3
100
200

Sample Output
1 2 3 4 5 6 7 8 9 10 11
3 2 1 6 5 4 9 8 7 10 11
100 3 2 1 6 5 4 9 8 7 10 11 200
*/

func TestKRevLL() {
	k := 11
	ll := utils.InitLL()
	for i := 0; i < k; i++ {
		ll.AddRear(utils.InitLLNode(i+1, ""))
	}

	//ll.PrintLL()

	kRevLL(ll, 3)
	ll.PrintLL()
}

func kRevLL(ll *utils.LinkedList, k int) {
	currNode := ll.Head

	loop := int(math.Floor(float64(ll.Count) / float64(k)))
	var prevll *utils.LLNode

	for i := 0; i < loop; i++ {
		var currtemp *utils.LLNode
		var initemp *utils.LLNode
		for j := k; j > 0; j-- {
			temp := currNode
			currNode = currNode.Next
			if currtemp == nil {
				initemp = temp
			}
			temp.Next = currtemp
			currtemp = temp
		}

		if prevll != nil {
			prevll.Next = currtemp
			prevll = initemp
		} else {
			ll.Head = currtemp
			prevll = initemp
		}
	}

	for currNode != nil {
		prevll.Next = currNode
		currNode = currNode.Next
		prevll = prevll.Next
	}

}

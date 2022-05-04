package linkedlist

import (
	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Multiply Two Linkedlist

1. You are given two single linkedlist of digits.
2. The most significant digit comes first and each of their nodes contain a single digit. Multiply the two numbers and return it as a linked list.
3. You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input
1->2->3->4->5->null
7->8->9->null

Output
9->7->4->0->2->0->5->null

Example
Sample Input
6
6 1 3 2 4 0
2
3 5

Sample Output
2 1 4 6 3 4 0 0
*/

func TestmultiLL() {
	ll1 := utils.InitLL()

	ll1.AddRear(utils.InitLLNode(1, ""))
	ll1.AddRear(utils.InitLLNode(2, ""))
	ll1.AddRear(utils.InitLLNode(3, ""))
	ll1.AddRear(utils.InitLLNode(4, ""))
	ll1.AddRear(utils.InitLLNode(5, ""))

	// ll1.AddRear(utils.InitLLNode(6, ""))
	// ll1.AddRear(utils.InitLLNode(1, ""))
	// ll1.AddRear(utils.InitLLNode(3, ""))
	// ll1.AddRear(utils.InitLLNode(2, ""))
	// ll1.AddRear(utils.InitLLNode(4, ""))
	// ll1.AddRear(utils.InitLLNode(0, ""))

	ll2 := utils.InitLL()

	ll2.AddRear(utils.InitLLNode(7, ""))
	ll2.AddRear(utils.InitLLNode(8, ""))
	ll2.AddRear(utils.InitLLNode(9, ""))

	// ll2.AddRear(utils.InitLLNode(3, ""))
	// ll2.AddRear(utils.InitLLNode(5, ""))

	mult2LL(ll1, ll2)
}

func mult2LL(ll1 *utils.LinkedList, ll2 *utils.LinkedList) {

	revLL(ll1)
	revLL(ll2)

	temp := ll2.Head
	multiLL := make([]*utils.LinkedList, ll2.Count)
	k := 0

	for temp != nil {
		currmultill := multiEachEle(ll1, temp.Key)
		for i := 0; i < k; i++ {
			currmultill.AddRear(utils.InitLLNode(0, ""))
		}
		multiLL[k] = currmultill
		k++
		temp = temp.Next
	}

	sumLL := multiLL[0]
	for i := 1; i < ll2.Count; i++ {
		sumLL = add2ll(sumLL, multiLL[i])
	}

	sumLL.PrintLL()

}

func multiEachEle(ll *utils.LinkedList, val int) *utils.LinkedList {
	resl := utils.InitLL()

	temp := ll.Head
	var carry int
	for temp != nil {
		sum := (temp.Key * val) + carry

		if sum > 9 {
			resl.AddFront(utils.InitLLNode(sum%10, ""))
			carry = sum / 10
		} else {
			resl.AddFront(utils.InitLLNode(sum, ""))
			carry = 0
		}

		temp = temp.Next
	}

	if carry > 0 {
		resl.AddFront(utils.InitLLNode(carry, ""))
	}

	return resl
}

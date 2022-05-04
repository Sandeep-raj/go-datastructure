package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Unfold Of Linkedlist

Given a singly linkedlist : l0 -> ln -> l1 -> ln-1 -> l2 -> ln-2 -> l3 -> ln-3 -> .....
reorder it :  l0 -> l1 -> l2 -> l3 -> l4 -> l5 -> l6 ..... -> ln-1 -> ln
for more information watch video.

Input
1->7->2->6->3->5->4->null

Output
1->2->3->4->5->6->7->null

Example
Sample Input
5 1 1 4 4 6 6 9 9

Sample Output
5 1 4 6 9 9 6 4 1
*/

func TestUnfoldLL() {
	ll := utils.InitLL()
	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(7, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(6, ""))
	ll.AddRear(utils.InitLLNode(3, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(4, ""))

	unfoldLL(ll)

	ll.AddFront(utils.InitLLNode(25, ""))
	ll.AddRear(utils.InitLLNode(11, ""))

	ll.PrintLL()
}

func unfoldLL(ll *utils.LinkedList) {
	temp := ll.Head
	var reartemp *utils.LLNode
	for temp != ll.Rear && temp.Next != ll.Rear {
		nextnode := temp.Next
		temp.Next = nextnode.Next

		rearNext := ll.Rear.Next
		ll.Rear.Next = nextnode
		nextnode.Next = rearNext
		if rearNext == nil {
			reartemp = nextnode
		}

		temp = temp.Next
	}

	ll.Rear = reartemp
}

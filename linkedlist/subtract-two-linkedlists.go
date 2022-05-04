package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Subtract Two Linked Lists

1. You are given two linked lists with N and M nodes respectively.
  2. The linked list as a whole represents a large positive number.
  3. You have to write a function that subtracts the smaller number
     from larger one and returns a pointer to the resultant linked list.
  4. display is a utility function which displays the contents of Linked List,
     feel free to use it for debugging purposes.
  5. main takes input from the users and creates the Linked Lists. You can use
     display to know its contents.
  6. This is a functional problem.
  7. You should code only the sublinkedList function. It takes as input the
     heads of the first and second linked list respectively. It should find the
     difference of two linked lists and return a pointer to the resultant linked list.
  8. Don't change the code of Node, main and display.

Sample Input
7
0 0 0 1 0 0 0
4
0 0 1 0

Sample Output
0 0 0 0 9 9 0
*/

func TestSub2LL() {
	ll1 := utils.InitLL()
	ll1.AddRear(utils.InitLLNode(0, ""))
	ll1.AddRear(utils.InitLLNode(0, ""))
	ll1.AddRear(utils.InitLLNode(0, ""))
	ll1.AddRear(utils.InitLLNode(1, ""))
	ll1.AddRear(utils.InitLLNode(0, ""))
	ll1.AddRear(utils.InitLLNode(0, ""))
	ll1.AddRear(utils.InitLLNode(0, ""))

	ll2 := utils.InitLL()
	ll2.AddRear(utils.InitLLNode(0, ""))
	ll2.AddRear(utils.InitLLNode(0, ""))
	ll2.AddRear(utils.InitLLNode(1, ""))
	ll2.AddRear(utils.InitLLNode(0, ""))

	sub2LL(ll1, ll2)
}

func sub2LL(ll1 *utils.LinkedList, ll2 *utils.LinkedList) {

	if ll1.Count < ll2.Count || (ll1.Count == ll2.Count && ll1.Head.Key < ll2.Head.Key) {
		temp := ll1
		ll1 = ll2
		ll2 = temp
	}

	resLL := utils.InitLL()

	node1 := ll1.Head
	node2 := ll2.Head

	sub2LLhelper(node1, ll1.Count, node2, ll2.Count, resLL)

	resLL.PrintLL()

}

func sub2LLhelper(node1 *utils.LLNode, index1 int, node2 *utils.LLNode, index2 int, resll *utils.LinkedList) int {
	if node1 == nil || node2 == nil {
		return 0
	}

	var carry int

	if index1 > index2 {
		carry = sub2LLhelper(node1.Next, index1-1, node2, index2, resll)
	} else {
		carry = sub2LLhelper(node1.Next, index1-1, node2.Next, index2-1, resll)
	}

	if index1 == index2 {
		if node1.Key-carry >= node2.Key {
			resll.AddFront(utils.InitLLNode(node1.Key-carry-node2.Key, ""))
			return 0
		} else {
			resll.AddFront(utils.InitLLNode((10+node1.Key)-carry-node2.Key, ""))
			return 1
		}
	} else {
		if node1.Key-carry >= 0 {
			resll.AddFront(utils.InitLLNode(node1.Key-carry, ""))
			return 0
		} else {
			resll.AddFront(utils.InitLLNode((10+node1.Key)-carry, ""))
			return 1
		}
	}
}

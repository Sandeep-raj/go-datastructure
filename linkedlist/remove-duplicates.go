package linkedlist

import "github.com/Sandeep-raj/go-datastructure/utils"

/*
Remove Duplicates In A Sorted Linked List

1. You are given a partially written LinkedList class.
2. You are required to complete the body of removeDuplicates function. The function is called on a sorted list. The function must remove all duplicates from the list in linear time and constant space
3. Input and Output is managed for you.

Sample Input
10
2 2 2 3 3 5 5 5 5 5

Sample Output
2 2 2 3 3 5 5 5 5 5
2 3 5
*/

func TestRemDupsLL() {
	ll := utils.InitLL()
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(3, ""))
	ll.AddRear(utils.InitLLNode(3, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(5, ""))
	ll.AddRear(utils.InitLLNode(5, ""))

	res := remDups(ll)
	res.PrintLL()
}

func remDups(ll *utils.LinkedList) *utils.LinkedList {
	resLL := utils.InitLL()
	temp := ll.Head

	resLL.AddRear(utils.InitLLNode(temp.Key, temp.Data))
	temp = temp.Next
	for temp != nil {
		if temp.Key != resLL.Rear.Key {
			resLL.AddRear(utils.InitLLNode(temp.Key, temp.Data))
		}

		temp = temp.Next
	}

	return resLL
}

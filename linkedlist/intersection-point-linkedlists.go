package linkedlist

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Intersection Point Of Linked Lists

1. You are given a partially written LinkedList class.
2. You are required to complete the body of findIntersection function. The function is passed two linked lists which start separately but merge at a node and become common thereafter. The function is expected to find the point where two linked lists merge. You are not allowed to use arrays to solve the problem.
3. Input and Output is managed for you.

Sample Input
7
1 2 3 4 5 6 7
5
9 8 5 6 7

Sample Output
5
*/

func TestIntersectionPointLL() {
	interLL := utils.InitLL()
	interLL.AddRear(utils.InitLLNode(5, ""))
	interLL.AddRear(utils.InitLLNode(6, ""))
	interLL.AddRear(utils.InitLLNode(7, ""))

	ll1 := utils.InitLL()
	ll1.AddRear(utils.InitLLNode(1, ""))
	ll1.AddRear(utils.InitLLNode(2, ""))
	ll1.AddRear(utils.InitLLNode(3, ""))
	ll1.AddRear(utils.InitLLNode(4, ""))

	ll1.Rear.Next = interLL.Head
	ll1.Rear = interLL.Rear
	ll1.Count = ll1.Count + interLL.Count

	ll2 := utils.InitLL()
	ll2.AddRear(utils.InitLLNode(9, ""))
	ll2.AddRear(utils.InitLLNode(8, ""))

	ll2.Rear.Next = interLL.Head
	ll2.Rear = interLL.Rear
	ll2.Count = ll2.Count + interLL.Count

	intersectionLLDiff(ll1, ll2)
	intersectionLLFloydCycle(ll1, ll2)
}

func intersectionLLDiff(ll1 *utils.LinkedList, ll2 *utils.LinkedList) {
	if ll1.Count < ll2.Count {
		temp := ll1
		ll1 = ll2
		ll2 = temp
	}

	k := ll1.Count - ll2.Count
	temp1 := ll1.Head
	for k > 0 {
		temp1 = temp1.Next
		k--
	}

	interNode := interLLDiffhelper(temp1, ll2.Head)
	log.Printf("%+v", interNode)
}

func interLLDiffhelper(node1 *utils.LLNode, node2 *utils.LLNode) *utils.LLNode {
	if node1 == nil || node2 == nil {
		return nil
	} else if node1 == node2 {
		return node1
	}

	return interLLDiffhelper(node1.Next, node2.Next)
}

func intersectionLLFloydCycle(ll1 *utils.LinkedList, ll2 *utils.LinkedList) {
	ll1.Rear.Next = ll2.Head

	node := cycleNodeLL(ll1)

	if node != nil {
		log.Printf("%+v", node)
	} else {
		log.Print("No intersection point")
	}
}

package linkedlist

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Reverse Linked List (pointer Iterative)

1. You are given a partially written LinkedList class.
2. Here is a list of existing functions:
    2.1 addLast - adds a new element with given value to the end of Linked List
    2.2. display - Prints the elements of linked list from front to end in a single line.
    All elements are separated by space
    2.3. size - Returns the number of elements in the linked list.
    2.4. removeFirst - Removes the first element from Linked List.
    2.5. getFirst - Returns the data of first element.
    2.6. getLast - Returns the data of last element.
    2.7. getAt - Returns the data of element available at the index passed.
    2.8. addFirst - adds a new element with given value in front of linked list.
    2.9. addAt - adds a new element at a given index.
    2.10. removeLast - removes the last element of linked list.
    2.11. removeAt - remove an element at a given index
3. You are required to complete the body of reversePI function. The function should be an iterative function and should reverse the contents of linked list by changing the "next" property of nodes.
4. Input and Output is managed for you.

Sample Input
addFirst 10
addFirst 20
addLast 30
addLast 40
addLast 50
addFirst 60
removeAt 2
display
reversePI
display
quit

Sample Output
60 20 30 40 50
50 40 30 20 60
*/

func TestRevLL() {
	ll := utils.InitLL()
	ll.AddFront(utils.InitLLNode(10, "test"))
	ll.AddFront(utils.InitLLNode(20, "test"))
	ll.AddRear(utils.InitLLNode(30, "test"))
	ll.AddRear(utils.InitLLNode(40, "test"))
	ll.AddRear(utils.InitLLNode(50, "test"))
	ll.AddFront(utils.InitLLNode(60, "test"))

	ll.PrintLL()
	revLL(ll)
	log.Print("----------------------")
	ll.PrintLL()

	ll.AddFront(utils.InitLLNode(100, "test"))
	ll.AddFront(utils.InitLLNode(200, "test"))
	ll.AddRear(utils.InitLLNode(300, "test"))
	ll.AddRear(utils.InitLLNode(400, "test"))

	log.Print("----------------------")
	ll.PrintLL()

}

func revLL(ll *utils.LinkedList) {
	var head *utils.LLNode
	currnode := ll.Head

	for currnode != nil {
		temp := currnode
		currnode = currnode.Next

		temp.Next = head
		head = temp
	}

	ll.Rear = ll.Head
	ll.Head = head
}

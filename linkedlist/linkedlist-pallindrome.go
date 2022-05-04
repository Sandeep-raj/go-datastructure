package linkedlist

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Is Linked List A Palindrome?

1. You are given a partially written LinkedList class.
2. You are required to complete the body of IsPalindrome function. The function is expected to check if the linked list is a palindrome or not and return true or false accordingly.
3. Input and Output is managed for you.

Sample Input
5
1 2 3 2 1

Sample Output
true
*/

func TestLLPallindrome() {
	ll := utils.InitLL()
	ll.AddRear(utils.InitLLNode(1, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(3, ""))
	ll.AddRear(utils.InitLLNode(2, ""))
	ll.AddRear(utils.InitLLNode(1, ""))

	llPallindrome(ll)
}

func llPallindrome(ll *utils.LinkedList) {
	_, ispallindrome := llpallindromehelper(ll.Head, ll.Head)
	if ispallindrome {
		log.Print("Linked List is Pallindrome")
	} else {
		log.Print("Linked List is not Pallindrome")
	}
}

func llpallindromehelper(lfront *utils.LLNode, lback *utils.LLNode) (*utils.LLNode, bool) {
	if lback == nil {
		return lfront, true
	}

	lfront, isPallindrome := llpallindromehelper(lfront, lback.Next)

	if isPallindrome && lfront.Key == lback.Key {
		return lfront.Next, true
	} else {
		return lfront, false
	}
}

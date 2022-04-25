package utils

import "log"

type BST_node struct {
	Key   int
	Data  interface{}
	Left  *BST_node
	Right *BST_node
}

type BST struct {
	Count int
	Root  *BST_node
}

func InitBSTNode(key int, data interface{}) *BST_node {
	return &BST_node{
		Key:   key,
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func InitBST() *BST {
	return &BST{
		Count: 0,
		Root:  nil,
	}
}

func (bst *BST) BSTInsert(node BST_node) {
	if bst.Count == 0 {
		bst.Root = &node
		bst.Count++
	} else {
		temp := bst.Root

		for {
			if node.Key <= temp.Key {
				if temp.Left == nil {
					temp.Left = &node
					break
				} else {
					temp = temp.Left
				}
			} else {
				if temp.Right == nil {
					temp.Right = &node
					break
				} else {
					temp = temp.Right
				}
			}
		}

		bst.Count++
	}
}

func (bst *BST) BSTSearch(key int) *BST_node {
	temp := bst.Root

	for temp != nil {
		if temp.Key == key {
			return temp
		} else {
			if key < temp.Key {
				temp = temp.Left
			} else {
				temp = temp.Right
			}
		}
	}

	return nil
}

func (bst *BST) BSTDelete(key int) {
	temp := bst.Root
	if temp.Key == key {
		if bst.Count == 1 {
			bst.Root = nil
		} else if temp.Left == nil {
			bst.Root = temp.Right
		} else {
			bst.remove(bst.Root)
		}
		bst.Count--
	} else {
		parentNode := temp
		for {
			if parentNode.Key < key {
				if parentNode.Right.Key == key {
					break
				} else {
					parentNode = parentNode.Right
				}
			} else {
				if parentNode.Left.Key == key {
					break
				} else {
					parentNode = parentNode.Left
				}
			}
		}

		if parentNode.Right != nil && parentNode.Right.Key == key {
			if parentNode.Right.Left == nil {
				parentNode.Right = parentNode.Right.Right
			} else {
				bst.remove(parentNode.Right)
			}
		}

		if parentNode.Left != nil && parentNode.Left.Key == key {
			if parentNode.Left.Left == nil {
				parentNode.Left = parentNode.Left.Right
			} else {
				bst.remove(parentNode.Left)
			}
		}

		bst.Count--
	}
}

func (bst *BST) remove(currNode *BST_node) {
	temp := currNode.Left
	if temp.Right == nil {
		currNode.Key = temp.Key
		currNode.Data = temp.Data
		currNode.Left = temp.Left
	} else {
		for temp.Right.Right != nil {
			temp = temp.Right
		}

		maxNode := temp.Right
		if maxNode.Left == nil {
			currNode.Key = maxNode.Key
			currNode.Data = maxNode.Data
			temp.Right = nil
		} else {
			currNode.Key = maxNode.Key
			currNode.Data = maxNode.Data
			bst.remove(maxNode)
		}
	}
}

func (bst *BST) BSTTraversal(order string) {
	switch order {
	case "pre":
		bst.preOrderTraversal(bst.Root)
	case "in":
		bst.inOrderTraversal(bst.Root)
	case "post":
		bst.postOrderTraversal(bst.Root)
	}
}

func (bst *BST) preOrderTraversal(node *BST_node) {
	if node == nil {
		return
	}

	bst.preOrderTraversal(node.Left)
	log.Print(node.Key)
	bst.preOrderTraversal(node.Right)
}

func (bst *BST) inOrderTraversal(node *BST_node) {
	if node == nil {
		return
	}

	log.Print(node.Key)
	bst.inOrderTraversal(node.Left)
	bst.inOrderTraversal(node.Right)
}

func (bst *BST) postOrderTraversal(node *BST_node) {
	if node == nil {
		return
	}

	bst.postOrderTraversal(node.Right)
	log.Print(node.Key)
	bst.postOrderTraversal(node.Left)
}

func (bst *BST) BSTMax() *BST_node {
	if bst.Count == 0 {
		return nil
	}

	temp := bst.Root

	for temp.Right != nil {
		temp = temp.Right
	}

	return temp
}

func (bst *BST) BSTMin() *BST_node {
	if bst.Count == 0 {
		return nil
	}

	temp := bst.Root

	for temp.Left != nil {
		temp = temp.Left
	}

	return temp
}
